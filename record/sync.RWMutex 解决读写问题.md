# sync.RWMutex 解决读写问题

在一个共享的数据需要被多个线程访问的时候就会出现很多读写问题(由于读写问题有很多变种，所以用许多来形容)。有两种类型的线程需要访问数据—读线程和写线程。读线程仅仅读数据，写线程修改数据。当写线程有权限访问数据的时候，其他线程(包括读线程和写线程)是不可以访问这个共享的数据。这个限制在日常生活中是真的发生的，当写线程无法以原子性的操作修改数据的时候，读线程必须被阻塞，以防读取到脏数据(译者注：为了使得说明的更加清晰，后面用写goroutine和读goroutine分别代替写读线程)。有许多核心问题的变种如下：

- 写线程不能处于饥饿状态(无限的等待他们执行的机会)
- 读线程不能处于饥饿状态
- 不应该有线程可以处于饥饿状态

多读/一写的互斥锁的具体实现(例如 [sync.RWMutex](https://golang.org/pkg/sync/#RWMutex))解决了读写问题的的其中之一。让我们看看这在Go总是如何做到的以及是它给到了一种到达什么样程度的保证。

作为一个奖励，我们深入的理解简化的竞争互斥锁。

## 使用方式

在深入到实现细节之前，让我们看看如何在实践中使用`sync.RWMutex`。如下的程序使用了读写互斥用来保护关键部分的操作—`sleep()`。通过对关键部分正在执行的读写线程进行计数来对整个程序执行过程的可视化([source code](https://play.golang.org/p/xoiqW0RQQE9))。

```go
package main
import (
    "fmt"
    "math/rand"
    "strings"
    "sync"
    "time"
)
func init() {
    rand.Seed(time.Now().Unix())
}
func sleep() {
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
func reader(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
    sleep()
    m.RLock()
    c <- 1
    sleep()
    c <- -1
    m.RUnlock()
    wg.Done()
}
func writer(c chan int, m *sync.RWMutex, wg *sync.WaitGroup) {
    sleep()
    m.Lock()
    c <- 1
    sleep()
    c <- -1
    m.Unlock()
    wg.Done()
}
func main() {
    var m sync.RWMutex
    var rs, ws int
    rsCh := make(chan int)
    wsCh := make(chan int)
    go func() {
        for {
            select {
            case n := <-rsCh:
                rs += n
            case n := <-wsCh:
                ws += n
            }
            fmt.Printf("%s%s\n", strings.Repeat("R", rs),
                    strings.Repeat("W", ws))
        }
    }()
    wg := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go reader(rsCh, &m, &wg)
    }
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go writer(wsCh, &m, &wg)
    }
    wg.Wait()
}
```

样例输出:

```
W

R
RR
RRR
RRRR
RRRRR
RRRR
RRR
RRRR
RRR
RR
R

W

R
RR
RRR
RRRR
RRR
RR
R

W
```

每当goroutines(包括读和写)执行关键部分的数据发生改变的时候，新的一行就会打印出来。通过这些打印的行可以表明RWMutex`允许至少一个读的goroutine或者仅仅允许一个写的goroutine的。

另外很重要的就是当写goroutine执行`Lock()`的时候，新来的读goroutine是会被阻塞的，这个也会在接下来进行讨论。写goroutine会等待当前已经开始执行的读goroutine执行完成它们的任务(departure)，在它们执行完成之后，写goroutine就会开始执行了。上述的过程可以在样例输出中得到展示：在某个时间点，读goroutine一个一个的减少之后，写goroutine开始展出出来。

```tex
...
RRRRR
RRRR
RRR
RR
R

W

...
```

一旦写goroutine执行完成，之前阻塞的读goroutine就会恢复并且其它的写goroutine也可以开始执行了(译者注：可以调用`Lock()`)。值得注意的是，当写goroutine执行完成，如果有读goroutine和写goroutine同时等待，那么其中的读goroutine会优先执行。正是这种新来的写goroutine需要等待之前阻塞的读goroutine执行之后再执行才不会使得读goroutine或者写goroutine处于饥饿状态(译者注：如果写goroutine优先执行的话，那么读goroutine就会处于饥饿状态，就是永远也没有机会执行)。

## 实现

![](https://miro.medium.com/max/1280/1*Gg_vmyWlU35r3w_L4r4SYw.jpeg)

>需要注意的是此次讲解版本是[718d6c58](https://github.com/golang/go/blob/718d6c5880fe3507b1d224789b29bc2410fc9da5/src/sync/rwmutex.go)，之后的实现可能会变更

`RWMutex`提供了两个方法(`RLock`和`RUnlock`)给读goroutine使用，两个方法(`Lock`和`Unlock`)给写goroutine使用。

### RLock

为了简洁，我们跳过竞争状态监测部分的代码(它们将会被`…`替代)

```go
func (rw *RWMutex) RLock() {
    ...
    if atomic.AddInt32(&w.readerCount, 1) < 0 {    
        runtime_SemacquireMutex(&rw.readerSem, false)
    }
    ...
}
```

变量`readerCount`是一个`int32`，用于表示等待的读goroutine数量—已经开始执行或者被写goroutine阻塞的读goroutine数量。这就是已经调用`Rlock`，而还没调用`RUnlock`的读goroutine的数量。

[atomic.AddInt32](https://golang.org/pkg/sync/atomic/#AddInt32)等价于原子性的如下操作：

```go
*addr += delta
return *addr
```

其中addr是一个`*int32`变量，delta是一个`int32`型的变量。由于是原子性的操作，所以在对addr添加delta的时候不会影响其他线程(更多的 [fetch-and-add](https://en.wikipedia.org/wiki/Fetch-and-add))

>在没有写goroutine的时候，变量`readerCount`永远都是一个大于或等于0的值。并且，读goroutine可以不阻塞并且快速的执行，其中仅仅掺杂的操作是`atomic.AddInt32`

### 信号量

信号量是一个由Edsger Dijkstra发明的数据结构，这种数据结构对于同步问题非常有用。信号量是i一个整型变量，有两个操作：

- acquire(也被称为wait，decrement或者P)
- release(signal，increment or V)

`acquire`操作用于对信号量的值减1。在信号量的值是负数的时候，线程会阻塞并不会恢复，只有在其他线程增加信号量的值的时候线程才会可以继续执行；在信号量的值为正数的时候，线程是可以无阻碍的执行的。

`release`操作会对信号量的值加1。如果有阻塞的线程，那么这些阻塞线程中的一个可以得到执行。

go中的`runtime`包提供了两个函数`runtime_SemacquireMutex`和`runtime_Semrelease用于`实现`sync.RWMutex`。

### Lock

```go
func (rw *RWMutex) Lock() {
    ...
    rw.w.Lock()
    r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
    if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {     
        runtime_SemacquireMutex(&rw.writerSem, false)
    }
    ...
}
```

`Lock`方法用于在写goroutine对共享的数据获取独有权限的时候使用。首先，写goroutine将获得一个互斥锁，用于禁止其他写goroutine访问共享数据。这个互斥锁会在函数`Unlock`调用的时候，立即释放。然后会使`readerCount`减去`rwmutexMaxReaders(1 << 30)`。当`readerCount`变成一个负值的时候，会使阻塞后面到来的读goroutine:

```go
if atomic.AddInt32(&rw.readerCount, 1) < 0 {
    // A writer is pending, wait for it.    
    runtime_SemacquireMutex(&rw.readerSem, false)
}
```

由于新来的读goroutine会被阻塞，那么已经开始运行的读的goroutine会怎么办呢？变量`readerWait`用于记录信号量中当前运行的和阻塞的读goroutine数量。此信号量会在最后一个读goroutine使用`RUnlock`方法的时候释放掉互斥锁，`RUnlock`会在下面讨论。

如果没有正在运行的读goroutine，那么接下来写goroutine会直接运行。

### rwmutexMaxReaders

在 [rwmutex.go](https://github.com/golang/go/blob/718d6c5880fe3507b1d224789b29bc2410fc9da5/src/sync/rwmutex.go)中有一个常量:

```go
const rwmutexMaxReaders = 1 << 30
```

这个常量是用来做什么的？并且`1 << 30`的含义是什么呢？

变量`readerCount`是一个 [int32](https://golang.org/pkg/builtin/#int32) ，其取值范围如下：

```go
[-1 << 31, (1 << 31) — 1] or [-2147483648, 2147483647]
```

`RWMutex`使用此变量用于表示调用了`RLock`函数的读goroutine和阻塞的写goroutine，在`Lock`方法中：

```go
r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
```

变量`readerCount`会减少`1 << 30`，当`readerCount`为负数的时候表示有一个阻塞的写goroutine，并且`readerCount + rwmutexMaxReaders`就是当前已经开始执行的读goroutines(译者注：其实就是在调用`Lock`之前已经调用`RLock`的读goroutine数量，因为`readerCount - rwmutexMaxReaders + rwmutexMaxReaders = readerCount`)。此变量也限制了可以有多少个读goroutine可以访问共享数据结构。如果我们大于等于`rwmutexMaxReaders`数量的读goroutine，那么`readerCount`就会是一个大于等于零的值，所以实际的最大读goroutine的数量是：

```go
rwmutexMaxReaders-1
```

这个值虽然有限制，但是仍然超过10亿——1073741823。

### RUnlock

```go
func (rw *RWMutex) RUnlock() {
    ...
    if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
        if r+1 == 0 || r+1 == -rwmutexMaxReaders {
            race.Enable()
            throw("sync: RUnlock of unlocked RWMutex")
        }
        // A writer is pending.
        if atomic.AddInt32(&rw.readerWait, -1) == 0 {
            // The last reader unblocks the writer.       
            runtime_Semrelease(&rw.writerSem, false)
        }
    }
    ...
}
```

此方法会会对`readerCount`减1(`Rlock`会是`readerCount`加1)。如果`readerCount`是一个负数，那表示有一个写goroutine正在等待或者运行，这是由于`readerCount`在调用`Lock`的时候减去了`rwmutexMaxReaders`。然后会检查是否是最后一个读的goroutine运行完成，如果为0则表示已经完成，那么写goroutine在最后获取到信号量。

### Unlock

```go
func (rw *RWMutex) Unlock() {
    ...
    r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
    if r >= rwmutexMaxReaders {
        race.Enable()
        throw("sync: Unlock of unlocked RWMutex")
    }
    for i := 0; i < int(r); i++ {
        runtime_Semrelease(&rw.readerSem, false)
    }
    rw.w.Unlock()
    ...
}
```

在释放由写goroutine持有的互斥锁的时候，第一步就是对`readerCount`加1，所以这个值会是一个非负数了。如果`readerCount`大于0则说明有些读goroutine在等待着写goroutine完成共享数据的修改，然后需要启动这些等待的读goroutine了。然后写goroutine持有的互斥锁会被释放，从而允许其他的写goroutine来获取`RWMutex`来获得修改共享数据的权限。

方法`Unlcok`和`RUnlock`会在读goroutine或者写goroutine释放一个没有锁住的互斥锁的时候([源码](https://play.golang.org/p/YMdFET74olU))抛异常：

```go
m := sync.RWMutex{}
m.Unlock()
```

输出：

```go
fatal error: sync: Unlock of unlocked RWMutex
...
```



## 递归读中的锁

文档中写道

> 如果一个goroutine持有一个读写锁的读权限，另外一个goroutine可能会调用lock。任何一个goroutine都不能期待可以获取一个读锁直到最初的读锁释放之后。特别的是，这个可以预防递归的读锁。这个可以确保锁最终是可以被获取的；一个阻塞的锁防止了新的readers的获取。

`RWMutex`的运行方式是，如果有一个阻塞的写goroutine，那么所有试图调用`RLock`的读goroutine都会被阻塞，不管读goroutine是否已经获得读锁了([源码](https://play.golang.org/p/XNndlaZ6Ema)):

```go
package main
import (
    "fmt"
    "sync"
    "time"
)
var m sync.RWMutex
func f(n int) int {
    if n < 1 {
        return 0
    }
    fmt.Println("RLock")
    m.RLock()
    defer func() {
        fmt.Println("RUnlock")
        m.RUnlock()
    }()
    time.Sleep(100 * time.Millisecond)
    return f(n-1) + n
}
func main() {
    done := make(chan int)
    go func() {
        time.Sleep(200 * time.Millisecond)
        fmt.Println("Lock")
        m.Lock()
        fmt.Println("Unlock")
        m.Unlock()
        done <- 1
    }()
    f(4)
    <-done
}
```

输出如下：

```go
RLock
RLock
RLock
Lock
RLock
fatal error: all goroutines are asleep - deadlock!
```

>译者注：我觉得这个其实也不是故意设计的不允许递归的调用`RLock`方法，这个在实际使用会出现问题是由于在写goroutine调用`Lock`的时候，会等待之前调用了`RLock`方法的读goroutine执行完成，而如果之前调用的`RLock`是递归的调用的话，后面还会出现调用`RLock`的读goroutine，之前调用了读goroutine则不能继续往下执行了。下面展示的是递归的调用过程，Lock方法不在递归的调用过程之中，放在下面这个位置只是为了说明出现的时机。
>
>>RLock
>>
>>>RLock
>>>
>>>>RLock
>>>>
>>>>>**Lock**(表示加入的时机，并不参与递归的过程)
>>>>>
>>>>>>RLock
>
>因为有一个写 goroutine 来了，会阻塞后来的读 goroutine 的执行。后来的读 goroutine 不能执行了，也就是递归也执行不下去了。递归执行不下去，就会造成之前获得读锁得不到释放，然后就写 goroutine 也没有机会执行，造成的结果就是读 goroutine 和 写 goroutine 都不会执行。最后的结果就和报错一样`fatal error: all goroutines are asleep - deadlock!`
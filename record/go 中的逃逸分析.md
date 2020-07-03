# go 中的逃逸分析

## 前言

本文翻译自https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html，省略了部分个人认为不重要的内容，也有部分内容作了一定的修改。

## 堆

除了栈以外，堆是用于存储数据的第二种方式。由于堆并不会像栈一样自动的清除，所以使用堆代价较大一些。代价大根本原因是需要通过垃圾回收(GC)来清理其中的无用数据。当 GC 运行的时候，它会消耗约 25% 的 cpu 资源。另外，它可能会造成微妙级别的"stop the world"延迟。由于堆中的内存管理一致比较难懂并且容易出错，使用 GC 的优点就是不用管堆中的内存。

堆中的数据组成了 Go 中的内存分配方式。这种分配方式给了 GC 压力，因为不再被指针引用的值需要清理。并且需要检查和清理的值越多，每次 GC 运行时需要做的工作就越多。所以调度算法需要在 GC 运行的频次和堆的大小上进行平衡。

## 共享栈

在 Go 中， goroutine 并不允许拥有一个指向其他 goroutine 运行时栈中的内存的指针。这是由于 goroutine 使用的栈内存在需要扩张或者收缩的时候会被新的内存块所占用。如果 runtime 需要跟踪所有指向其他 goroutine 栈空间的指针，这将耗费巨大的代价。并且由于需要更新这些指针的地址，"stop the world"的延迟也会变得巨大。

下面是一个例子，由于 goroutine 的递归，需要使用的空间在不断的增大，其空间地址变换了好几次。可以查看`println`打印的`string`的地址变换了多次。(译者注：不同的版本，变换的次数可能会不一致)

```go
// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how stacks grow/change.
package main

// Number of elements to grow each stack frame.
// Run with 10 and then with 1024
const size = 1024

// main is the entry point for the application.
func main() {
   s := "HELLO"
   stackCopy(&s, 0, [size]int{})
}

// stackCopy recursively runs increasing the size
// of the stack.
func stackCopy(s *string, c int, a [size]int) {
   println(c, s, *s)

   c++
   if c == 10 {
      return
   }

   stackCopy(s, c, a)
}
```

## 逃逸分析

一个值被分配的内存地址处于函数栈帧之外，就会被分配到堆。逃逸分析算法的工作就是找到这种情况，并保证程序的一体性。一体性是指保证操作某些值的时候是准确的、一致的以及高效的。

通过下面的例子，我们可以学到逃逸算法背后的基本机制。

**Listing 1**

```go
01 package main
02
03 type user struct {
04     name  string
05     email string
06 }
07
08 func main() {
09     u1 := createUserV1()
10     u2 := createUserV2()
11
12     println("u1", &u1, "u2", u2)
13 }
14
15 //go:noinline
16 func createUserV1() user {
17     u := user{
18         name:  "Bill",
19         email: "bill@ardanlabs.com",
20     }
21
22     println("V1", &u)
23     return u
24 }
25
26 //go:noinline
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

在代码中使用了指令`go:noinline`来避免编译器将这些函数的操作直接内联到`mian`函数之中(译者注：内联的含义就是不调用函数，而把函数部分的操作直接嵌入到调用此函数的位置)。内联将会清除函数的调用并使得例子复杂化。在下一篇文章我将会介绍内联操作的副作用。

打印的结果如下，可以看到函数`createUserV1`返回结果的地址和创建时候的地址是不一致的，而`createUserV2`则是一致的。

```
V1 0xc00003ff00
V2 0xc000004060
u1 0xc00003ff58 u2 0xc000004060
```



在 Listing 1，你会看到代码中有两个函数都创建了一个`user`，并返回给调用者。函数的 Version 1 版本返回的是`user`的值。

**Listing 2**

```go
16 func createUserV1() user {
17     u := user{
18         name:  "Bill",
19         email: "bill@ardanlabs.com",
20     }
21
22     println("V1", &u)
23     return u
24 }
```

之所以说函数 Version 1 版本是返回值得原因是函数的返回值会被复制一份，然后返回给调用栈。这意味着调用者接受的值是函数返回值复制之后的结果。

可以看到第 17 行到第 20 行会构建`user`对象。然后再第 23 行 会把`user`复制一份，返回给调用者。在函数返回之后，栈会如下图所示

**Figure 1**

![](https://www.ardanlabs.com/images/goinggo/81_figure1.png)

可以在 Figure 1 中看到，在调用`createUserV1`之后`user`值在栈中存在两份。在函数的 Version 2 之中，返回的则是指针。

**Listing 3**

```go
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

之所以说的返回的是指针的原因是此函数创建的值被调用栈所共享。这意味着调用的函数接受的值是`user`地址所复制的值。

你可以看到第 28 行到 31 行同样的创建`user`值得部分，但是在第 34 行不一样。和返回`user`复制的结果不同，这次返回的是`user`值地址的复制结果。基于此，你可能会认为调用栈会和下图一样

**Figure 2**

![](https://www.ardanlabs.com/images/goinggo/81_figure2.png)

如果你看到的 Figure 2 是真实发生的，那么就会出现一致性的问题，因为指针所指向的地址是无效的。在`main`函数调用的下个函数的时候，被指针所指向的区域就会被覆盖。

这就是逃逸分析用来维护一致性的地方。在这个例子中，编译器会判断把`user`值构建于`createUserV2`函数之内是否是安全的，如果不安全就会把对象构建于堆。这就是函数第 28 行所发生的。

## 可读性

正如你在上一篇文章所看到的一样，函数是可以通过帧指针直接访问帧内的地址的，但是对于帧外的地址需要使用间接的方式。这也就意味着访问堆上的对象也需要以指针的形式进行访问。

记住函数`createUserV2`的代码样子

**Listing 4**

```go
27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

语法隐藏于代码中真实发生之下。第 28 行声明的变量`u`表示的是一个`user`值。Go 中的构建对象的过程并没有告诉我们对象存在于内存中的哪些地方，直到第 34 行，才知道值是需要逃逸的。这意味着，即使变量`u`代表的是一个类型为`user`的值，但是访问`user`值得时候，必须通过指针。

你可以设想栈在调用了函数之后的样子如下图

**Figure 3**

![](https://www.ardanlabs.com/images/goinggo/81_figure3.png)

函数`createUserV2`栈帧中的变量`u`，代表的值存在于堆，而不是栈。这也就意味着使用`u`访问值，需要使用指针，而不是直接的如语法表示一样的方式。你可能回想，为什么不直接把`u`作为一个指针呢？因为访问一个变量所代表的的值需要使用指针。

**Listing 5**

```go
27 func createUserV2() *user {
28     u := &user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", u)
34     return u
35 }
```

如果你这样做，你就会使你的代码不具有可读性。暂时从完整的函数离开，聚焦于`return`。

**Listing 6**

```go
34     return u
35 }
```

`return`告诉你的是什么？它所说的是复制一份`u`返回给调用者。然而，如果使用`&`操作符的时候，使用`return`，这个会告诉你什么呢？

**Listing 7**

```go
34     return &u
35 }
```

由于`&`操作符，`return`函数告诉你这个变量`u`会在栈之间共享，所以需要逃逸到堆之中。记住，指针的目的是共享数据，当你阅读代码的时候碰到操作符`&`，你可以用`共享`来替换。这在提高代码可读性上非常有用。

下面的这个例子使用了指针，但是损害了可读性

**Listing 8**

```go
01 var u *user
02 err := json.Unmarshal([]byte(r), &u)
03 return u, err
```

在第 2 行代码中，你必须共享一个指针值给`json.Unmarshal`调用，才可以正常的工作。`json.Unmarshal`调用的时候，会创建一个`user`值，并把它的地址传递给指针值。

代码所表达的如下

1. 创建一个`user`类型的指针值，并置为零型。
2. 与函数`json.Unmarshal`共享变量`u`
3. 给函数调用者返回`u`的复制值

这个代码对于表达一个由`json.Unmarshal`创建的`user`值和调用者共享并不明显。

如果通过值变量的方式，可读性会提高吗？

**Listing 9**

```go
01 var u user
02 err := json.Unmarshal([]byte(r), &u)
03 return &u, err
```

代码所表达的如下

1. 创建一个`user`类型的值，并置为零型。
2. 与函数`json.Unmarshal`共享变量`u`
3. 与调用者共享变量`u`

所有事都非常清晰。第 2 行是与函数`json.Unmarshal`的函数栈共享变量`user`， 第 3 行是与调用此函数的函数栈共享变量`user`，这样就会造成`user`逃逸到堆。

创建一个变量的时候使用值类型，共享变量的时候通过操作符`&`来提高代码的可读性。

## 编译报告

为了看到编译器的决定，我们需要使用让编译器提供报告。你需要做的就是在通过`go build`的时候，传递参数`-gcflags`。

实际上有四个层级的`-m`可以使用，但是超过 2 层的信息就足够多了。我会在接下来使用两级的`-m`

**Listing 10**

```go
$ go build -gcflags "-m -m"
./main.go:16: cannot inline createUserV1: marked go:noinline
./main.go:27: cannot inline createUserV2: marked go:noinline
./main.go:8: cannot inline main: non-leaf function
./main.go:22: createUserV1 &u does not escape
./main.go:34: &u escapes to heap
./main.go:34: 	from ~r0 (return) at ./main.go:34
./main.go:31: moved to heap: u
./main.go:33: createUserV2 &u does not escape
./main.go:12: main &u1 does not escape
./main.go:12: main &u2 does not escape
```

你可以看到编译器报告了逃逸决定。编译器说了什么？在看之前，我们先再次看看`createUserV1` 和`createUserV2` 函数的定义。

**Listing 13**

```go
16 func createUserV1() user {
17     u := user{
18         name:  "Bill",
19         email: "bill@ardanlabs.com",
20     }
21
22     println("V1", &u)
23     return u
24 }

27 func createUserV2() *user {
28     u := user{
29         name:  "Bill",
30         email: "bill@ardanlabs.com",
31     }
32
33     println("V2", &u)
34     return &u
35 }
```

从第一行开始看。

**Listing 14**

```
./main.go:22: createUserV1 &u does not escape
```

这句说的是在函数`createUserV1`内部调用`println`并没有使得`user`逃逸到堆。这个必须要检查，因为`createUserV1`和函数`println`共享了变量。

下面来看看报告中的接下来的几行。

**Listing 15**

```
./main.go:34: &u escapes to heap
./main.go:34: 	from ~r0 (return) at ./main.go:34
./main.go:31: moved to heap: u
./main.go:33: createUserV2 &u does not escape
```

这些行表示了由于第 34 行，`user`值关联的变量`u`是一个`user`类型并在第 31 行赋值，逃逸到了堆。最后一行和之前说过的一样，第 33 行调用`println`没有导致`user`值逃逸到堆。

## 结论

创建一个变量的时候决定不了此变量分配的内存位置。只有变量的共享方式才决定了编译器如何分配此变量所在的内存区域。在一个函数创建的变量需要和调用此函数的函数位置共享的时候，此变量需要逃逸到堆。还有其他的变量需要逃逸的情况，会在下一篇的博客介绍。
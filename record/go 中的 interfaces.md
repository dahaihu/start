# go 中的 interfaces

## 接口值

拥有方法的语言中通常分为两个阵营: 预先准备好静态的方法表(例如 C++ 和 Java)，或者在每次调用的方法的时候进行方法的搜索(例如 Smalltalk 以及其模仿者，包括 JavaScript 和 Python)，并加入缓存提高效率。Go 是处于两者之间的: 它有方法表，但是是在 runtime 期间计算出来的。我不知道 Go 是不是编程语言中第一个这样做的，但是它肯定不是唯一一个这样做的。

作为一个预热，一种类型为`Binary`的值是一个 64 位的整型，由两个 32 位组成(假设电脑是一个 32 位的机器)

![](http://research.swtch.com/gointer1.png)

其方法以及接口`Stringer`定义如下:

```go
type Binary uint64

func (i Binary) String() string {
    return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
    return uint64(i)
}

type Stringer interface {
    String() string
}
```

接口值是由两指针组成，一个指针指向类型相关的信息，一个指针指向其关联的数据。在给变量`b`赋值一个类型为`Stringer`的接口时候，都会设置这两个指针的值。

![](http://research.swtch.com/gointer2.png)

(图中接口包含的指针是灰色的，表示它们并不直接暴露给 go 编码者)

interface 中的第一个指针指向的一个 interface table 或者 itable (读成 i-table; 在 [runtime 源码](http://golang.org/src/pkg/runtime/iface.c#L23)，其 C 实现的名称是 `Itab`)。itable 结构中的起始部分是一些例如类型的元数据，然后是一个函数指针的列表。注意：itable 对应的类型是接口类型，而不是动态类型。在我们的例子中，接口`Stringer`的 itable 包含 `Binary`中用来实现`Stringer`的方法，只有`String`方法。`Binary`中的其他方法(例如`Get`)并不会出现在 itable 之中。

interface 中的第二个指针指向的是实际的数据，在这个例子中是`b`复制的数据。赋值语句`var s Stringer = b`会复制一份`b`的值赋给`s`，而不是让`s`的指针指向`b`，正如`var c uint64 = b`一样。如果`b`的值在后面改变了，那么`s`和`c`并不会改变。interface 中的数据可以是任意大的，但是 interface 结构中只有一个字的长度用来存储这个值，所以赋值的时候，会在堆上分配内存，并在 interface 中使用指针来指向堆上的数据。(如果实际值得大小只有一个字的长度，那么这个地方是显然是可以优化的，在后面会讲到)。

为了检测 interface 的类型(和上面的 type switch 示例一样)，Go 编译器会产生和 C 表达式`s.tab->type`一样的代码来类型指针，并对比类型。如果类型一致，那么可以通过`s.data`复制。

为了调用`s.String()`，Go 编译器会产生和 C 表达式一样的`s.tab->fun[0](s.data)`的代码：调用 itable 中合适的方法，并把 interface 中的 value 值作为方法的第一个参数(在这个例子中是一个唯一的参数)。如果你运行`8g -S x.go`(详细的细节在文章末尾)，你就可以看到这个代码。你可以注意到传递给 itable 中的函数的参数是一个 32 位的指针，而不是 64 位的指针实际指向的值。通常，接口调用方法的时候并不知道字长值的含义或者这个指针指向值的大小。interface  代码会会使得 itable 中的函数指针接受一个 32 位的存储在 interface 中的值。所以这个例子中的函数指针实际上是`(*Binary).String`而不是`Binary.String`。

这个例子中的接口只有一个方法。如果一个接口中有多个方法，那么 itable 底部就会有一个列表。

## 计算 Itable

现在我们知道了 itable 的样子，但是这些 itable 到底是如何获得的呢？Go 中的类型可以动态的转化，这也就意味着编译器不可能提前计算所有可能的 itables。因为接口类型和具体类型可以组合成许多可能，但是绝大部分都是无用的。go 编译器选择为每个具体类型(例如`Binary`，`int`或者`func(map[int]string)`)产生一个类型描述的结构。除了一些元数据，类型描述结构也包含此类型实现的所有方法。go 中也为接口(例如`Stringer`)产生一个类似的结构体，这个结构体当然也包括一个方法列表。在运行时，go 会通过在结构体的方法列表中查询结构体的每个方法来计算 itable，runtime 在产生 itable 之后会对齐进行缓存，从而这种关联就只会计算一次了。

在我们的简单的例子中，`Stringer`的方法表中只有一个方法，然而`Binary`的方法表中有两个方法。如果接口类型的方发表中有`ni`个方法，具体的类型有`nt`个方法，那么时间复杂度就会是`O(ni*nt)`，但是我们可以对此进行优化。如果对两个方法表进行排序(译者注: 排序可能是接口或者具体类型的方法表在构建的时候进行的)，然后同时遍历，这样构建动态类型的方法表的时间复杂度就会是`O(ni + nt)`

## 内存优化

上述具体的实现方法中消耗的内存可以有两种方法来进行优化。

首先，如果涉及的接口是一个空接口-没有方法-那么这个 itable 的作用就仅仅是存储一个指向原始类型的指针。在这个例子中就是，itable 可以不要，然后其指针仅仅指向类型即可:

![](https://research.swtch.com/gointer3.png)

其次，如果关联值的大小只有只有一个字的长度的时候，就没有必要使用指针关联了或者把值分配到堆上。如果我们定义一个和`Binary`相似的`Binary32`，但是其实现的类型单位`uint32`，那么接口指向指针的数据就可以更换为具体的值了。

![](http://research.swtch.com/gointer4.png)

实际的值到底是通过指针引用还是直接存储，取决于值类型的大小。编译器可以根据传递进来的值来安排具体的类型来进行正确的调用。如果接受值正好是一个字长，那么可以直接使用；如果不是的，那么可以指针引用。图中展示了，`Binary`版本的 itable 中实际调用的方法是`(*Binary).String`，而在`Binary32`版本中调用的方法是`Bianry32.String`。

当然，空接口中如果值得大小正好是一个字长(或者更短)的时候可以使用上面两种方式进行优化:

![](http://research.swtch.com/gointer5.png)


# go 中的 interfaces

## 前言
本文翻译自:https://research.swtch.com/interfaces ， 部分内容作了调整，也增加了些细节验证。
## 接口值

拥有方法的语言中通常分为两个阵营: 预先准备好静态的方法表(例如 C++ 和 Java)，或者在每次调用的方法的时候进行方法的搜索(例如 Smalltalk 以及其模仿者，包括 JavaScript 和 Python)，再加入缓存提高效率。Go 是处于两者之间的: 它有方法表，但是是在 runtime 期间计算出来的。不知道 Go 是不是编程语言中第一个这样做的，但是它肯定不是唯一一个这样做的。

作为一个预热，一种类型为`Binary`的值是一个 64 位的整型，由两个 32 位组成(假设电脑是一个 32 位的机器)

![](https://user-gold-cdn.xitu.io/2020/7/4/1731914c79f9b957?w=159&h=83&f=png&s=1516)

其拥有的方法，接口`Stringer`定义，以及接受接口的函数`ToString`的定义如下:

```go
type Binary int64

func (i Binary) String() string {
	return strconv.FormatInt(i.Get(), 10)
}

func (i Binary) Get() int64 {
	return int64(i)
}

type Stringer interface {
    String() string
}

func ToString(any interface{}) string {
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    switch v := any.(type) {
    case int:
        return strconv.Itoa(v)
    case float64:
        return fmt.Sprintf("%f", v)
    }
    return "???"
}
```

接口值是由两指针组成，一个指针指向类型相关的信息，一个指针指向其关联的数据。在给变量`b`赋值一个类型为`Stringer`的接口时候，这两个指针的值都会被设置。

![](https://user-gold-cdn.xitu.io/2020/7/4/1731914c69e730c5?w=486&h=238&f=png&s=10238)

(图中接口包含的指针是灰色的，表示它们并不直接暴露给 go 编码者)

interface 中的第一个指针指向的一个 interface table 或者 itable (读成 i-table; 在 [runtime 源码](https://golang.org/src/pkg/runtime/iface.c#L23)，其 C 实现的名称是 `Itab`)。itable 结构中的起始部分是一些例如类型的元数据，然后是一个函数指针的列表。注意：itable 对应的类型是接口类型，而不是动态类型。在我们的例子中，接口`Stringer`的 itable 包含 `Binary`中用来实现`Stringer`的方法，只有`String`方法。`Binary`中的其他方法(例如`Get`)并不会出现在 itable 之中。

interface 中的第二个指针指向的是实际的数据，在这个例子中是`b`复制的数据。赋值语句`var s Stringer = b`会复制一份`b`的值赋给`s`，而不是让`s`的指针指向`b`，正如`var c uint64 = b`一样。如果`b`的值在后面改变了，那么`s`和`c`并不会改变。interface 中的数据可以是任意大的，但是 interface 结构中只有一个字的长度用来存储这个值，所以赋值的时候，会在堆上分配内存，并在 interface 中使用指针来指向堆上的数据。(如果实际值得大小只有一个字的长度，那么这个地方是显然是可以优化的，在后面会讲到)。

为了检测 interface 的类型(和`ToString`方法中的 type switch 示例一样)，Go 编译器会产生和 C 表达式`s.tab->type`一样的代码来类型指针，并对比类型。如果类型一致，那么可以通过`s.data`复制。

为了调用`s.String()`，Go 编译器会产生和 C 表达式一样的`s.tab->fun[0](s.data)`的代码：调用 itable 中合适的方法，并把 interface 中的 value 值作为方法的第一个参数(在这个例子中是一个唯一的参数)。如果你运行`8g -S x.go`(详细的细节在文章末尾)，你就可以看到这个代码。你可以注意到传递给 itable 中的函数的参数是一个 32 位的指针，而不是 64 位的指针实际指向的值。通常，接口调用方法的时候并不知道字长值的含义或者这个指针指向值的大小。interface  代码会会使得 itable 中的函数指针接受一个 32 位的存储在 interface 中的值。所以这个例子中的函数指针实际上是`(*Binary).String`而不是`Binary.String`。

这个例子中的接口只有一个方法。如果一个接口中有多个方法，那么 itable 底部就会有一个列表。

## 计算 Itable

现在我们知道了 itable 的样子，但是这些 itable 到底是如何获得的呢？Go 中的类型可以动态的转化，这也就意味着编译器不可能提前计算所有可能的 itables。因为接口类型和具体类型可以组合成许多可能，但是绝大部分都是无用的。go 编译器选择为每个具体类型(例如`Binary`，`int`或者`func(map[int]string)`)产生一个类型描述的结构。除了一些元数据，类型描述结构也包含此类型实现的所有方法。go 中也为接口(例如`Stringer`)产生一个类似的结构体，这个结构体当然也包括一个方法列表。在运行时，go 会通过在结构体的方法列表中查询结构体的每个方法来计算 itable，runtime 在产生 itable 之后会对齐进行缓存，从而这种关联就只会计算一次了。

在我们的简单的例子中，`Stringer`的方法表中只有一个方法，然而`Binary`的方法表中有两个方法。如果接口类型的方发表中有`ni`个方法，具体的类型有`nt`个方法，那么时间复杂度就会是`O(ni*nt)`，但是我们可以对此进行优化。如果对两个方法表进行排序(译者注: 排序可能是接口或者具体类型的方法表在构建的时候进行的)，然后同时遍历，这样构建动态类型的方法表的时间复杂度就会是`O(ni + nt)`

## 内存优化

上述具体的实现方法中消耗的内存可以有两种方法来进行优化。

首先，如果涉及的接口是一个空接口(没有方法)那么这个 itable 的作用就仅仅是存储一个指向原始类型的指针。在这个例子中就是，itable 可以不要，然后其指针仅仅指向类型即可:

![](https://user-gold-cdn.xitu.io/2020/7/4/1731914c80f6aced?w=317&h=170&f=png&s=5351)

其次，如果关联值的大小只有只有一个字的长度的时候，就没有必要使用指针关联了或者把值分配到堆上。如果我们定义一个和`Binary`相似的`Binary32`，但是其实现的类型单位`uint32`，那么接口指向指针的数据就可以更换为具体的值了。

![](https://user-gold-cdn.xitu.io/2020/7/4/1731914c70ea3b8e?w=418&h=238&f=png&s=9300)

实际的值到底是通过指针引用还是直接存储，取决于值类型的大小。编译器可以根据传递进来的值来安排具体的类型来进行正确的调用。如果接受值正好是一个字长，那么可以直接使用；如果不是的，那么可以指针引用。图中展示了，`Binary`版本的 itable 中实际调用的方法是`(*Binary).String`，而在`Binary32`版本中调用的方法是`Bianry32.String`。

当然，空接口中如果值得大小正好是一个字长(或者更短)的时候可以使用上面两种方式进行优化:

![](https://user-gold-cdn.xitu.io/2020/7/4/1731914c7023a06c?w=299&h=83&f=png&s=4132)

## 代码展示

Selected output of `8g -S x.go`(其实我也看不懂，，，，，):

```shell
0045 (x.go:25) LEAL    s+-24(SP),BX
0046 (x.go:25) MOVL    4(BX),BP
0047 (x.go:25) MOVL    BP,(SP)
0048 (x.go:25) MOVL    (BX),BX
0049 (x.go:25) MOVL    20(BX),BX
0050 (x.go:25) CALL    ,BX
```

## 疑问
关于字长的优化，有点疑虑。可能作者针对的版本比较老，所以和现在版本的数据有些不一致。
```
type Binary bool

func (i Binary) String() string {
	return strconv.FormatInt(i.Get(), 10)
}

func (i Binary) Get() int64 {
	return int64(100)
}

type Stringer interface {
	String() string
}

type emptyInterface struct {
	typ unsafe.Pointer
	val unsafe.Pointer
}

func Change2interface() {
	var n Binary = false
	var nS Stringer = n

	ei := (*emptyInterface)(unsafe.Pointer(&nS))
	_type, _value := ei.typ, ei.val
	fmt.Println(*(*bool)(_type))
	fmt.Println(*(*bool)(_value))

}
// Output: true
// false
```
如果是一个`int32`类型的，结果也是一致的，并没有进行优化(至于我为什么使用`fmt.Println`打印这么多东西，是为了表明验证的方法是没有问题的)
```

type Binary int32

func (i Binary) String() string {
	return strconv.FormatInt(i.Get(), 10)
}

func (i Binary) Get() int64 {
	return int64(i)
}

type Stringer interface {
	String() string
}

type emptyInterface struct {
	typ unsafe.Pointer
	val unsafe.Pointer
}

func Change2interface() {
	var n Binary = 100
	var nS Stringer = n

	ei := (*emptyInterface)(unsafe.Pointer(&nS))
	_type, _value := ei.typ, ei.val
	fmt.Println(*(*int32)(_type))
	fmt.Println(uintptr(_type))
	fmt.Println(*(*int32)(_value))
	fmt.Println(uintptr(_value))
}
// Output: 17914976
// 18214720
// 100
// 824633811632
```
由于类型`Binary`是一个长度不足一个字长的字符(这个测试特意写的是一个`bool`类型，我的机子是一个 64 位的系统，大家用`int32` 类型也可以做一样的测试)，但是测试的结果中指向值的指针并没有被优化为一个值。
虽然文章的内容的部分细节可能过时，但是整体的思想还是 go 中接口实现的思想，所以还是值得阅读的。

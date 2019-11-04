#Go切片：使用方式以及组成

## 介绍

**本文翻译自：https://blog.golang.org/go-slices-usage-and-internals**

Go中的切片提供了一种方便、有效的处理一系列特定类型值的方式。切片在其他语言中和数组是相似的，但是有一些不同的特性。这篇文章将会讨论切片，以及如何使用它们。

## 数组

在go中切片是建立于数组之上的，所以在理解切片之前，我们必须先理解数组。

数组在定义的时候需要明确长度和元素的类型。例如，类型`[4]int`表示的是一个长度为4的数组。数组的长度是固定的，长度是数组类型的一部分。也就是说`[4]int`和`[5]int`是两个截然不同的类型。数组通过索引的方式取值，表达式`s[n]`用于获取数组的第n个元素，数组的索引是从0开始的。

数组不用明显的初始化(译者注:不用每个位置的值都初始化)，数组的元素的初始值都是该数组对应类型的零值。

```go
// a[2] == 0, the zero value of the int type
```

内存中`[4]int`是连续分布的四个整型值。

![](https://blog.golang.org/go-slices-usage-and-internals_slice-array.png)

go中的数组是值传递的。一个数组变量表示的是整个数组；而不是像在C语言中一样是一个指向数组第一个元素的指针。这意味着当你赋值或者传递数组的时候，你使用的是数组的一个copy。(你可以通过传递数组指针的方式来进行这种操作)。我们可以认为数组是一个通过索引而不是使属性取值的结构体:一个长度固定的符合类型元素。

你可以通过如下方式定义一个数组

```go
b := [2]string{"Penn", "Teller"}
```

或者，你也可以让编译器来为你计算数组的长度

```go
b := [...]string{"Penn", "Teller"}
```

上面两个case中，数组的类型都是`[2]string`

## 切片

go中数组是由存在的意义的，但是由于有一点僵化(大小固定)，所以你很少在Go的代码中见到它们。然而，到处都可以见到切片。切片是基于数组构建的，为开发者提供了巨大的能力和方便。

切片通过`[]T`来定义，其中`T`代表的是切片中元素的类型。不需要和数组一样，切片并不明确的定义一个长度。

切片在字面上的声明和相似，除了你不需要说明长度:

```go
letters := []string{"a", "b", "c", "d"}
```

切片也可以通过go内置的函数`make`来进行创建，其签名(函数定义的方式？？？)如下,

```go
func make([]T, len, cap) []T
```

其中`T`表示切片中元素的类型。函数`make`接受的参数包括类型(type)、长度(len)以及可选的容量(cap)。当make被如此调用的时候，就会分配一个数组并返回表示此数组的切片。

```go
var s []byte
s = make([]byte, 5, 5)
// s == []byte{0, 0, 0, 0, 0}
```

当容量(cap)参数省略的时候，其默认值和长度一样。下面是和上面一致的更简明的代码:

```go
s := make([]byte, 5)
```

切片的长度和容量可以通过内置函数`len`和`cap`分别进行检查

```go
len(s) == 5
cap(s) == 5
```

下面的两部分内容会用于讨论长度和容量之间的联系。

切片的零值是`nil`,在切片为零值的时候，调用`len`和`cap`函数都会返回0.

切片也可以通过对切片或者数组进行`切片`操作来创建。切片操作通过一个由冒号分隔的半开区间(左边包括，右边不包括)完成的。例如，操作表达式`b[1:4]`创建切片包含b中从1到3的元素(返回的元素的切片范围为从0到2)。

```go
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
```

切片表达式的初始位置和结束位置都可以可以省略的，初始值默认为0，结束值默认是切片(或数组)的长度。

```go
// b[:2] == []byte{'g', 'o'}
// b[2:] == []byte{'l', 'a', 'n', 'g'}
// b[:] == b
```

下面是通过数组的方式来创建切片：

```go
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // a slice referencing the storage of x
```

## Slice的内部结构

切片是一个用于表述数组片段的结构。它包含一个指向数组的指针，片段的长度，以及容量(片段的最大长度)。

![](https://blog.golang.org/go-slices-usage-and-internals_slice-struct.png)

我们之前通过`make([]byte, 5)`创建的变量s，其内部是按照如下方式组织的：

![](https://blog.golang.org/go-slices-usage-and-internals_slice-1.png)

长度表示的是slice所代表的元素的个数。容量表示的是切片所在的数组的元素的个数(起始位置为切片表示)。长度和容量的区别会随着下面的几个例子的出现愈加清晰。

当我们对`s`进行切片操作，可以观察到切片数据结构的变化，以及和底层数组的关系：

```go
s = s[2:4]
```

![](https://blog.golang.org/go-slices-usage-and-internals_slice-2.png)

切片操作并不会复制切片的值。它会创建一个新的切片用于指向原来的数组。这样就会使得切片的操作和通过索引对数组取值一样搞笑。因此，修改切片中元素的值，也会修改之前切片中的值。因为他们共享一个数组的空间。

```go
d := []byte{'r', 'o', 'a', 'd'}
e := d[2:] 
// e == []byte{'a', 'd'}
e[1] = 'm'
// e == []byte{'a', 'm'}
// d == []byte{'r', 'o', 'a', 'm'}g
```

在之前，我们创建的切片的长度小于底层数组的长度，我们可以再次通过切片的操作来延长`s`：

```go
s = s[:cap(s)]
```

![](https://blog.golang.org/go-slices-usage-and-internals_slice-3.png)

切片并不能增长到超过其容量。尝试进行这种操作会造成一个运行时异常，和对数组或者切片通过索引取值，而索引超过范围时的异常是一样的。通过对切片进行小于0的切片来获取数组之前的元素也会造成同样的异常。

##增长的切片(copy和append函数)

为了增长一个切片的容量，我们必须创建一个新的并且容量更大的切片并且将原来切片中的数据复制到新的切片之中。这个就是其他语言动态数组背后进行的操作。下面这个例子就会通过创建一个新的切片`t`，然后把切片`s`中的数据复制到`t`之中，再把`t`赋值给`s`来对`s`进行容量翻倍的操作：

```go
t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
for i := range s {
        t[i] = s[i]
}
s = t
```

像代码中的for循环来赋值的操作可以通过Go的内置函数`copy`来操作。正如名称所表明的一样，`copy`从原始的切片中复制元素到目标切片之中，函数返回复制的元素的个数

```go
func copy(dst, src []T) int
```

`copy`函数支持两个不同长度的切片进行复制操作(仅支持复制长度较小的切片的长度元素)。另外，`copy`可以正确的处理原始切片和目标切片共享一个底层数组的情况。(这个可以试试)

使用`copy`，我们可以简化上面的翻倍切片的代码:

```go
t := make([]byte, len(s), (cap(s)+1)*2)
copy(t, s)
s = t
```

一个常用的操作就是给一个切片的末尾添加一个元素。下面这个函数就会在一个切片的末尾添加byte元素，并且在必要的时候增长切片的容量，返回一个更新了的切片值：

```go
func AppendByte(slice []byte, data ...byte) []byte {
    m := len(slice)
    n := m + len(data)
    if n > cap(slice) { // if necessary, reallocate
        // allocate double what's needed, for future growth.
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:n]
    copy(slice[m:n], data)
    return slice
}
```

我们可以像下面这样使用`AppendByte`函数：

```go
p := []byte{2, 3, 5}
p = AppendByte(p, 7, 11, 13)
// p == []byte{2, 3, 5, 7, 11, 13}
```

像`AppendByte`这样的函数是非常有用的，它们在切片增长的时候提供了完全的控制。根据程序的特性，可能需要分配更小的或者更大的块，或者给出再次分配的最大值。

但是大部分程序都不需要这样的完全的控制，所以Go提供了一个内置的`append`函数；函数的签名如下

```go
func append(s []T, x ...T) []T
```

函数`append`将元素`x`添加到切片`s`的末尾，在需要更多的容量的时候会增大切片的容量。

```go
a := make([]int, 1)
// a == []int{0}
a = append(a, 1, 2, 3)
// a == []int{0, 1, 2, 3}
```

我们可以通过`…`的操作将一个切片的所有元素添加到另一个切片的尾部。

```go
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// a == []string{"John", "Paul", "George", "Ringo", "Pete"}
```

由于切片的零值是`nil`，表现的像一个长度为0的切片，所以你可以声明一个切片变量，然后在一个for循环中给其添加元素:

```go
// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}
```

## 一个可能的坑

正如之前提到过的一样，再次切片并不会复制底层数组的元素。底层的数组会一直保存在内存中，直到没有变量引用到此数组为止。这会偶尔的造成程序把全部数据保存在内存中，而使用的仅仅是其中的一小部分。

例如，函数`FindDigits`会把一个文件加载到内存之中，然后在里面搜索第一个连续的数字，并把它们通过切片的形式返回

```go
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    return digitRegexp.Find(b)
}
```

这个代码就如之前说到的一样，会返回一个`[]byte`指向一个包含整个文件的数组。由于切片指向原始的数组，所以数组会一直保存在内存中，并且垃圾回收并不会释放其内存；仅仅使用的部分bytes，造成整个文件保存在内存中。

为了解决这个问题，我们可以将需要的值赋值到新创建的切片之中然后返回这个切片即可：

```go
func CopyDigits(filename string) []byte {
    b, _ := ioutil.ReadFile(filename)
    b = digitRegexp.Find(b)
    c := make([]byte, len(b))
    copy(c, b)
    return c
}
```


# JavaScript的作用域

## 前言

js中的一个常见异常是ReferenceError，表示引用的变量不存在

js其实不像大多数认为的那样是一个脚本语言，他是一个编译语言。代码中包括变量和函数在内的所有声明都会在任何代码被执行前首先被处理

## 编译语言

### 变量提升

``` javascript
console.log(a)
a = 2
console.log(a)
var a 
```

上述代码是不会报错的。只不过console.log(a)的时候，打印的是undefined。由于是编译语言，所以`var a`被提升，在代码被执行之前被处理了。在第一次执行`console.log(a)`的时候，变量是存在的，只不过没有被赋值，所以是undefined。第二次执行的时候就会打印出2了。此时的运行时的结果如下

```
➜  ~ node test.js
undefined
2
```

如果把代码中变换如下

```javascript
console.log(a)
a  = 2
```

就会报错了，出现ReferenceError的错误。应为`console.log(a)`运行的时候，变量a并不存在。运行时的结果如下

```
➜  ~ node test.js
/Users/.../test.js:1
console.log(a)
            ^
ReferenceError: a is not defined
    at Object.<anonymous> (/Users/.../test.js:1:13)
    at Module._compile (internal/modules/cjs/loader.js:759:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:770:10)
    at Module.load (internal/modules/cjs/loader.js:628:32)
    at Function.Module._load (internal/modules/cjs/loader.js:555:12)
    at Function.Module.runMain (internal/modules/cjs/loader.js:824:10)
    at internal/main/run_main_module.js:17:11
```

### 函数提升

变量被提升的时候，可能很少人知道。但是函数提升，大家就一定都清楚了。因为在代码中的位置，函数调用是可以在函数声明之前的。如下，此时`foo()`是可以正常调用的。

``` javascript
foo()
function foo() {
	console.log(2)
}
```

虽然函数是可以提升的，但是函数表达式不会提升，这点是需要注意的。例如，如下代码就会报错，但是错误并不是ReferenceError，而是TypeError。

```javascript
foo()
var foo = function() {
	console.log(1)
}
```

原因是在执行`foo()`这行代码的时候，foo变量是存在的。而此时还没有被赋值为函数，对于一个undefined的变量，当做函数来执行，报错当然是TypeError了。详细错误如下

```
/Users/.../test.js:1
foo()
^
TypeError: foo is not a function
    at Object.<anonymous> (/Users/.../test.js:1:1)
    at Module._compile (internal/modules/cjs/loader.js:759:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:770:10)
    at Module.load (internal/modules/cjs/loader.js:628:32)
    at Function.Module._load (internal/modules/cjs/loader.js:555:12)
    at Function.Module.runMain (internal/modules/cjs/loader.js:824:10)
    at internal/main/run_main_module.js:17:11
```

### 变量和函数同名

```javascript
foo()
var foo
function foo() {
	console.log(1)
}
foo = function() {
	console.log(2)
}
```

上述代码的输出结果会是什么呢？是出现异常TypeError，还是打印1，还是打印2？

首先确定可以排除的是打印2，因为`foo()`的调用是在`foo = function() {console.log(2)}`之前执行的。

剩下的就判断不了了。但是，结果是会打印1，而不会触发TypeError这个错误。

**<font color=red>这是由于函数被提升的优先级要高于变量被提升的优先级。</font>**从代码中可以看出来，虽然`var foo`在`function foo() {…}`之前被声明，但是`foo()`被执行的时候，foo还是代表的是函数的定义。(这个不是因为函数后声明，所以把之前变量foo的声明给覆盖了，换个位置结果是一样的，读者可以试试)

## 块作用域

下面的代码，大家都会很熟悉。

```javascript
for (var i=0; i<10; i++) {
	console.log(i)
}
```

这个会从0到9的打印10个数字。但是，也会造成一个问题，就是之后的代码中也是可以调用变量i的。在之后执行命令`console.log(i)`就会打印出10了。这个可能不是我们想要的。因为在for循环中，我们希望的是i仅仅在这个for循环之中使用的，这个在JavaScript之中是可以简单有效的解决的。

在ES6之中，提供了除了var之外声明变量的方式，有let和const。let关键字可以将变量绑定到所在的任意作用域中(通常是{..}内部)。换句话说，let为其声明的变量的隐饰的劫持在了所在的块作用域之中。const声明的变量也是属于块作用域的，只是const声明的变量不能再次被赋值。<font color=red>for循环头部的let声明还会有一个特殊的行为，这个行为为指出变量在循环的过程中不止被声明一次，每次迭代都会声明。声明后的每个迭代都会使用上一个迭代结束时的值来初始化这个值</font>

```javascript
for (let i=0; i<5; i++) {
	console.log(i)
}
console.log(i)
```

输出如下

```
➜  ~ node test.js
0
1
2
3
4
/Users/.../test.js:4
console.log(i)
            ^
ReferenceError: i is not defined
    at Object.<anonymous> (/Users/.../test.js:4:13)
    at Module._compile (internal/modules/cjs/loader.js:759:30)
    at Object.Module._extensions..js (internal/modules/cjs/loader.js:770:10)
    at Module.load (internal/modules/cjs/loader.js:628:32)
    at Function.Module._load (internal/modules/cjs/loader.js:555:12)
    at Function.Module.runMain (internal/modules/cjs/loader.js:824:10)
    at internal/main/run_main_module.js:17:11
```

除了会顺序输出0到4的五个数外，还会报错ReferenceError。这个就是let的作用。

### 题外话

python的作用域也挺有意思的，下面是做python开发经常用来的面试题(可是我从来没遇到过)。

```python
a = [lambda : x for x in range(3)]
for func in a:
    print func()
    
2
2
2
```

三个输出竟然都是2，这是为什么呢？是由于三个函数在执行的时候，由于函数内部没有变量x，所以通过legb原则逐层向上搜索变量x，而在执行的时候的x为2，所以输出的都是2。

那么如何修改为输出为0，1，2呢？可以给匿名函数传递个默认参数

```python
a = [lambda x = x: x for x in range(3)]
for func in a:
    print func()
    
0
1
2
```

JavaScript之中，也会遇到相似的问题。

```javascript
for (var i=0; i<3; i++) {
       setTimeout( function time() {
              console.log(i)
        }, 1000)
}
```

一秒之后的输出结果是3个3。那么如何修改为分别输出0，1，2呢？

1. 通过和python相似的方法

   ```javascript
   for (var i=0; i<3; i++) {
           (function (i){setTimeout( function time() {
                   console.log(i)
           }, 1000)})(i)
   }
   ```

2. 通过JavaScript的块作用域方法

   ```javascript
   for (let i=0; i<3; i++) {
           setTimeout(function() {console.log(i)}, 1000);
   }
   ```

上述两个方法都是可以正确的输出的

```
➜  ~ node test.js
0
1
2
```


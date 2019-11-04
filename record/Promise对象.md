# Promise对象和任务队列

## Promise对象的两个特点

1. 对象的状态不受外界影响。Promise对象代表一个异步操作，有三种状态：`pending(进行中)`、`fulfilled(已成功)`和`rejected(已失败)`
2. 一旦状态改变，就不会再变。状态的改变有两种可能：`从pending到fulfilled和从pending变为rejected。`

## 优点

有了`Promise`对象，就可以将异步操作以同步操作的流程表达出来，避免了层层嵌套的回调函数。此外，`Promise`对象提供统一的接口，使得控制异步操作更加容易。

## 基本用法

下面是一个Promise实例

```javascript
const promise = new Promise(function(resolve, reject) {
  // ... some code
  if (/* 异步操作成功 */){
    resolve(value);
  } else {
    reject(error);
  }
});
```

Promise构造函数接受一个函数作为参数，该函数的两个参数分别为resolve和reject。resolve函数的作用是将Promise对象的状态由pending转化为fulfilled。reject函数的作用是将Promise对象的状态由pending转化为rejected。

Promise实例生成之后，可以使用then方法分别指定resolved状态和rejected状态的回调函数，其中rejected状态的回调函数是可选的。

```javascript
promise.then(function(value) {
  // success
}, function(error) {
  // failure
});
```

还有下面一种更为常见的方式，网络请求就常用此方式。此方式和上述方式是等价的。

```javascript
promise.then(function(value) {
  // success
}).catch(function(error) {
  // failure
});)
```

下面是一个Promise对象的简单例子。

```javascript
function timeout(ms) {
  return new Promise((resolve, reject) => {
    setTimeout(resolve, ms, 'done');
  });
}

timeout(100).then((value) => {
  console.log(value);
});
```

在知道上面函数调用的过程之前，另外需要知道的就是`setTimeout`函数的调用方式，`用于在指定的毫秒数后调用函数或计算表达式，还可以给指定调用的函数传递参数`。例如代码`setTimeout((a) => {console.log(a)}, 100, 'hello world')`就会在0.1秒之后打印`hello world`。

这样就可以理解上述的Promise对象例子的用法了。在返回的Promise对象状态由pending状态到fulfilled状态之后，会执行传入的resolve函数。setTimeout中传递的第一个参数`resolve`就是用于回调的函数，第二个参数`ms`就是在`ms`毫秒之后执行`resolve`函数，第三个参数`'done'`就是传递给resolve函数调用时的参数。timeout(100)返回的Promise对象在状态由`pending`转化为`fulfilled`的时候会调用传入的匿名函数，该函数仅会打印传入的参数值，由于传入的参数为字符串`'done'`，所以会打印`'done'`。

## 注意

Promise建立之后会立即执行。

```javascript
let promise = new Promise(function(resolve, reject) {
  console.log('Promise');
  resolve();
});

promise.then(function() {
  console.log('resolved.');
});

console.log('Hi!');
// 打印结果顺序如下:
// Promise
// Hi!
// resolved
```

上述代码中，Promise新建之后会立即执行，所以首先输出的是`Promise`，但是then指定的回调函数会在最后将在当前脚本所有同步任务执行完才会执行，所以`resolved`最后输出。原因可在之后的Event Loop中找到答案。

## 任务队列

JavaScript中所有的任务可以分成两种，一种是`同步任务(synchronous)`，一种是`异步任务(asynchronous)`。同步任务指的是在主线程上排队执行的任务，只有前一个任务执行完毕，才能执行后一个任务；异步任务指的是，不进入主线程，而进步`"任务队列"(task queue)`的任务，只有`"任务队列"`通知主线程，某个异步任务可以执行了，该任务才会进入主线程执行。

具体来说，异步执行的运行机制如下。（同步执行也是如此，因为它可以被视为没有异步任务的异步执行。

> （1）所有同步任务都在主线程上执行，形成一个[执行栈](http://www.ruanyifeng.com/blog/2013/11/stack.html)（execution context stack）。
>
> （2）主线程之外，还存在一个"任务队列"（task queue）。只要异步任务有了运行结果，就在"任务队列"之中放置一个事件。
>
> （3）一旦"执行栈"中的所有同步任务执行完毕，系统就会读取"任务队列"，看看里面有哪些事件。那些对应的异步任务，于是结束等待状态，进入执行栈，开始执行。
>
> （4）主线程不断重复上面的第三步。

下图就是主线程和任务队列的示意图。

![avatar](http://www.ruanyifeng.com/blogimg/asset/2014/bg2014100801.jpg)

只要主线程空了，就会去读取"任务队列"，这就是JavaScript的运行机制。这个过程会不断重复。

## 事件和回调函数

"任务队列"是一个事件的队列（也可以理解成消息的队列），IO设备完成一项任务，就在"任务队列"中添加一个事件，表示相关的异步任务可以进入"执行栈"了。主线程读取"任务队列"，就是读取里面有哪些事件。

"任务队列"中的事件，除了IO设备的事件以外，还包括一些用户产生的事件（比如鼠标点击、页面滚动等等）。只要指定过回调函数，这些事件发生时就会进入"任务队列"，等待主线程读取。

所谓"回调函数"（callback），就是那些会被主线程挂起来的代码。异步任务必须指定回调函数，当主线程开始执行异步任务，就是执行对应的回调函数。

"任务队列"是一个先进先出的数据结构，排在前面的事件，优先被主线程读取。主线程的读取过程基本上是自动的，只要执行栈一清空，"任务队列"上第一位的事件就自动进入主线程。但是，由于存在后文提到的"定时器"功能，主线程首先要检查一下执行时间，某些事件只有到了规定的时间，才能返回主线程。

## Event Loop

主线程从"任务队列"中读取事件，这个过程是循环不断的，所以整个的这种运行机制又称为Event Loop（事件循环）。

![avatar](http://www.ruanyifeng.com/blogimg/asset/2014/bg2014100802.png)

`heap`中存放对象，`stack`中存放当前运行的代码。dom，ajax和setTimeout(setInterval一样)之类的操作会交给`WebAPIs`执行，在执行完成之后放到`callback queue(和上述的任务队列一致)`。在`stack`中代码全部执行完成之后，会查看`callback queue`中是否有任务需要执行。如此反复的进行就是事件循环。图是截取自[名为《Help, I'm stuck in an event-loop》的演讲](http://www.ruanyifeng.com/blogimg/asset/2014/bg2014100802.png)。

上述关于队列的介绍可以用一个简单的例子来说明

```javascript
setTimeout(function() {console.log(1)}, 0)
console.log(2)
// 打印顺序如下：
// 2
// 1
```

setTimeout设置的定时时间为0ms，即在0ms后打印，按照'正常的'步骤应该是立即打印1，然后再打印2：但是实际的打印结果中的顺序是先打印2，然后再打印1。这是由于在执行setTimeout的时候，会交给`WebAPIs`定时，定时完成之后交给`callback queue`，而执行`stack`会继续往下执行，在执行完`console.log(2)`之后，才会查看在`callback queue`中的任务并执行。由于这个定时任务定时0ms后执行，在咱看似没有必要，却仍然会按照既定的步骤执行。


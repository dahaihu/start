# JavaScript异步编程: 从回调地狱到async和await

写好一个优秀的web应用关键之一就是可以在一个页面上做许多AJAX请求

## 前言

我们可以从一个简单的例子的每个解决方式来思考JavaScript异步编程的进步

为了做到这些，我们可以来做一个简单的任务，这个任务是完成下面这些流程：

1. 验证用户的名称和密码
2. 获取应用中用户的角色
3. 打印用户访问应用的时间

## 回调地狱的方式

最古老的解决这些问题的方式是通过一层套一层的的回调。这在过去是解决简单的异步任务的优雅方式，但是呢，由于回调地狱的原因，并不能进行拓展升级。

![](https://miro.medium.com/max/1281/0*dcS--RdIWbccmWXE.png)

用于解决三个简单问题的代码如下

```javascript
const verifyUser = function(username, password, callback){
   dataBase.verifyUser(username, password, (error, userInfo) => {
       if (error) {
           callback(error)
       }else{
           dataBase.getRoles(username, (error, roles) => {
               if (error){
                   callback(error)
               }else {
                   dataBase.logAccess(username, (error) => {
                       if (error){
                           callback(error);
                       }else{
                           callback(null, userInfo, roles);
                       }
                   })
               }
           })
       }
   })
};
```

每个函数的调用都需要传递一个参数，这个参数也是一个函数，会接受前一个函数的返回值。

仅仅阅读上面的句子都会使得太多人大脑麻木，而如果一个应用拥有数以百计的这种代码的话，会对维护这些代码的人造成极大的困扰，即使是这些人自己写的这些代码(不要高估明天的自己，明天的你不一定可以看懂今天的你写的代码)。

当你知道database.getRoles`也是一个嵌套回调的函数的时候，你会意识到这个例子变得更加复杂。

```javascript
const getRoles = function (username, callback){
   database.connect((connection) => {
       connection.query('get roles sql', (result) => {
           callback(null, result);
       })
   });
};
```

这些代码的问题，除了代码难以维护以外，还有就是违背DRY原则(Do not repeat yourself)。例如，异常处理在每个函数中都被重复的进行(if else)并且调用的callback也在每个嵌套的function中调用。

## JavaScript中的Promises

`Promise`是避免回调地狱的一个进步。此方法并没有移除回调函数，但是将函数的调用连接起来并且简化了代码，使得代码更易于阅读。

![](https://miro.medium.com/max/1281/0*_if1EyoEM4I4jwpB.png)

使用了`Promise`的代码如下

```javascript
const verifyUser = function(username, password) {
   database.verifyUser(username, password)
       .then(userInfo => dataBase.getRoles(userInfo))
       .then(rolesInfo => dataBase.logAccess(rolesInfo))
       .then(finalResult => {
           //do whatever the 'callback' would do
       })
       .catch((err) => {
           //do whatever the error handler needs
       });
};
```

为了做到这些，代码中使用到的函数必须Promise化。我们看一下`getRoles`如何更新的返回一个Promise对象的。

```javascript
const getRoles = function (username){
   return new Promise((resolve, reject) => {
       database.connect((connection) => {
           connection.query('get roles sql', (result) => {
               resolve(result);
           })
       });
   });
};
```

我们将此方法修改为返回一个`Promise`对象，此对象需要传入两个回调函数，并且在Promise对象中执行操作。现在，`resolve`和`reject`两个回调函数会分别映射为`Promise.then`和`Promise.catch`。

你可能意识到`getRoles`方法的内部仍然是可能遭受地狱回调的，因为database的方法并没有返回一个Promise对象。如果database的连接方法也返回一个Promise对象，那么`getRoles`就会像下面这样

```javascript
const getRoles = new function (userInfo) {
   return new Promise((resolve, reject) => {
       database.connect()
           .then((connection) => connection.query('get roles sql'))
           .then((result) => resolve(result))
           .catch(reject)
   });
};
```

## 方式3：Async/Await

JavaScript默认是异步的。这也可能是为什么JavaScript花了很长时间才使得代码看起来像同步的。但是，迟到了总比没有强。地狱回调在引入了Promise对象之后改善了很多。但是我们仍然需要传递回调函数给Promise对象`.then`和`.catch`。

Promise给JavaScript带来了最酷之一的改变。[ECMAScript 2017](https://en.wikipedia.org/wiki/ECMAScript#8th_Edition_-_ECMAScript_2017)在Promise之上已async和await表达式的方式带来了语法糖。这些使得我们可以基于Promise的代码看起来不像是异步的一样，并且也不会阻塞主线程。

代码如下：

```javascript
const verifyUser = async function(username, password){
   try {
       const userInfo = await dataBase.verifyUser(username, password);
       const rolesInfo = await dataBase.getRoles(userInfo);
       const logStatus = await dataBase.logAccess(userInfo);
       return userInfo;
   }catch (e){
       //handle errors as needed
   }
};
```

`await promise`的操作仅仅允许在`async`修饰的函数中使用，`async`修饰的函数表示verifyUser必须定义为`异步函数`

然而，这些简单的改变就可以`await`任何`Promise`，并且不需要修改任何代码。并且异步代码可以想同步代码一样愉快的编码了。

## Async 一个等待已久的Promise的进步

异步函数是JavaScript中异步编程的下一个里程碑。他们会使得代码更为干净并且更为简单的维护。声明一个函数为async会确保函数返回的是一个Promise对象(即不需要单独的定义返回的对象为一个`Promise`对象)，所以你并不需要为此再担心了。

我们如今为什么需要使用JavaScript中的`async`函数？

1. 代码会更为整洁
2. 异常处理更为简单，仅仅像在其他异步代码使用`try/catch`即可
3. 调试更为简单。在`.then`代码块设置断点不会移动到下一个断点，`.then`只会执行同步代码。但是调试`async`函数就和调试同步代码一样

**本文翻译自https://blog.hellojs.org/asynchronous-javascript-from-callback-hell-to-async-and-await-9b9ceb63c8e8**


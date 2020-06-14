# context 源码解析

## 基本的类型

首先来看的是`Context`到底是什么？源码中的定义是一个接口，有四个方法。

```go
type Context interface {
   Deadline() (deadline time.Time, ok bool)
   Done() <-chan struct{}
   Err() error
   Value(key interface{}) interface{}
}
```

四个方法中除了`Err()`是一个通用的方法外，其他三个方法都各对应一种`Context`类型。也就是说`Context`的实际结构体类型主要有三种(后面会对为什么说主要做解释)。

通常，在初始化一个`Context`的时候，我们都会使用一个方法来创建一个初始化的参数`context.Background()`。

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
   return
}

func (*emptyCtx) Done() <-chan struct{} {
   return nil
}

func (*emptyCtx) Err() error {
   return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
   return nil
}

func (e *emptyCtx) String() string {
   switch e {
   case background:
      return "context.Background"
   case todo:
      return "context.TODO"
   }
   return "unknown empty Context"
}

var (
   background = new(emptyCtx)
   todo       = new(emptyCtx)
)

func Background() Context {
   return background
}

func TODO() Context {
   return todo
}
```

可以看出来，`background`其实是一个`*emptyCtx`，并且实现了`Context`的四个方法。文档中对`Background`方法的注释就说了: 这个`Context`不会被取消，没有值，也没有截止时间。通常的使用方式有如下四种

1. `main`函数
2. `Context`初始化的时候
3. 测试用例
4. 作为处理请求的顶层的`Context`

在服务处理一个请求的时候，各种操作依照依赖的顺序和执行的顺序可以组成一个树状的结构，由根节点像外扩散。为了做好整体的控制，在超时或者某些条件下，后续的操作就不用执行了，这个就需要用户自己实现。而各种情况下的自己实现，是比较耗费时间以及精力的。于是`context`包就诞生了。

其实`context`就是为了在树状的结构中，控制请求在没有必要的时候不再执行。也就是说到了没有必要的时候，我就需要让树装的结构中，某个子树下的所有操作都取消。所以取消操作是`Context`的根本操作。取消操作的结构体定义如下:

```go
type cancelCtx struct {
   Context

   mu       sync.Mutex            // protects following fields
   done     chan struct{}         // created lazily, closed by first cancel call
   children map[canceler]struct{} // set to nil by the first cancel call
   err      error                 // set to non-nil by the first cancel call
}
```

这个里面的`Context`是比较有意思的，可以是四种`Context`中的任何一个，还可以是`emptyCtx`，然后就可以任意的组合。其中的`done`表示是否已经取消，`children`表示此`Context`下的子`Context`，从而控制子树下所有的操作都取消。

## 操作流程

### 初始化

初始化一个`cancelCtx`的步骤通常如下

```go
ctx, cancel := context.WithCancel(context.Background())
```

通过查看源码，可以看到`WithCancel`的操作如下

```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
   c := newCancelCtx(parent)
   propagateCancel(parent, &c)
   return &c, func() { c.cancel(true, Canceled) }
}

func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}
// propagateCancel arranges for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
	if parent.Done() == nil {
		return // parent is never canceled
	}
	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			// parent has already been canceled
			child.cancel(false, p.err)
		} else {
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
		// 什么时候会处于这个状态呢？？？
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}
```

`newCancelCtx`比较简单，初始化了一个`cancelCtx`。复杂点的操作是`propagateCancel`，这个操作就和原文的注释一样了，为了在父`Context`取消时候，此`Context`也可以进行取消的操作。而传入的`parent`可能并不是一个`cancelCtx`的类型，所以需要不断的往父节点寻找，直到找到一个`cancelCtx`类型的`Context`。这个过程需要调用方法`parentCancelCtx`，对其类型的判断有三种`cancelCtx`，`timerCtx`以及`valueCtx`。这个函数应该是不会返回`false`的，因为在调用此函数之前，判断了`parent.Done() == nil `，如果成立，则说明是`background`或者`todo`。如果不成立，则实现了`Context`的结构体类型只有这三种了。

```go
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
   for {
      switch c := parent.(type) {
      case *cancelCtx:
         return c, true
      case *timerCtx:
         return &c.cancelCtx, true
      case *valueCtx:
         parent = c.Context
      default:
         return nil, false
      }
   }
}
```

个人认为，`propagateCancel`函数中`p, ok := parentCancelCtx(parent)`中的`ok`不会是`false`的，并且`parentCancelCtx`的返回值也不会是`false`的。至于这种此时不可能到达的代码，我猜测是为了使得判断的最为完备。

### 取消

`WithCancel`还返回了一个`cancel`的函数，这个函数的作用是什么呢？

```go
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
   if err == nil {
      panic("context: internal error: missing cancel error")
   }
   c.mu.Lock()
   if c.err != nil {
      c.mu.Unlock()
      return // already canceled
   }
   c.err = err
   if c.done == nil {
      c.done = closedchan
   } else {
      close(c.done)
   }
   for child := range c.children {
      // NOTE: acquiring the child's lock while holding parent's lock.
      child.cancel(false, err)
   }
   c.children = nil
   c.mu.Unlock()

   if removeFromParent {
      removeChild(c.Context, c)
   }
}
func removeChild(parent Context, child canceler) {
	p, ok := parentCancelCtx(parent)
	if !ok {
		return
	}
	p.mu.Lock()
	if p.children != nil {
		delete(p.children, child)
	}
	p.mu.Unlock()
}

```

这个操作中的第一个参数`removeFromParent`比较有意思，表示此`Context`节点是否应该从父节点的子节点中删除。照理说此节点的所有子节点都应该从父节点中删除，但是没有，只有在第一次调用`cancel`函数的时候，才会传入参数`removeFromParent`为`true`，其他的时候都是`false`。其实仔细想想也就不难理解了，没有必要。这个节点可能有多个子节点，并且子节点也可能有很多子节点，这些节点从不从父节点删除都是无所谓的，因为一颗子树已经删除了，后续的每个节点的剥离只不过是浪费时间。

由于这种删除操作是深度优先的，如果都传入`true`，则会从最底部的节点开始删除。并不会因为传入`true`，就会造成删除过程出现 bug。

### 其他类型
调用方法有四种，分别如下
1. WithCancel
2. WithDeadline
3. WithTimeout
4. WithValue
`WithDeadline`返回的是`timerCtx`类型，就是包了一层的`cancelCtx`。可以定时到指定的时间执行`cancel`的操作，或者手动的执行`cancel`操作。
`WithTimeout`是转化为`WithDeadline`执行的。
`WithValue`大家可以在网上找找例子看看如何使用，其返回的类型为`valueCtx`也没什么说的，各位看看代码就可以理解了。

## 总结

个人感觉这个包的源码挺简单的，但是解决的问题是非常具有重要意义的。这个包虽然简单，但是可以通过各种`Context`组合，形成复杂的操作，这就是其厉害之处。


# time/rate 限流源码解析



## 背景

限流器一般采用的算法是`令牌桶`算法。简单来说就是有一个容量一定的桶，可以用来装令牌。系统会以固定的速率往桶里面放令牌，如果令牌数量超过了桶的大小，则抛弃。用户则从桶中取令牌，在取得令牌的情况下才可以执行操作，否则需要等待令牌，或者放弃操作。

## 基本结构

`/time/rate`有两个基本的结构体，先来介绍下两个基本的结构体，因为在仔细了解参数的含义之后，才可以更好的理解代码的执行过程。

### Limiter

```go
type Limit float64

type Limiter struct {
	mu     sync.Mutex
	limit  Limit
	burst  int
	tokens float64
	// last is the last time the limiter's tokens field was updated
	last time.Time
	// lastEvent is the latest time of a rate-limited event (past or future)
	lastEvent time.Time
}

```

这个就是限流器了，第一个参数是`mu`，用来控制限流器内部的参数的更新，参数更新的操作是互斥的。第二个参数是`limit`，本质上就是一个`float64`类型的数据，表示的是限流器产生令牌的速度。第三个参数是`burst`，就是桶能装的令牌数量的大小。第四个参数`tokens`，表示的是当前桶内的令牌数量。第五个参数`last`，是一个时间类型的参数，表示的是桶内`tokens`最近更新的时间。第六个参数`lastEvent`，也是一个时间类型的参数，表示的是事件最近执行的时间(可以在过去，也可以在将来，这个后面会讲到)。

### Reservation

```go
// A Reservation holds information about events that are permitted by a Limiter to happen after a delay.
// A Reservation may be canceled, which may enable the Limiter to permit additional events.
type Reservation struct {
	ok        bool
	lim       *Limiter
	tokens    int
	timeToAct time.Time
	// This is the Limit at reservation time, it can change later.
	limit Limit
}
```

这个结构体就比较有意思了，就是一种预约的操作。在给定的时间限制内，如果可以预约的上，那么我就执行。所以参数也就比较好解释了。第一个参数是`ok`，就是表示用户预约没预约上，`true`就表示预约上了。第二个参数是`lim`，用于指向具体的限流器。第三个参数是`tokens`，则表示的是我这次预约的数量。第四个参数就是`timeToAct`，也比较有意思了，就是预约的时间了。第五个参数是`limit`，则是在预约的时候，限流器的产生速度，其实是可以通过变量`lim`来获取限流器产生令牌的速度，那么为什么还要单独的整出这个参数呢？是因为`Limiter`中的`limit`是可变的，修改`limit`参数的方式如下:

```go
// SetLimit is shorthand for SetLimitAt(time.Now(), newLimit).
func (lim *Limiter) SetLimit(newLimit Limit) {
   lim.SetLimitAt(time.Now(), newLimit)
}

// SetLimitAt sets a new Limit for the limiter. The new Limit, and Burst, may be violated
// or underutilized by those which reserved (using Reserve or Wait) but did not yet act
// before SetLimitAt was called.
func (lim *Limiter) SetLimitAt(now time.Time, newLimit Limit) {
   lim.mu.Lock()
   defer lim.mu.Unlock()

   now, _, tokens := lim.advance(now)

   lim.last = now
   lim.tokens = tokens
   lim.limit = newLimit
}
```

## 涉及的方法

```go
// Allow is shorthand for AllowN(time.Now(), 1).
func (lim *Limiter) Allow() bool {
	return lim.AllowN(time.Now(), 1)
}

// AllowN reports whether n events may happen at time now.
// Use this method if you intend to drop / skip events that exceed the rate limit.
// Otherwise use Reserve or Wait.
func (lim *Limiter) AllowN(now time.Time, n int) bool {
	return lim.reserveN(now, n, 0).ok
}

// Reserve is shorthand for ReserveN(time.Now(), 1).
func (lim *Limiter) Reserve() *Reservation {
	return lim.ReserveN(time.Now(), 1)
}

// ReserveN returns a Reservation that indicates how long the caller must wait before n events happen.
// The Limiter takes this Reservation into account when allowing future events.
// The returned Reservation’s OK() method returns false if n exceeds the Limiter's burst size.
// Usage example:
//   r := lim.ReserveN(time.Now(), 1)
//   if !r.OK() {
//     // Not allowed to act! Did you remember to set lim.burst to be > 0 ?
//     return
//   }
//   time.Sleep(r.Delay())
//   Act()
// Use this method if you wish to wait and slow down in accordance with the rate limit without dropping events.
// If you need to respect a deadline or cancel the delay, use Wait instead.
// To drop or skip events exceeding rate limit, use Allow instead.
func (lim *Limiter) ReserveN(now time.Time, n int) *Reservation {
	r := lim.reserveN(now, n, InfDuration)
	return &r
}


// Wait is shorthand for WaitN(ctx, 1).
func (lim *Limiter) Wait(ctx context.Context) (err error) {
	return lim.WaitN(ctx, 1)
}

// WaitN blocks until lim permits n events to happen.
// It returns an error if n exceeds the Limiter's burst size, the Context is
// canceled, or the expected wait time exceeds the Context's Deadline.
// The burst limit is ignored if the rate limit is Inf.
func (lim *Limiter) WaitN(ctx context.Context, n int) (err error) {
	lim.mu.Lock()
	burst := lim.burst
	limit := lim.limit
	lim.mu.Unlock()

	if n > burst && limit != Inf {
		return fmt.Errorf("rate: Wait(n=%d) exceeds limiter's burst %d", n, burst)
	}
	// Check if ctx is already cancelled
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	// Determine wait limit
	now := time.Now()
	waitLimit := InfDuration
	if deadline, ok := ctx.Deadline(); ok {
		waitLimit = deadline.Sub(now)
	}
	// Reserve
	r := lim.reserveN(now, n, waitLimit)
	if !r.ok {
		return fmt.Errorf("rate: Wait(n=%d) would exceed context deadline", n)
	}
	// Wait if necessary
	delay := r.DelayFrom(now)
	if delay == 0 {
		return nil
	}
	t := time.NewTimer(delay)
	defer t.Stop()
	select {
	case <-t.C:
		// We can proceed.
		return nil
	case <-ctx.Done():
		// Context was canceled before we could proceed.  Cancel the
		// reservation, which may permit other events to proceed sooner.
		r.Cancel()
		return ctx.Err()
	}
}

```

上面的列出来的方法虽然多，但是底子里用的都是一个方法`reserveN`。实际中用到的最多的一个方法是`advance`，主要作用是更新下到达时间`now`的时候，`Limiter`中的`tokens`应该是多少。里面需要注意的`now.Before(last)`可能出现的场景是事件操作的时间并不是一定是按照时间的顺序进行操作的。

```go
// advance calculates and returns an updated state for lim resulting from the passage of time.
// lim is not changed.
// advance requires that lim.mu is held.
func (lim *Limiter) advance(now time.Time) (newNow time.Time, newLast time.Time, newTokens float64) {
   last := lim.last
   if now.Before(last) {
      last = now
   }

   // Avoid making delta overflow below when last is very old.
   maxElapsed := lim.limit.durationFromTokens(float64(lim.burst) - lim.tokens)
   elapsed := now.Sub(last)
   if elapsed > maxElapsed {
      elapsed = maxElapsed
   }

   // Calculate the new number of tokens, due to time that passed.
   delta := lim.limit.tokensFromDuration(elapsed)
   tokens := lim.tokens + delta
   if burst := float64(lim.burst); tokens > burst {
      tokens = burst
   }

   return now, last, tokens
}
```

接下来就要见识庐山真面目了，`reserveN`的源码如下：

```go
// reserveN is a helper method for AllowN, ReserveN, and WaitN.
// maxFutureReserve specifies the maximum reservation wait duration allowed.
// reserveN returns Reservation, not *Reservation, to avoid allocation in AllowN and WaitN.
func (lim *Limiter) reserveN(now time.Time, n int, maxFutureReserve time.Duration) Reservation {
	lim.mu.Lock()

	if lim.limit == Inf {
		lim.mu.Unlock()
		return Reservation{
			ok:        true,
			lim:       lim,
			tokens:    n,
			timeToAct: now,
		}
	}

	now, last, tokens := lim.advance(now)

	// Calculate the remaining number of tokens resulting from the request.
	tokens -= float64(n)

	// Calculate the wait duration
	var waitDuration time.Duration
	if tokens < 0 {
		waitDuration = lim.limit.durationFromTokens(-tokens)
	}

	// Decide result
	ok := n <= lim.burst && waitDuration <= maxFutureReserve

	// Prepare reservation
	r := Reservation{
		ok:    ok,
		lim:   lim,
		limit: lim.limit,
	}
	if ok {
		r.tokens = n
		r.timeToAct = now.Add(waitDuration)
	}

	// Update state
	if ok {
		lim.last = now
		lim.tokens = tokens
		lim.lastEvent = r.timeToAct
	} else {
		lim.last = last
	}

	lim.mu.Unlock()
	return r
}
```

判断用户是否可以执行的条件是有两个，其一是`n`小于等于`令牌桶`的容量，其二是用户是否可以等待足够的时间以便`令牌桶`可以产生足够的令牌。如果满足了这两个条件，那么会给一个用户执行任务的时间`timeToAct=now.Add(waitDuration)`。在`now`时间点，如果令牌的数量减去`n`是一个负值，就说明用户需要等待。那么等待可以执行的条件是什么呢？在`timeToAct`的时间点，此时令牌桶的容量是0，之前是负的。这就是一种预约的操作，我现在的量不够你消费，如果你可以等待，那么我会首先生产你需要的。

其实剩下还有一个问题，就是有用户预约到了一个时间，但是他又取消了。那么应该怎么做呢？这个就是`CancelAt`的作用呢

```go
// CancelAt indicates that the reservation holder will not perform the reserved action
// and reverses the effects of this Reservation on the rate limit as much as possible,
// considering that other reservations may have already been made.
func (r *Reservation) CancelAt(now time.Time) {
   if !r.ok {
      return
   }

   r.lim.mu.Lock()
   defer r.lim.mu.Unlock()

   if r.lim.limit == Inf || r.tokens == 0 || r.timeToAct.Before(now) {
      return
   }

   // calculate tokens to restore
   // The duration between lim.lastEvent and r.timeToAct tells us how many tokens were reserved
   // after r was obtained. These tokens should not be restored.
   restoreTokens := float64(r.tokens) - r.limit.tokensFromDuration(r.lim.lastEvent.Sub(r.timeToAct))
   if restoreTokens <= 0 {
      return
   }
   // advance time to now
   now, _, tokens := r.lim.advance(now)
   // calculate new number of tokens
   tokens += restoreTokens
   if burst := float64(r.lim.burst); tokens > burst {
      tokens = burst
   }
   // update state
   r.lim.last = now
   r.lim.tokens = tokens
   if r.timeToAct == r.lim.lastEvent {
      prevEvent := r.timeToAct.Add(r.limit.durationFromTokens(float64(-r.tokens)))
      if !prevEvent.Before(now) {
         r.lim.lastEvent = prevEvent
      }
   }

   return
}
```

这部分代码比较难以理解的就是这个操作了: `restoreTokens := float64(r.tokens) - r.limit.tokensFromDuration(r.lim.lastEvent.Sub(r.timeToAct))`。作者为什么要这么减呢？为什么不能直接`restoreTokens := float64(r.tokens)`呢？这样这部分产生的令牌就不会浪费了，给后续需要令牌的用户进行消费。

根据`令牌桶`算法是可以知道如下两点的：

1. 令牌产生的速度一定，所以在一个时间段内产生的令牌是有限的
2. `令牌桶`的大小固定，所以如果一个时间段内产生的令牌无用户使用，这些令牌是不能进行无限保存的。

令牌的产生，可以理解为把时间线进行切分，并且可以把所有要根据令牌得到执行权力的事件当成时间线中的某一段，可以如下图所示一样。中间的一条直线可以理解为时间线，线上的每个箭头可以理解为一个令牌。下图中的`事件N`旁边的两条数线表示的就是此事件占用的令牌数量。

![此图还需要再改改](/Users/hushichang/Library/Application Support/typora-user-images/image-20200726143025811.png)

那么`事件2`在红色箭头指向的时间点进行了`CancelAt`的操作的时候(操作之前已经有`事件3`和`事件4`预约好了)，如何恢复此事件占用的令牌的影响呢？

肯定是要恢复一定数量的`restoreTokens`的，先来说说加上`restoreTokens`的影响吧！**加上一定数量的`restoreTokens`会使得本来在`事件4`执行时间点的令牌数量为`0`的变为`restoreTokens`。也就是说如果在`事件4`执行之前的一个时间点进来一个事件要执行(假设需要的`tokens`数量少于等于`restoreTokens`)，是可以得到令牌进行执行的操作。**那么新来的事件和`事件4`执行的时间会有重叠。此时就要看看`令牌桶`算法的基本含义了，本来在`事件4`执行时间段内的令牌只够`事件4`执行，但是由于`事件2`的取消，造成了新来的事件在`事件4`执行的时间段内也会得到执行，但是这段时间内令牌产生的数量是只够`事件4`使用的，这样就违背了`令牌桶`算法了。

作者的做法是`restoreTokens := float64(r.tokens) - r.limit.tokensFromDuration(r.lim.lastEvent.Sub(r.timeToAct))`，如果后续事件需要的`tokens`数量少于`事件2`需要的`r.tokens`数量，那么加上一个差值，即使加上`restoreTokens`之后，虽然也有重叠，但是重叠之后的时间段加上`事件3`和`事件4`的时间段可以当成一个整体，作为之前取消的`r.tokens`数量。

>本文中的`事件3`和`事件4`可以是任意数量的事件，在文中仅仅是为了好理解，所以有了具体的实例。
# redis事务

## 背景

在多线程个线程对同一个变量修改的时候，往往会希望在自己修改的过程中，不要让其他线程来对此变量进行修改。我们在使用redis的时候，也是会产生这种需求的。

## 解决方案

Redis 的事务需要用到 MULTI 命令和 EXEC 命令，被MULTI和EXEC包围的命令相当于一个原子性的操作，执行的过程中不会夹杂着其他的redis操作。

### 不使用事务

在不使用事务，多个客户端对键`no_trans:`进行修改值的时候，会遇到一个背景中提到的问题。一个客户端修改的过程中，另外的客户端的命令也在对同一个键进行修改，就会造成结果不是咱想要的情况。下面的代码中通过多个线程来模拟这个过程，会得到变量`no_trans:`打印的结果是递增的，由1到3。

```python
import time
import threading
from redis import Redis

redis_conn = Redis()

def no_trans():
    print redis_conn.incr("no_trans:")
    time.sleep(0.1)
    redis_conn.incr("no_trans:", -1)

def test(func):
    threads = []
    for i in range(3):
        thread = threading.Thread(target=func)
        thread.start()
        threads.append(thread)
    # 等待每个线程执行完毕
    for thread in threads:
        thread.join()
    
if __name__ == '__main__':
    test(no_trans)

```

### 使用事务

在使用了事务的情况下就不会这样了。python中使用过`pipeline`来标识事务，函数默认参数是`True`，在执行了`pipeline.execute()`的时候，多个命令会一起传递给redis服务器，并且这些命中执行的过程中间不会掺杂其他命令。为了避免redis和客户端之间的网络消耗，也可以给`pipeline`函数传递参数`False`，多个命令同时传递给redis，不过执行的过程中不会按照事务来进行处理，命令执行的过程中可能会执行其他命令。由于每个事务执行的过程中是不会有其他的命令打扰的，所以在使用事务对上面的操作修改之后，就不会出现一样的问题了。

```python
def trans():
    pipeline = redis_conn.pipeline()
    pipeline.incr("trans:")
    time.sleep(0.1)
    pipeline.incr("trans:", -1)
    print pipeline.execute()[0]
    
if __name__ == '__main__':
    test(trans)

```


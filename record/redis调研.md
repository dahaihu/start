# redis调研

年前有一个需求，要统计10min内的某个ip或者用户的请求数量，根据此做广告反作弊系统。为此，我首先想到的是使用redis对键设置10min的过期时间，然后就可以快速的访问了。而由此耗费的内存也是不大的，经过验证也是可以接受的。

## 前期

为了做最坏的打算，可以假设请求的ip都是ipv6(ipv4长度比ipv6短，ipv6更占用内存)，并且ip没有进行省略的表达。由于不想写生成ipv6的代码，所以直接使用python的第三方库`netaddr`，可以按照指定的范围生成所有的ip，最后生成的ip数量为15658735。

```python
from netaddr import *
r1 = IPRange('1111:1111:1111:1111:1111:1111:1111:1111', '1111:1111:1111:1111:1111:1111:11ff:ffff')
with open('ip.txt', 'w') as file:
  for ip in r1:
      file.write('{}\n'.format(ip))
```

>为了生成足够的ip，所以我设置的范围比较大。读者可以按需生成

## 压测

由于redis可以使用lua做扩展，所以开始设置的是lua读取`ip.txt`文件，然后随机返回一个ip，再做`incr ip`操作。写好的lua脚本文件名为`redis_script.lua`，内容如下：

```lua
local data = {}
for line in io.lines('ip.txt') do
        table.insert(data, line)
ends
math.randomseed(os.time())
redis.call('set', data[math.random(1, #data)], 1)
```

下一步就是将此函数加载到redis之中，执行之后会返回一个`sha`，需要将此`sha`记住，此后可以在redis中直接调用。

```lua
redis-cli --eval redis_script.lua
```

可是呢，在这一步的时候出错了，代码中使用了`io`库，而在使用的lua扩展的redis，[只加载了必要的几个库](https://redis.io/commands/eval#available-libraries)，所以到此就game over了，不能继续下去。

### 感想

- 此脚本的执行是非常消耗时间的，因为每次执行都要读取一个一千多万行的文件。
- redis压根不可能读取文件的，lua拓展在redis中相当于一个可执行的函数，而在redis中肯定是没法读取本地文件的。

为了模拟ip请求然后计数的情况，修改了执行方式，不再使用lua扩展。而是把所有的ip放到redis中的一个集合内，每次模拟请求的时候，从集合中随机抽取ip作为请求的ip。

### 加载ip到redis

首先使用的是命令行，读取所有文件，然后一个ip一个ip的添加，其中使用到了`xargs`命令，一个我觉得非常有用的命令。

```shell
cat ip.txt | xargs -I {} redis-cli sadd ip_set {}
```

一千五百多万次的操作，让我等得实在不耐烦了。在每个请求需要`1ms`的假设下，估算了下总共需要的时间`15658735 * 0.001 / 60 / 60 = 4h`，大约需要4个小时。其中主要耗时是在网络请求上，所以在操作较为频繁的时候，避免网络耗费时间，redis建议使用`pipeline`。

为了速度的解决战斗，我使用了python，每次`sadd`添加的元素数为100个，当然也可以设置更多。

```python
with open('ip.txt') as file:
    ips = file.read().strip().split('\n')

def write_to_redis(ips):
    length = len(ips)
    for i in range(length / 100 + 1):
        redis.sadd('ip_set', *ips[i*100, (i+1) * 100])

write_to_redis(ips)
```

后来想想，其实也是不用写python代码来实现的，因为`sadd`可以对集合一次添加多个元素，那么在命令行也是可以同样做到的。`xargs`可以将每100个ip作为参数，传递给`redis-cli`作为添加到`ip_set`中的元素。

```shell
cat ip.txt | xargs -n 100 redis-cli sadd ip_set
```

### 压测开始

压测部分的代码比较简单，还是使用到了lua拓展。

第一步是完成lua脚本，脚本内容如下，有两步：1. 获取ip；2. ip计数加1

```lua
local key = redis.call('srandmember', 'ip_set', 1);return redis.call('incr', key[1])
```

第二步是加载脚本到redis，因为代码量较少，所以直接写在了命令行里面，而不是从文件中加载。由于`srandmember`可以指定随机返回元素的数量，所以返回的是一个数组。脚本加载到redis会返回一个`sha`，之后在redis内部，可以直接执行这个`sha`来执行对应的lua脚本。

```shell
redis-cli script load "local key = redis.call('srandmember', 'ip_set', 1);return redis.call('incr', key[1])"
# "dd575305e1aa4c3953449602b9b2bcbe59c51bad"
```

第三步完成测试

由于广告反作弊的请求的qps峰值为1万，所以10min的总请求数量为`10 * 60 * 100000 = 6000000`,有六百万的请求量

是可以直接使用redis提供的`redis-benchmark`来进行压测的

```shell
redis-benchmark -n 6000000 evalsha 'dd575305e1aa4c3953449602b9b2bcbe59c51bad' 0
```

命令中参数`n`表示的请求数量，`evalsha`表示的是需要执行的lua脚本的`sha`，后面的0表示传递给脚本的参数个数为0。

执行的结果如下，六百万个请求在一分半就结束了，每秒的请求数量可达六万五。

```
====== evalsha dd575305e1aa4c3953449602b9b2bcbe59c51bad 0 ======
  6000000 requests completed in 92.54 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1

89.89% <= 1 milliseconds
97.76% <= 2 milliseconds
99.06% <= 3 milliseconds
99.49% <= 4 milliseconds
99.74% <= 5 milliseconds
99.89% <= 6 milliseconds
99.96% <= 7 milliseconds
99.98% <= 8 milliseconds
99.99% <= 9 milliseconds
100.00% <= 10 milliseconds
64834.03 requests per second
```

总的来说说，redis是可以满足要求的。此时在redis中通过`info`命令可以查看使用的内存量为2.61G，此数据包括了所有ip的数量，以及几次执行`redis-benchmark`的结果。2.6G的内存消耗，对于反作弊的需求也是可以接受的。
## 初识NGINX

### 前言

此文是翻译的http://nginx.org/en/docs/beginners_guide.html，由于本人不做php部分的开发，所以去掉了最后一节关于php部分的引用。

### 正文

此文对nginx有一个基本的介绍，并提供了几个nginx可以解决的任务。本文假设读者在电脑上已经安装了nginx，如果没有安装，可以按照[nginx安装页面](http://nginx.org/en/docs/install.html)的介绍进行安装。此文档介绍了如何启动和安装nginx，根据配置文件重启，介绍了配置文件的结构，并描述了如何设置nginx用于服务静态文件，如何配置nginx为代理，和如何设置nginx连接FastCGI应用。

nginx有一个master进程和几个worker进程。master进程的作用是读取和使得配置文件生效，并维护worker进程。worker进程则用于处理实际的请求。nginx使用的是基于事件模型和依赖于系统的机制来在worker进程之间有效的分配请求。worker进程的数量可由配置文件配置，也可由指定的配置文件进行修改(这句话的意思是nginx可以由指定的配置文件启动)或者自动适应cpu的核数。

nginx和其模块的工作方式是由配置文件决定的。文件的默认名为`nginx.conf`，并且默认的存放在`/usr/local/nginx/conf`, `/etc/nginx`, 或者 `/usr/local/etc/nginx`文件夹下。

### nginx的启动，停止，重启

可以通过执行nginx的执行文件来启动nginx。一旦nginx启动了， 它就可以由命令行参数`-s`来进行控制，控制方式如下:

```shell
nginx -s signal
```

其中*signal*可以为如下四种：

* stop - 快速的关闭
* quit - 优雅的关闭
* reload - 重启配置文件
* reopen - 重新打开log文件

例如，为了在nginx处理了当前已经接受的请求之后再退出，可以使用如下的命令

```shell
nginx -s quit
```

>此命令应该在同一个用户的环境下执行

配置文件只有在使用重载配置文件或者重启nginx之后才会生效。为了重启配置，可以执行如下命令:

```shell
nginx -s reload
```

一旦nginx收到重启配置的指令，它就会检查新配置文件在语法上的有效性并尝试应用新提供的配置文件。在执行成功之后，nginx会启动新的worker进程，并发送消息给旧的工作进程，让它们停止；否则，master进程会回滚，继续让旧的worker进程按照老的配置工作。旧的工作进程会接受停止命令，然后停止接受新的连接并把当前的连接处理完毕。然后，旧的工作进程会退出。

另外也可以使用Unix提供的`kill`命令来传递指令给nginx。这个时候指令是通过指定的进程ID来传递给nginx的。nginx的master进程的ID会默认写到 `/usr/local/nginx/logs` 或者 `/var/run`目录下`nginx.pid`文件里。例如，在master进程的ID是1628的时候，可以使用如下命令来优雅的关闭nginx的:

```shell
kill -s QUIT 1628
```

为了获得所有正在运行的nginx进程，可以使用`ps`命令，例如：

```shell
ps -ax | grep nginx
```

可以看[控制nginx](http://nginx.org/en/docs/control.html)，来获取更多的发送指令给nginx的消息

### 配置文件的结构

nginx包含的模块是由配置文件中的指令控制的。指令可以区分为简单指令和块指令。一个简单指令包含由空格分割的名称和参数，并以分号(;)进行结尾。一个块状的指令整体的结构和简单指令相同，但是这些指令由大小括号({和})包围。如果在大括号内部之中还有大括号，那么外部括号则是内部括号的上下文环境(例如: [events](http://nginx.org/en/docs/ngx_core_module.html#events), [http](http://nginx.org/en/docs/http/ngx_http_core_module.html#http), [server](http://nginx.org/en/docs/http/ngx_http_core_module.html#server), and [location](http://nginx.org/en/docs/http/ngx_http_core_module.html#location))。

置于任何括号之外的指令是[main](http://nginx.org/en/docs/ngx_core_module.html)上下文环境。`events`和`http`指令在main上下文环境之中，`server`置于`http`环境之中，`location`置于`server`环境之中。

每行指令中置于`#`之后的是注释

### 服务于静态文件

web服务器的一个重要功能就是返回文件给客户端(例如图片或者静态的HTML页面).你可以实现一个实例，根据请求的不同来将文件由本地 `/data/www` (包含html文件) 和`/data/images` (包含图片)进行分发。这将需要设置配置文件，设置一个置于 [http](http://nginx.org/en/docs/http/ngx_http_core_module.html#http) 模块之中的 [server](http://nginx.org/en/docs/http/ngx_http_core_module.html#server) 模块 ，并且`server`模块之中有两个 [location](http://nginx.org/en/docs/http/ngx_http_core_module.html#location) 指令块。

首先，创建一个`/data/www`的目录，并将`index.html`置于其中。并创建目录`data/images/`，然后放置几张图片于其中。

然后，打开配置文件。默认的配置文件中已经包含几个`server`块的示例，通常是被注释掉的。此时，注释掉所有这种指令块，然后开始一个新的指令块:

```shell
http {
	server {
	}
}
```

通常，配置文件中包含多个由监听的端口或者[服务名称](http://nginx.org/en/docs/http/server_names.html)区分的`server`块。一旦nginx决定哪个`server`处理请求，它会测试请求头部的URI和定义在`server`块中的`location`中的指令是否匹配。

添加如下`location`块到`server`指令块之中：

```shell
location / {
	root /data/www;
}
```

此`location`块会指定前缀`/	`和请求的URI进行比较。*在匹配的时候，请求的URI会添加到指定的[root](http://nginx.org/en/docs/http/ngx_http_core_module.html#root)路径之后*，来组成访问的本地文件的路径，在此例中就是添加到`/data/www`之后。`如果有多个匹配的路径，那么会匹配最长匹配的路径。`定义的`location`块可以提供一个最短的长度为1的前缀，然后在其他`location`都不匹配的时候使用。

接下来添加第二个`location`块：

```shell
location /images/ {
	root /data;
}
```

此块可以用来匹配URI以`/images/`开头的请求(以`/`开头的`location`也会匹配，但是匹配的前缀更短)。

最后的`server`块配置如下：

```shell
server {
    location / {
        root /data/www;
    }

    location /images/ {
        root /data;
    }
}
```

这是一个服务器可以运行的配置，用来监听标准的80端口，可以在本地通过 `http://localhost/`进行访问。对于URI以`/images/`开头的请求，服务器将从路径`/data/images`中发送文件。例如，对于uri为`http://localhost/images/example.png`的请求，nginx将会发送文件`/data/images/example.png`。如果此文件不存在，那么nginx将会发送一个表示404错误的响应。对于不是以`/images/`开头的请求，nginx将会从'/data/www'文件夹下发送文件。例如，对于URI为`http://localhost/some/example.html`的请求，nginx将会发送文件`/data/www/some/example.html`。

如果nginx没有启动，那么使得新的配置文件生效只需要启动nginx即可；如果nginx已经启动了，那么可以传递`reload`指令给nginx的master进程即可，命令如下：

```shell
nginx -s reload
```

### 设置一个简单的代理服务器

nginx常常使用的一个方式是设置为一个代理服务器，这意味着nginx接受请求，然后把请求传递给代理的服务器，再从这些代理的服务器中获取相应，最后返回给客户端。

接下来我们将配置一个基础的代理服务器，此服务器将会把图片和文件的响应直接从本地文件夹返回，而其他响应则传递给代理的服务器。在这个例子中，两个服务器将会定义在一个nginx实例之中。

首先，通过添加一个`server`块的方式添加一个代理服务器，`server`块中包含如下内容

```shell
server {
    listen 8080;
    root /data/up1;

    location / {
    }
}
```

这是一个简单的服务器，会监听8080端口(在之前是没有使用`listen`指令的原因是nginx默认监听的是80端口)，然后把所有请求映射到本地的`/data/up`文件夹。创建此文件夹，然后把`index.html`文件放在此文件夹下面啊。注意，`root`指令放在了`server`环境之中，此`root`指令会在处理请求的`location`块中没有自己的`root`指令的时候使用。

下一步就是修改上一步制定的配置文件，使得它可以作为一个代理服务器使用。在第一个`location`块，使用 [proxy_pass](http://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass) 指令，并且参数中包括使用的协议、名称以及端口。在我们的例子中是`http://localhost:8080`:

```shell
server {
    location / {
        proxy_pass http://localhost:8080;
    }

    location /images/ {
        root /data;
    }
}
```

接下来我们将修改第二个`location`块，此时该`locatinon`指令块会在请求的URI匹配`/images/`的时候将置于`/data/images`文件夹下的文件返回给客户端。修改后的`location`块如下

```shell
location ~ \.(gif|jpg|png)$ {
    root /data/images;
}
```

参数是一个匹配所有以`.gif`,`.jpg`或`.png`结尾的URI。正则表达式之前需要一个符号`~`表示后面的是按照正则表达式匹配的。对的请求会映射到对应的`/data/images`文件夹下的文件。

在nginx选择处理请求的`location`块的时候，会首先检查`location`指令，记住最长的`location`前缀，然后会检查正则匹配。如果有正则匹配匹配上的`location`，那么会使用此`location`指令，否则会使用之前记住的`location`指令。

最终的`server`配置如下：

```shell
server {
    location / {
        proxy_pass http://localhost:8080/;
    }

    location ~ \.(gif|jpg|png)$ {
        root /data/images;
    }
}
```

此服务器会将以`.gif`,`.jpg`或者`.png`结尾的请求映射到文件夹`/data/images`下的文件。然后把其他请求传递给代理的服务器上。

要想使得配置的文件生效，需要想前面讲到的一样，把`reload`指令传递给nginx。

另外还有[更多的指令](http://nginx.org/en/docs/http/ngx_http_proxy_module.html)，用于处理代理服务器。

## 安装buildout
`pip install zc.buildout`
## 初始化
创建文件夹firstbuildout，然后进入该文件夹执行buildout命令,会产生如下就结果
```
├── bin
│   └── buildout
├── buildout.cfg
├── develop-eggs
│   ├── setuptools.egg-link
│   └── zc.buildout.egg-link
├── eggs
└── parts
```

## 编辑文件
首先是完成脚本，创建main包(必须有一一个`__init__.py`文件)，里面有一个main.py的文件。文件内容如下
```
import requests
def main():
    print requests.get('https://www.baidu.com')
```
然后需要编辑的是setup.py，里面设置了脚本的位置main包下main文件里的main函数。这样就可以准备完成之后，执行buildout就可以在bin的文件夹中生成一个可以直接执行的脚本。
```
#!/usr/bin/python
# coding: utf-8
from setuptools import setup
setup(name='firstbuildout',
      version='0.0.0',
      author='***',
      license='PRIVATE',
      install_requires=[
          'requests',
      ],
      entry_points={
          'console_scripts': [
            'test = main.main:main'
          ],
      }, )
```
另外还需要编辑的就是buildout.cfg。文件中配置了extends为versions.cfg，表示的是配置文件应该从versions.cfg中读取(配置文件可以有多个)。我们可以设置update-versions-file为versions.cfg，表示的是我们希望buildout来维护versions.cfg，在执行buildout的时候，会更新第三方库的版本(当然也会满足已配置的版本限制的)，这个时候会写入文件verssions.cfg之中的。文件内容如下
```
[buildout]
develop  = .
include-site-packages = true
update-versions-file = versions.cfg
extends = versions.cfg
parts =
    tools

[tools]
recipe = zc.recipe.egg
eggs =
    firstbuildout
scripts = test

```

## 必要操作
**首先需要将当前包制作为egg的形式，要不然在buildout执行过程中会找不到firstbuildout这个包的。因为这个包是此项目的包。**
`python setup.py sdist`

此操作会产生文件夹dist，里面有一个压缩文件firstbuildout-0.0.0.tar.gz

然后解压，之后就可以直接安装的。但是，可以等着buildout来给你安装
`tar xzf firstbuildout-0.0.0.tar.gz`


## 产生对应的脚本
首先需要创建versions.cfg文件，使用`touch versions.cfg`即可，因为buildout需要从中读取配置。然后执行`buildout`即可大功告成，就会产生想要的脚本了
```
➜  firstbuildout tree bin
bin
├── buildout
└── test

0 directories, 2 files
```

直接执行bin/test即可
```
➜  firstbuildout bin/test
<Response [200]>
```

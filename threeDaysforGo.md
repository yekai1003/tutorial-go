## 1. 写在前面
### 1.1 为什么要学习go语言？

Go（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。创作者包括罗伯特·格瑞史莫（Robert Griesemer），罗勃·派克（Rob Pike）及肯·汤普逊（Ken Thompson）Ian Lance Taylor、Russ Cox。这些在google都是响当当的角色，如果对unix历史有所了解的小伙伴就会知道，Thompson和Dennis.Ritchie被合称为unix之父以及C语言之父，由此可见，go语言创造团队相当强大！
学习Go语言的理由可以是以下几点：
- 并发
- 语法简洁
- 开发周期短
- 内存回收
- 运行效率高
- 统一的编程规范
- 活跃的社区文化


### 1.2 go语言能做什么？

Go语言利用其天生的并发特性，可以很好的开发服务器后端，目前在云计算领域以及区块链行业，Go都受到了热烈追捧，顶顶大名的Docker也是用Go语言开发的，Go语言的产业链正在日趋成熟。

### 1.3 go语言应该如何学？
Go语言的入门也非常简单，学习的路径可以是先搭建环境，然后基础语法掌握，此后学习Go语言核心编程和思想，再然后进入项目阶段实战，掌握行业应用所常用到的开源框架，整体而言，Go语言的学习成本并不高！
## 2. 环境搭建

开发环境始终是学习一门语言最重要的一环，Go语言作为2009年才被推出的新面孔来说，跨平台是必须支持的。Go语言的安装很简单，在Go语言的官网可以下载到安装包，但是由于某些原因，我们访问会有一些问题，但是也有一些好心人将Go语言的安装包同步到了[Go语言中文网](https://studygolang.com/dl)，在这里可以很顺利的下载。


![image](https://note.youdao.com/yws/public/resource/70bb16150523238188a2ec3bae0292ea/xmlnote/B9326AB72B9E4D2BADDE5E67B03BDF27/29568)

### 三大平台安装方法介绍

- windows**

对于windows的用户来说，下载windows版的msi文件即可，当然安装的过程也非常的傻瓜，各种下一步即可。默认情况下Go语言很关键的环境变量GOROOT和GOPATH在安装的时候会设定好，如果没有，需要手动设置一下这两个环境变量。GOROOT就是安装路径，而GOPATH则是后续要用到的重要路径，可以简单粗暴的认为它是开发路径。在它的下面，会有bin、pkg、src三个子目录，bin是存放二进制文件的目录，pkg是存放包编译后的.a文件，src就是我们的源码路径。（GOROOT与GOPATH的原理在所有平台是相同的）在命令行窗口，可以通过go env指令来查看Go语言的环境变量情况。

![image](https://note.youdao.com/yws/public/resource/70bb16150523238188a2ec3bae0292ea/xmlnote/C7C49EB2745343F2A7C41FBCD3CCEDFA/29590)

- Linux

在linux平台可以下载安装包安装，也可以使用命令行来操作。以ubuntu系统举例（centOS需要使用yum进行安装），下面两句指令就可以搞定Go语言的安装：


```
sudo apt-get update
sudo apt-get install golang
```
下面介绍使用安装包安装的方式，先下载安装包（amd64代表该版本支持64位操作系统，amd代表的则是cpu的架构，与之对应的是x86）

```
mkdir ~/install
cd ~/install
wget https://studygolang.com/dl/golang/go1.12.7.linux-amd64.tar.gz
```

将压缩包解压
```
tar -zxvf go1.12.7.linux-amd64.tar.gz -C ~/
```

设置环境变量
```
cd ~
export PATH=$PATH:~/go/bin
echo 'export PATH=$PATH:~/go/bin' >>~/.bashrc
mkdir ~/gowork
export GOPATH=$HOME/gowork
echo 'export GOPATH=$HOME/gowork' >>~/.bashrc
```

来查看一下环境变量

```
ubuntu@ip-172-31-26-216:~$ go env
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/ubuntu/.cache/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/ubuntu/gowork"
GOPROXY=""
GORACE=""
GOROOT="/home/ubuntu/go"
GOTMPDIR=""
GOTOOLDIR="/home/ubuntu/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build052037651=/tmp/go-build -gno-record-gcc-switches"
```

到这里，ubuntu上Go开发环境其实具备了，也可以趁热打铁，将GOPATH路径下的子目录都创建一下。


```
mkdir -p $GOPATH/src 
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg
```

- macOS

在macOS上安装Go语言也可以选择下载安装包和命令行的形式，如果想下载安装包安装的话，可以参考之前在ubuntu上的安装做法，两者是相同的，只不过包名不同。

```
https://studygolang.com/dl/golang/go1.12.7.darwin-amd64.pkg
```
命令行安装也是非常简单，使用brew工具即可！


```
brew install golang
```

### IDE开发工具安装

Go语言的IDE开发工具有很多，开发者完全可以根据自己喜好来选择，比如GoLand，VsCode，sublime text3等等，本人比较喜欢用LiteIDE。开发人员可以前往[github地址下载](https://github.com/visualfc/liteide/releases/tag/x36)。

![image](https://note.youdao.com/yws/public/resource/70bb16150523238188a2ec3bae0292ea/xmlnote/8E4C36B77EE44C4FAEBB755806B43BE6/29643)

安装后效果如下图所示：

![image](https://note.youdao.com/yws/public/resource/70bb16150523238188a2ec3bae0292ea/xmlnote/05A1C6B3F2AA4DE0872AD8CF0C520BA1/29649)

至于IDE的使用，本文不做详细说明，点点就会了，不过需要注意一点的是，Go语言是面向工程型的语言，每个工程里只能有一个main，所以同一个目录下有多个包含main函数的文件时是不能使用IDE编译并运行的，这时仍然需要借助命令行的操作，我们将在后面进行介绍如何编译和运行Go代码。


## 语法篇
从本篇开始，我们将介绍Go语言的语法知识，接下来跟我一起来写代码吧！
### 初识golang

别的先不说，先来看一段hello-world的代码：

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
```

### go工具集

- go get url  - 下载第三方包
- go run 


### 变量
### 常量
### 指针
### 函数
## 核心编程
### 容器编程
### 面向对象篇
### json处理
### 并发
### 同步
### 多路通道使用
### sync同步
## 网络编程
### 文件IO
### defer语句
### TCP编程
### HTTP编程
## 数据库编程

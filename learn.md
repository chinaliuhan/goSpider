## 内建变量类型

- bool
- string
- int

int 不加u代表有符号整数

(u)int 代表无符号 

规定长度 int8,int16,int32,int64

故不规定长度,按照操作系统来, 32位系统是32,在64位系统是64

自动 uintptr



- byte
- rune
这里的rune就是go的char

go语言中已经没有char了,只有rune

因为char里只有1字节,在多国语言时有很多坑

在utf-8里面有很多字符是3字节,因为GO采用了4字节的int32来做的rune

这个byte是8位的, rune是32位的 他们和整数都是可以混用的

文档里说的, 他们和整数来说就是一个别名

- 浮点型float32 float64  原生支持复数类型complex64 complex128




## GoPath的环境变

- GOPATH 一般就是我们自己定义的系统环境变量路径

- GOROOT 一般就是go语言自己设定的路径, 存放一些go语言自带的包之类的


Linux和Unix默认位置  ~/go下面

window下环境变量默认位置 %USERPROFIL%\go 

官方推荐: 所有项目和第三方库都放在同一个GOPATH下

也可以将不同的项目放在不同的GOPATH下面

go语言编译的时候回去到不同的GOPATH的路径中取寻找自己依赖的包

go语言自带的包会到自己原来自己的src目录中取找, 这个我们不用管

我们自己import 的包是到我们定义的GOPATH中取找

设置环境变量
export GOPATH=/Users/liuhao/go
一要用的话也可以把这个也设置进去
export PATH="$GOPATH/bin:$PATH"

一般go会默认设置环境变量,我们在编写代码时,使用import时,系统的库他回去自动找自己的$GOROOT库.
我们写的库的话,会到$GOPATH中取查找.


## intellij IDEA自动清掉我们文件中错误的import

如果直接用goland这个IDE的话不用搞这些,格式化代码的时候会自动去掉

老版本idea是在系统设置中的 language & framework 里面的go 里面有一个On save的几个选项
新版本idea需要在plugin 安装 file watch 然后在 设置中的tools 中的file watch 中添加
我的添加目录案例: /Users/liuhao/go/bin/goimports 然后文件自动保存或手动保存时就会清掉无用的imports

nothing     这个就不说了 

go fmt      只做格式化代码用

go imports  不仅格式化代码, 而且还会对错误的Import进行清除和格式化

在这之前呢需要我们下载和安装这个第三方的库,否则你也没有goimports文件,自动保存的时候会暴一个Can't find `goimports` in GOPATH.....

使用go get 获取第三方库

go get 获取golang.org 是不行的貌似被墙了

要用gopm 来获取无法下载的包

安装gopm
go get -v github.com/gpmgo/gopm

这是我们回到~/go 中的我们自己的GOPATH的目录 ls一下

可以看到$GOPATH下的src下多了一个github.com 的目录,和我们的目录放在一起
```shell
    ~/go ⌚ 2:06:05
    $ tree -L 2
    .
    ├── bin
    │   └── gopm
    └── src
        ├── github.com
        └── learnGo
```

这时候我们到~/go/bin/gopm 运行该文件

`~/go/bin/gopm get  -g -v -u golang.org/x/tools/cmd/goimports`

这是安装成功之后, $GOPATH 下的src中会多一个golang.org的目录和我们的目录放在一起
```
~/go ⌚ 2:19:31
$ tree -L 2
.
├── bin
│   └── gopm
└── src
    ├── github.com
    ├── golang.org
    └── learnGo
```

很显然我们在编码时的目录也是在src中运行的, bin是可执行文件

在~/go/bin 目录下执行 go install 安装goimports 安装之后会在~/go/bin目录下生成一个可执行的文件
~/go/bin ⌚ 2:26:38
$ go install  ../src/golang.org/x/tools/cmd/goimports


之后配置完之后, 我们在保存文件时,打码会自动按照标准尽心格式化, 无效的import也会被删除掉


1. go get 命令演示
2. 使用gopm 来获取无法下载的包
3. go build来编译 我们编写的go文件,但是会建立在当前目录下
4. go install 产生Pkg文件和可执行文件 将我们编写的go文件安装到~/go/bin下
    使用go install ./... 安装当前目录下所有的go的包
    go install 的时候一个package的main函数只能有1个,所以go语言的main函数都要在自己的一个目录下面
5. go run 直接编译运行

```aidl
~/go ⌚ 2:38:05
$ tree -L 2
.
├── bin
│   ├── goimports
│   └── gopm
├── pkg
│   └── darwin_amd64
└── src
    ├── github.com
    ├── golang.org
    └── learnGo

```

```
src 有很多第三方的包和我们的自己的项目放在里面, 每个人站一个目录
pkg 和src是对应的是我们build出来的一些中间过程,我们不用去管他
bin 就是我们生成的可执行文件

src
    git repository 1
    git repository 1
pkg
    git repository 1
    git repository 1
bin
    可执行文件1,2,3,4...
```
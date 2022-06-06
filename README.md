### Golang简介和特征

Go 语言是一个可以编译高效，支持高并发的，面向垃圾回收的全新语言。

- 秒级完成大型程序的单节点编译。
- 依赖管理清晰。
- 不支持继承，程序员无需花费精力定义不同类型之间的关系。
- 支持垃圾回收，支持并发执行，支持多线程通讯。
- 对多核计算机支持友好。



### Golang的下载安装并设置IDE（VS Code）

1. 下载Golang  **https://golang.google.cn/dl/**
2. 下载对应平台的二进制文件并安装
3. 配置环境变量
   - GOROOT --> go的安装目录
   - GOPAHT 
     - src: 存放源代码
     - pkg：存放依赖包
   - GOPROXY -->  设置**goproxy --> export GO111MODULE=on --> export GOPROXY=https://goproxy.cn**
4. 下载VS Code -->  **https://code.visualstudio.com/download**
5. 设置Golang的VS Code插件  -->  **https://marketplace.visualstudio.com/items?itemName=golang.go**
6. 



### 一些基本指令

- bug 生成错误报告。运行之后生成了一个给Golang提issus的模板。

![](/Users/hezefeng/Library/Mobile Documents/com~apple~CloudDocs/技术文稿/ 学习Golang/gobug.png)

- **build  编译包和依赖生成一个二进制文件**
  - 指定输出目录  go build -o /home/file
  - 常用环境变量设置编译操作系统和CPU架构 GOOS=linux GOARCH=amd64 go build

```shell
root@u20:~/go/src/learngolang/gobasecommand# go build main.go
root@u20:~/go/src/learngolang/gobasecommand# ls
main  main.go
root@u20:~/go/src/learngolang/gobasecommand# ./main 
hello world
```

- clean
  - -i 清除关联的包和二进制文件
  - -n 将执行的指令打印出来。
  - -cache 删除`go build`命令的缓存
  - -testcache 删除当前包所有的测试结果

```
root@u20:~/go/src/learngolang/gobasecommand# ls
main  main.go
root@u20:~/go/src/learngolang/gobasecommand# go clean -i main.go
root@u20:~/go/src/learngolang/gobasecommand# ls
main.go
root@u20:~/go/src/learngolang/gobasecommand# 
```

- doc 显示包或符号的文档
- env 打印Go环境

```
root@u20:~/go/src/learngolang/gobasecommand# go env
GO111MODULE="on"
GOARCH="amd64"
GOBIN="/usr/local/go/bin"
GOCACHE="/root/.cache/go-build"
GOENV="/root/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/root/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/root/go"
GOPRIVATE=""
GOPROXY="https://goproxy.cn"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.16"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build148419871=/tmp/go-build -gno-record-gcc-switches"
```

- fix 更新包以使用新的api
- fmt 格式化包的源代码
- generate   通过处理源文件生成Go文件
- **get 向当前模块添加依赖项并安装**
- **install 编译和安装包和依赖项**
- list 列出包或模块
- **mod  模块的维护**
- **run 编译并运行Go程序**
- **test   测试**
  - go test指令会自动读取目录下名为 *_test.go的文件，并生成运行测试用的可执行文件。
  - 在测试文件中函数以Testxxx开头会被识别成测试的函数

```go
func TestAdd(t *testing.T) {
	t.Log("Start testing")
	result := Add(1, 2)
	if result == 3 {
		t.Log("PASS")
	}
}
```

```
root@u20:~/go/src/learngolang/gobasecommand# go test
PASS
ok      learngolang/gobasecommand       0.004s
```

- tool 运行指定的go工具

- version 打印Go的版本

- vet 报告包中可能的错误

  

### Go的控制和结构

#### if 语句

跟其他语言的if循环一样，只不过go这里判断条件不用写括号。

```go
	if a > 0{
		xxxxx
	}else if b > 3{
		xxxxx		
	}
```



#### switch case 语句

```go
	var a, b, num int
	switch num {
		case a:
			fmt.Println(a)
		case b:
			fmt.Println(b)
    default:
    	// 如果都未匹配则会走default
	}
```



#### for 循环

- 有限的循环
- 等价于while的循环（Go语言不支持while）
- 无限循环

```go
// 有限的循环	
var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}

//等价于while的循环
	sum := 1
	for sum < 1000 {
		sum += sum
	}

// 无限循环 就写一个for 就完成了
	for {
		fmt.Println("hi")
	}

```



#### for-range语句

For range循环是将一个东西循环到底，可以遍历字符串、切片、数组、Map等等。

```go
// 遍历字符串
	str := "helloworld"
	for _, v := range str {
		fmt.Println(string(v), v)
	}
// 循环出来v是int32类型的，需要再用string()转一下
/*
h 104
e 101
l 108
l 108
o 111
w 119
o 111
r 114
l 108
d 100
*/

// 遍历Map
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["No1"] = "Apple"
	myMap["No2"] = "Orange"

	for i, v := range myMap {
		fmt.Println(i, v)
	}
	
/*
No1 Apple
No2 Orange
*/
```



### Go常用数据结构

#### 常量和变量以及定义

- 常量定义： const xxx type

- 变量定义： var xxx type

```
var a, b, c, d int = 1, 2, 3, 4

var a, b, c, d = 1, true, "six", false
```



#### 类型转换

通过int, float32,float64,uint这些函数进行类型转换

```
	var i int = 12
	var f float64 = float64(i)
	var u uint = uint(f)
```

#### 数组和切片

- 数组的定义：相同类型且长度固定连续内存片段
- 访问元素：以下标形式访问
- 定义一个数组： var myarray [len] type

```
//新建一个长度为3类型为int的数组
	var myintarray [3]int
	myintarray[0] = 1
	
//若不知道数组长度可以用[...]替代
	var myintarray = [...]int{7, 8, 9, 10, 5, 2, 3}
	for i, v := range myintarray {
		fmt.Println("i= ", i, "v= ", v)
	}
```



- 切片的定义：切片是对数组有一个连续片段的引用
- 定义一个切片：数组定义中不指定程度即为切片
- 切片在未初始化之前默认为nil,长度为0

#### Make和New

- New返回指针地址
- Make返回第一个元素，可以预设内存空间，避免未来的内存拷贝

#### Map

- 声明方法  var myMap = make(map[string]string)
- 添加Map元素  myMap["a"] = ["b"]
- 遍历Map元素 for range

#### 结构体、结构体标签和指针

- 通过type ... struct 关键字定义结构体
- Go语言支持指针，但不支持指针运算
  - 指针变量的值为内存地址
  - 未赋值的指针为nil

```go
type MyType struct {
	Name string
	Age  int
}

var m1 MyType
m1.Name = "Feng"
m1.Age = 18

fmt.Println(m1.Name, m1.Age)
```





### 函数

#### Main函数

- 每个Go语言程序都应该有个main package
- Main package里的main函数是Go语言程序入口

```go
package main
func main() {
	...
}
```



#### Init函数

- Init函数会在包初始化时运行
- 谨慎使用init函数--> 当多个依赖项目引用统一项目，且被引用项目的初始化在 init 中完成，并且不可重复运行时，会导致启动错误

```
var myinit = 0
func main() {
	fmt.Println("myinit", myinit)
}
func init() {
	myinit = 1
}
```



#### 返回值

- 多返回值

- 命名返回值
  - Go的返回值可以命名，它们会被当做定义在函数顶部的变量
  - 返回值的名称有一定的意义，它可以当做文档使用
  - 没有参数的return语句返回已命名的返回值，就直接返回
- 调用者忽略部分返回值
  -  a ,_ := returnfunc(args)



#### 可变参数

Go语言中允许传入

```
func main() {
	a, b, _ := RetrunSomething(1, 2, 3, 4, 5, 6, 9, 1)
	fmt.Println(a, b)
}

func RetrunSomething(args ...int) (a, b, c int) {
	fmt.Println("args ==>", args)
	a = 1
	b = 2
	c = 3
	return a, b, c
}

//args ==> [1 2 3 4 5 6 9 1]
//1 2
```



#### 内置函数

|       函数名        |            字作用             |
| :-----------------: | :---------------------------: |
|        close        |           管道关闭            |
|       len,cap       | 返回数组/切片/Map的长度或容量 |
|      new,make       |           内存分配            |
|    copy, append     |           操作切片            |
|   panic, recover    |           错误处理            |
|   print, println    |             打印              |
| complex, real, imag |           操作复数            |



#### 回调函数(callback)

- 函数作为参数传入其他函数，并在其他函数内部调用执行

```
func main() {
	execAdd(1, 2, Add)
}

func Add(a int, b int) {
	fmt.Println("result = ", a+b)
}

func execAdd(a int, b int, f func(int, int)) {
	f(a, b)
}
```



#### 闭包

- 匿名函数
  - 不能独立存在
  - 可以复制给其他变量
  - 可以直接调用
  - 可作为函数的返回值

#### 接口

- 接口定义一组方法集合
- 使用场景：kubernetes中有大量的接口抽象和多种实现
- Struct无需显示声明实现interface,只需直接实现方法
- Struct除实现interface定义的接口外还有额外的方法
- 一个类型可实现多个接口（go语言的多重继承）
- Go语言中接口不接受属性定义
- 接口可以嵌套其他接口

**Interface 是可能为 nil 的，所以针对 interface 的使用一定要预先判空，否则会引起程序 crash(nil panic)**

**Struct 初始化意味着空间分配，对 struct 的引用不会出现空指针**



#### 反射

- reflect.TypeOf() 返回被检查对象的类型
- reflect.TypeOf() 返回被检查对象的值



#### JSON编解码

- Unmarshal  从string --> struct
- Marshal 从struct --> string



### 常用语法

#### 错误处理

#### defer

- 函数返回之前执行某个语句或函数

- 常见的defer场景

  - defer file.Close()
  - defer mu.Unlock()
  - defer println("xxx")

  

#### Panic 和 recover

- panic: 在系统出现不可恢复错误时主动调用panic，panic会使当前线程直接crash
- defer :保证执行并把控制权交还给收到panic的函数调用者
- recover : 函数从panic或错误场景中恢复



### 多线程

#### 并发和并行

​	并发和并行的概念。并行：两个或多个事件在同一时间间隔发生；并行：两个或多个事件在同一时刻发生

#### 协程

- 进程

  - 分配系统资源（CPU时间、内存等）基本单位
  - 有独立的内存空间、切换开销大

- 线程：进程中的一个执行流，是CPU调度并能独立运行的基本单位

  - 同一进程中的多线程共享内存空间，线程切换代价小

  - 多线程通讯方便

  - 从内核层面来看线程其实也是一种特殊的进程，它跟父进程共享了打开的文件和文件系统信息，共

    享了地址空间和信号处理函数

- 协程

  - Go 语言中的轻量级线程实现

  - Golang 在 runtime、系统调用等多方面对 goroutine 调度进行了封装和处理，当遇到长时间执行

    或者进行系统调用时，会主动把当前 goroutine 的 CPU (P) 转让出去，让其他 goroutine 能被调度

    并执行，也就是 Golang 从语言层面支持了协程



#### 线程和协程的差异

- 每个goroutine(协程)默认占用内存远比Java、C的线程少
  - Goroutine: 2KB
  - 线程：8MB
- 线程/goroutine切换开销方面，goroutine远比线程小
  - 线程：设计模式切换（从用户态切换到内核态）、16个寄存器、PC、SP等寄存器的刷新
  - Goroutine: 只有三个寄存器的值修改 PC / SP /DX
- GOMAXPROCS: 控制并行线程数量



#### channel 多线程通讯

- Channel 是多个协程之间通讯的管道
  -  一段发送数据，一段接收数据
  -  同一时间只有一个协程可以访问数据，无共享内存模式可能出现内存竞争
  -  协调协程的执行顺序
- 声明方式
  - var identifier chan datatype
  - 操作符 <- 

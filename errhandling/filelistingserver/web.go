package main

import (
	"learnGo/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

//统一的错误处理逻辑

//统一的错误处理函数
type appHandler func(writer http.ResponseWriter, request *http.Request) error

//返回一个http.HandleFunc()需要的函数
//这个函数输入是一个函数,输出也是一个函数
//把输入的函数包装一下,包装成一个输出的函数
//这就是函数式编程,函数既能做参数,又能作为返回值
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {
		//定义一个延时执行的函数,捕捉错误信息
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			//将具体的系统的错误信息在命令行中输出
			log.Printf("Error ocurred"+"handling request: %s", err.Error())

			//如果错误信息假定为userError类型,给用户展示用户的错误信息
			if userErr, ok := err.(userError); ok {
				http.Error(
					writer,
					userErr.Message(),
					http.StatusBadRequest,
				)
				return
			}

			//匹配错误信息,将不必要的信息过滤,不展示给前台看
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(
				writer,
				//这里写err.Error()也可以但是会把我们内部的错误报漏出去
				http.StatusText(code),
				code,
			)
		}
	}
}

//给用户看的错误信息
//go语言的接口,两边是各自管各自的定义的,这边定义的userError只是说我要实现一个error和message()
//然后那边用的人呢,自己去管实现自己的userError,就能给外面的人用得到,这两个接口的本身是不需要相互看得到的
type userError interface {
	error            //给系统看的错误信息
	Message() string //给用户看的错误信息
}

func main() {

	//设置文件访问的路径, 原本第二个参数是一个回调函数, 我们之前直接把回调函数写在后面的, 现在剥离出去了, 因为业务部分写在这里不合适
	//这里我们把/list/改成"/",filelisting.HandleFileList保持不变, 访问的时候就会报错,此时访问时,里面的函数,就是提示slice越界
	//因为我们访问的时候不加/list/直接用文件名,他会URL进行剪切了/list/然后对剩下的字符串进行匹配,在用改字符串访问文件,现在的slice肯定不够长了
	//虽然越界导致错误,但是这个httpserver没有挂, 因为httpserver做了保护,使用recover做的保护,但是还会panic()报错,我们不想报错,所以在上面使用recover判断加上了保护
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	//开启监听服务器
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

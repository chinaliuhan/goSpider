package main

import (
	"learnGo/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
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
		err := handler(writer, request)
		if err != nil {
			//将具体的系统的错误信息在命令行中输出
			log.Printf("Error ocurred"+"handling request: %s", err.Error())
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

func main() {

	//设置文件访问的路径, 原本第二个参数是一个回调函数, 我们之前直接把回调函数写在后面的, 现在剥离出去了, 因为业务部分写在这里不合适
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	//开启监听服务器
	err := http.ListenAndServe(":8881", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func simpleHttpDemo() {
	response, error := http.Get("http://www.baidu.com")
	//有错误处理错误
	if error != nil {
		panic(error)
	}
	//body资源要求必须关闭
	defer response.Body.Close()

	//处理返回值,第二个值标识是否打印body
	byte, error := httputil.DumpResponse(response, true)
	if error != nil {
		panic(error)
	}
	fmt.Printf("%s\n", byte)
}

func httpControllerDemo() {
	//对请求内容进行控制
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}
	//设置请求头
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3493.3 Safari/537.36")
	//发送请求
	//response, error := http.DefaultClient.Do(request)	//默认发送请求
	//自定义请求内容
	//可以设置transport代理服务器 CheckRedirect重定向会走这里  jar设置cookie  timeout
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	response, error := client.Do(request)

	//有错误处理错误
	if error != nil {
		panic(error)
	}
	//body资源要求必须关闭
	defer response.Body.Close()

	//处理返回值,第二个值标识是否打印body
	byte, error := httputil.DumpResponse(response, true)
	if error != nil {
		panic(error)
	}
	fmt.Printf("%s\n", byte)

}

func main() {
	//simpleHttpDemo()
	httpControllerDemo()
	//使用http客户端发送请求
	//使用http.Client控制请求头部
	//使用httputil简化工作

	//http服务器的性能分析,要用的话首先需要一个http的服务器
	//import _ "net/http/pprof" 这里的下划线标识,这个引用我可以使用
	//访问/debug/pprof/
	//使用go tool pprof分析,获取30秒内的CPU性能分析
	// 比如在当前web服务下的命令行中 go tool pprof http://localhost:8888/debug/pprof/profile  把profile改成heap可以看内存使用情况,具体可以看goroot里面的pprof或者手册
	//访问这个URL后, 会在当前命令行中打印很多东西, 输入web 通过浏览器打开
}

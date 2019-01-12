package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	//请求页面
	response, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	//关闭资源
	defer response.Body.Close()
	//会打印头部之类的信息
	//httputil.DumpResponse()
	//判断头部
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", response.StatusCode)
		return
	}
	//下载官方库,这个库可以用来为HTML页面转码
	//依赖gopm get -g -v golang.org/x/text
	//转码,将GBK转码为UTF8
	//utf8Reader := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	//使用封装的这个来自动判断编码
	e := determineEncoding(response.Body)
	utf8Reader := transform.NewReader(response.Body, e.NewDecoder())

	//只要body
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)

}

//自动检测编码
func determineEncoding(r io.Reader) encoding.Encoding {
	//如果把response.body直接放进来的话,读完后面就没法再读了,所以这里我们单独处理一下
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	//自动检测编码
	//依赖gopm get -g -v golang.org/x/net
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

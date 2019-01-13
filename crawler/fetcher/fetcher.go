package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
//通过URL下载对应的页面并自动转码为utf-8
func Fetch(url string) ([]byte, error) {
	//处理URL返回处理完的具体内容,和错误信息

	//请求页面 "http://www.zhenai.com/zhenghun"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//关闭资源
	defer response.Body.Close()
	//会打印头部之类的信息
	//httputil.DumpResponse()
	//判断头部
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", response.StatusCode)
	}

	//下载官方库,这个库可以用来为HTML页面转码
	//依赖gopm get -g -v golang.org/x/text
	//转码,将GBK转码为UTF8
	//utf8Reader := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
	//使用封装的这个来自动判断编码
	e := determineEncoding(response.Body)
	utf8Reader := transform.NewReader(response.Body, e.NewDecoder())

	//只要body
	return ioutil.ReadAll(utf8Reader)
}

//自动检测编码
func determineEncoding(r io.Reader) encoding.Encoding {
	//如果把response.body直接放进来的话,读完后面就没法再读了,所以这里我们单独处理一下
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("fetcher error:%v", err)
		//返回默认编码
		return unicode.UTF8
	}
	//自动检测编码
	//依赖gopm get -g -v golang.org/x/net
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

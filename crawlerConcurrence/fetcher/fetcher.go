package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//请求频率的限制
var rateLimiter = time.Tick(10 * time.Millisecond)
//通过URL下载对应的页面并自动转码为utf-8
func Fetch(url string) ([]byte, error) {
	//要限制请求频率,这里不用做别的限制,因为time.Tick返回的本身就是一个不断递减的channel
	//利用channel的阻塞特性,放这里就行
	<-rateLimiter

	//处理URL返回处理完的具体内容,和错误信息

	//请求页面 "http://www.zhenai.com/zhenghun"
	//response, err := http.Get(url)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//request.Header.Add("Accept-Encoding", "gzip, deflate")//不能有这个,解不开
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	request.Header.Add("Cache-Control", "no-cache")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
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
	bodyReader := bufio.NewReader(response.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	//只要body
	return ioutil.ReadAll(utf8Reader)
}

//自动检测编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//如果把response.body直接放进来的话,读完后面就没法再读了,所以这里我们单独处理一下,只用peek拿出前面的1024个
	//bytes, err := bufio.NewReader(r).Peek(1024)
	//上面读完后面1024后,后面的内容会少1024,所以这里改成这样,同时将 bufio.NewReader拿到外面去
	bytes, err := r.Peek(1024)
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

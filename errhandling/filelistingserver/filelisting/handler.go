package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

//这里因为我们要用到userError,所以这里我们也定义一个userError
//go语言的接口,两边是各自管各自的定义的,那边定义的interface只是说他要实现一个error和message()
//但是对这边是不影响的,如果这边要用就要实现自己的userError
type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//如果不存在/list/提示错误,避免内存越界
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must start with" + prefix)
		//return errors.New("path must start with" + prefix)
	}

	//获得请求的URL地址,通过len()计算出/list/的长度,然后进行剪切得到URI文件地址
	path := request.URL.Path[len(prefix):]
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		//处理http错误号,如果直接使用panic()的话,用户在访问不存在的文件时,会报错,没有任何提示
		//在外面, 我们专门做了错误的处理, 所以这里直接返回错误,在外面处理
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	//关闭资源
	defer file.Close()

	//读取文件的所有内容
	all, err := ioutil.ReadAll(file)
	if err != nil {
		//返回错误,让外面处理
		return err
	}

	//直接输出文件内容
	writer.Write(all)
	return nil
}

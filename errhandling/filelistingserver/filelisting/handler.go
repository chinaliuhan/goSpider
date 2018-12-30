package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//获得请求的URL地址,通过len()计算出/list/的长度,然后进行剪切得到URI文件地址
	path := request.URL.Path[len("/list/"):]
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

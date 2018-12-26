package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//将整数转换为2进制
func convertTobin(n int) string {
	//省略起始条件
	//使用for循环把整数转换为2进制的表达式,用数学中的短除法转换为2进制
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//读取文件内容打印
func printFile(fileName string) {
	fmt.Println(os.Getwd())
	//这里的这个路径是按照项目目录走的, 并非当前文件的相对路径
	//比如我之前在这里直接传入aaa.log,会提示no such file or directory
	//通过os.Getwd()打印知道GO的路径为项目根路径
	//所以"basic/loop/aaa.log"这种全路径
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

//go语言中的Reader和Writer不仅可以操作文件,网络还有string,byte,slice 都是作为Reader和Writer
//我们在点开os.Open的时候会发现返回的是一个file的结构体里面就有一个Read()就是对Read()接口的一个实现
//而且我们发现bufio.NewScanner()里面要的就是一个Reader,通常情况下我们不用file,而是什么都用Reader和Writer
//Reader和Write更加关键的地方是,fmt.Fprintf(),fmt.Fscanf()里面的第一个参数就是一个reader,这时候我们给网络连接啊什么的都是可以用的
//针对Reader和Writer各写一个接口之后,可以装在文件上也可以装在网络上.非常的方便
//因此我们在比较底层的读写相关的东西都希望把他做成Reader和Writer,这样我们就可以和许多系统的函数一起共用
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	//go 语言没有while,这个就是while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	//省略初始条件和递增条件,就是一个死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	//将10进制的正整数转换为2进制
	fmt.Println(
		convertTobin(5),
		convertTobin(13),
		convertTobin(72387885),
		convertTobin(0),
	)
	//读取文件内容,输出
	printFile("basic/loop/aaa.log");
	//在go语言中反引号代表跨行的字符串
	s := `abc"d"
kkkk
123

pppp`
	//像文件一样打印字符串,除了strings.NewReader之外还有bytes.NewReader两个差不多只是用的类型不一样
	printFileContents(strings.NewReader(s))

	//forever();
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	scanner := bufio.NewScanner(file)
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
	forever();
}

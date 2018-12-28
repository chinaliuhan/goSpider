package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//斐波那契数列 生成器,返回一个无参数的函数
//func fibonacci() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//要实现接口要先实现类型,类型是intGen 内容是function() int
type intGen func() int

//把斐波那契数列生成器的数字当做文件一样处理
//函数也能实现接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	//如果大于10000,返回0,告诉Read 文件已经到头
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

//为函数实现接口
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {

	f := fibonacci()
	/**

	fmt.Println(f()) //1
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //5
	fmt.Println(f()) //8
	fmt.Println(f()) //13
	fmt.Println(f()) //21
	*/

	fmt.Println("--------------\n")
	printFileContents(f)
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func demoIf() {
	//if 的条件语句可以赋值
	//if的条件赋值的变量作用域就在这个if语句里
	//if不需要括号

	const filename = "abc.txt"
	//go的函数可以返回两个值
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//
	//	fmt.Printf("%s\n", err) //输出错误信息
	//} else {
	//
	//	fmt.Printf("%s\n", contents) //输出文件内容
	//}
	//也可在这里输出
	//fmt.Printf("%s\n", contents)

	//这里的if else 也可以这么写,不过这种形式写的话, 在PHP中会被打死.
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("%s\n", err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//这里就不可用在这里输出,因为这种写法的contents, err在if中定义的,他的生存期只在if的这个block中才有
	//fmt.Println(err)
}

func eval(a, b int, op string) int {
	//switch 不需要加break, 如果不想break反而要用 fallthrough
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupport operator:" + op)
	}
	return result
}

func grade(score int) string {
	//go语言的switch case命中之后会默认自动进行break
	g := ""
	//switch 后面可以没有表达式,只要在case中定义了就行
	switch {
	case score < 0 || score > 100:
		//panic 中断程序执行进行报错
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"

	}
	return g
}

func demoFor() {
	sum := 0
	//for循环不能有括号
	//for的条件里可以省略初始条件,结束条件,递增表达式 go语言中没有while 省略了初始条件,结束条件,就可以当做while用
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func convertToBin(n int) string {
	//省略起始条件
	//使用for循环把整数转换为2进制的表达式,用数学中的短除法转换为2进制
	result := ""
	for ; n > 0; n /= 2 {
		lasb := n % 2
		result = strconv.Itoa(lasb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	//for的条件里可以省略初始条件,结束条件 go没有while语句, 这样就可以当做while来用
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forEver() {
	//for省略初始条件,结束条件,递增表达式
	//这里会形成一个死循环输出
	for {
		fmt.Println("abc")
	}

}

func main() {

	//demoIf()
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		//下面这个就会中断程序执行就行报错
		//grade(101),
	)
	demoFor()
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
	)
	printFile("abc.txt")
	//forEver()
}

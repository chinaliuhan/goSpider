package main

import (
	"fmt"
	"math"
)

//函数外面的变量只是一个包变量, 也只是作用于包内部,GO语言没有全局变量这个说法
var aa = "3"
var ss = "kkk"

//函数外禁止使用:=定义变量 在函数外,即包内定义时必须以一个关键字开头
var bb = true

//常量也可以定义在包内,即函数外面, 这样的话所有的包内的函数都可以用一组常量也可以像变量那样用()来声明
//go语言的常量最好不要进行全大写,go语言中首字母大写是有其他含义的
const demo = "demo"

//函数内和包内都可以, 一次定义多个变量
//var (
//	aa = "3"
//	ss = "kkk"
//	bb = true
//)

func variableZeroValue() {
	//定义变量, 变量名在前,类型在后, 变量定义后必须使用,不用不行
	//定义变量时不设置值, 都会有默认值, int:0 string: ""
	var a int
	var s string
	//该函数就像PHP中的printf , 需要import "fmt"
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	//同时定义多个变量时,也可以不赋初始值
	//规定类型后,不同类型的值不可以写在一行
	var a, b int = 1, 2
	var s string = "abc"
	fmt.Println(a, b, s)

}

//编译器可以推测变量的类型
func variableTypeDeduction() {
	//在不规定类型的情况下, 可以将多个类型写在一行
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	//在函数内可以,使用短标签 := 进行定义变量,但是在后面赋值时不可再用 := 只能用=
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	//注意go的类型转换必须是显式的,没有隐式的
	// 这里的结果不加手动int() 会报cannot use math.Sqrt(float64(a * a + b * b)) (type float64) as type int in assignment
	//这里为什么要先把输入float()再把结果int() 是因为math.Sqrt只接受float64并且返回的也是float64而我们c是int
	//但是int的变量c 没法直接 = float64 所以要把结果显式的int转换一下
	c = int(math.Sqrt(float64(a*a + b*b)))
	//错误的写法 错误原因就是上面说的那样
	//var c int = math.Sqrt(a*a + b*b)
	fmt.Println(c)
}

//常量
func consts() {
	//常量的数值可以作为各种类型使用,所以下面计算的时候需要float时可以自动转换
	//常量可以规定类型也可以不规定类型, 不规定类型时为不确定类型,常量不像变量,常量定义后key不用,不用也不会报错
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	//因为上面没有定义变量类型, 这里会认为a,b既可以是int也可以是float,因为这里会认为是一个float 不用转换
	//但是c因为规定了变量类型, 所以这里必须进行转换
	c = int(math.Sqrt(a*a + b*b))
	//常量不像变量,常量定义后key不用,不用也不会报错
	fmt.Println(c, filename)

}

//枚举类型
func enum() {

	//一组常量放在一起就是一个枚举类型
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)

	//下面使用一个关键字iota表示这一组常量是一个自增的枚举 这里可以使用一个 _ 进行站位
	//这里的iota其实是一个表达式,下面有更加复杂的用法,iota还可参与运算,作为自增值的种子
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)
	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//该函数就像PHP中的echo , 需要import "fmt"
	//fmt.Println("hello word")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(aa, ss, bb)
	//triangle()
	//consts()
	enum()
}

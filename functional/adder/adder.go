package main

import "fmt"

//函数是编程 VS函数指针
/**
函数是一等公民: 参数,变量,返回值都可以是函数
高阶函数  他的提法就是他的参数还可以是一个函数
函数->闭包
*/

//正式编程
/**
不可变性: 不能有状态,只有常量和函数
函数只能有一个参数
本课程不做上述严格规定
*/

//通过闭包来做累加,该函数返回一个闭包
func adder() func(int) int {
	//闭包特征之一,,自由变量
	sum := 0
	//闭包特征之一,这里的这个v是一个参数,也是一个局部变量
	return func(v int) int {
		//闭包特征之一,sum不是在函数体里面定义的,他是这个函数所处的这个环境,sum是外面的
		sum += v
		return sum
	}
}

//稍微正统式函数编程的写法
//返回一个新的函数,递归定义
type iAddr func(int) (int, iAddr)

//接收一个base返回的是一个iAdder即下一个函数
func addr2(base int) iAddr {
	//传入一个参数v,返回一个整型和iAdder函数
	return func(v int) (int, iAddr) {
		//返回一个部分和,和一个递归
		return base + v, addr2(base + v)
	}

}

func main() {

	//a := adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("0 + 1 + ...... + %d = %d\n", i, a(i))
	//}

	//为实例赋值,同时初始化数据
	a := addr2(0)
	for i := 0; i < 10; i++ {
		//部分和
		//这样赋值变量,没用s,a:=a(i)是因为a不用的话会报错
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ...... + %d = %d\n", i, s)
	}

}

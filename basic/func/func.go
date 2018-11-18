package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数定义, 函数名在前, 返回类型在后 可以传入一个匿名函数,可传入可变参数列表,可以返回多个参数,没有默认参数,可选参数,函数重载,操作符重载,
//参数用逗号分隔, 参数也是一样名在前,类型在后 这里发红是因为在别的文件中也声明了该名字的函数
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		//返回两个值的时候, 如果只想用一个的话, 另一个值可以用 _ 来做代表,意思就是不接受这个值
		q, _ := div1(a, b)
		return q
	default:
		panic("unsupport operator:" + op)
	}
}

//除法 返回多个值, 多返回值不要乱用, 一般就是返回一个值,再返回一个错误信息
func div(a, b int) (int, int) {
	return a / b, a % b
}

//在返回多个值的情况下, 给返回值定义名称,但是在起名字的, 需要好好想想, 要不然散在函数体中, 搞不清哪里调用,会很难读
func div1(a, b int) (q, r int) {
	//返回两个值的时候, 如果只想用一个的话, 另一个值可以用 _ 来做代表,意思就是不接受这个值
	//里面不动,但是外面的返回值也需要用名字来接一下, 这里的q,r只是在函数体中的名字,在外面接的时候可以随意定,叫a,b也无所谓,但是最好保持一致
	//这里还可以这样写
	//return a / b, a % b
	//但是最好这样写, 因为已经返回值的名称了定义了, 但是一旦函数体比较长, 这样写就会很难读, 比较长的时候, 最好就用上面那种
	q = a / b
	r = a % b
	return
}

//多返回值不要乱用, 一般就是返回一个值,再返回一个错误信息,这样的话函数就不会中断
func eval1(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//返回两个值的时候, 如果只想用一个的话, 另一个值可以用 _ 来做代表,意思就是不接受这个值
		q, _ := div1(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupport operator: %s", op)
	}
}

//go语言是一个函数式编程语言, 他的函数中可以套函数 作为复合函数
//主要流程, 我们把收到的a,b参数放到op()中执行 然后返回一个Int
//apply中的函数, 我们把收到的两个a,b参数放在第一个函数op()中执行,op()返回的是一个int
//细节 ope()接收两个参数, 返回一个int
//applay 接收两个参数a,b返回一个int
func apply(op func(int, int) int, a, b int) int {
	//输出函数调用名称
	fmt.Printf("Calling %s with %d, %d \n", runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(), a, b)
	//将apply接收到的参数放到op()中执行
	return op(a, b)
}

//可变参数, 任意传入多少个参数, 取参数的时候就像数组一样去用
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

//指针 注意go语言只有值传递没有引用传递

//指针类型,本函数 将参数a,b 的值互换为b,a 必然会失败
func swap(a, b int) {
	a, b = b, a
}

//想要正常换可以采用指针类型
func swap1(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {

	fmt.Println(eval(3, 4, "*"))
	fmt.Println(eval(3, 4, "/"))
	fmt.Println(div(13, 3))
	q, r := div1(13, 3)
	fmt.Println(div1(q, r))

	if result, err := eval1(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	//也可以通过传入一个匿名函数的方式来调用
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	//用跟这样的方式,输出结果不会互换
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)

	//通过采用指针类型,互换两个参数的值,想要更换两个参数的值, 必须要传递一个指针进去
	c, d := 3, 4
	swap1(&c, &d)
	fmt.Println(c, d)

	//如果不想用指针的话,可以反着,返回两个参数,即可
	e, f := 3, 4
	e, f = swap2(e, f)
	fmt.Println(e, f)

}

package main

import "fmt"


/**
panic
停止当前函数运行
一直向上返回,执行每一个defer
如果没有遇到recover 程序退出

recover
仅在defer中调用
在dfer中可以获取panic的值
如果无法处理,可重新panic
 */

func tryRecover() {

	//匿名函数后面必须有一个()
	defer func() {
		r := recover()
		//如果是一个错误信息,输出错误信息,只是一行错误信息,如果没有这个错误处理的话,整个输出都很难看
		//这里用r.(error)  是类型断言，意思就是，我推断 r是一个error 的类型
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			//如果不知道是一个什么东西,不知道怎么处理,直接输出
			panic(fmt.Sprintf("I don't know what to do%s", r))
		}

	}()

	//panic(errors.New("this is an error"))

	//比如现在这里就会输出Error occurred runtime error: integer divide by zero
	//否则就是一大堆的代码错误信息
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	//而这个直接输出的错误信息,则是上面的recover无法捕捉的
	//所以直接走到 panic(fmt.Sprintf("I don't know what to do%s", r))
	panic(123)

}

func main() {
	tryRecover()
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//defer 确保调用结束时发生, 同时多个defer 时遵循现金后出的原则
//defer 参数在defer语句时计算
//defer	列表为后进先出

func tryDefer() {
	//输出结果3 panic: error occurred 2 1
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occurred")
	//fmt.Println(4)

	//这里在输出的时候是从30,29...0倒过来算的, 可以发现,参数在defer时已经计算完毕,而且先进后出
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}

}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	//创建文件
	//file, err := os.Create(filename)
	//if err != nil {
	//	panic(err)
	//}

	//这里制造点错误, 作为错误处理的案例
	//这里的os.OpenFile,点进去看的时候发现If there is an error, it will be of type *PathError.
	//所以下面在处理错误的时候,我们使用了*os.PathError来处理
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//这里使我们自定义声明的错误,目的是让err不是PathError, 以触发下面的panic
	err = errors.New("this isa custom error");
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf(
				"%s,%s,%s",
				pathError.Op,
				pathError.Path,
				pathError.Err,
			)
		}
		return
	}

	//函数结束时,关闭资源
	defer file.Close()
	//将数据保存在内存中,比上面的file快
	write := bufio.NewWriter(file)
	//	//函数结束时,保存数据
	defer write.Flush()
	f := fibonacci()
	for i := 0; i < 20; i++ {
		//功能同上面三个函数，只不过将转换结果写入到 w 中。
		fmt.Fprintln(write, f())
	}
}

func main() {
	//tryDefer()
	//将斐波那契的输出,保存在项目根目录下
	writeFile("fib.txt")
}

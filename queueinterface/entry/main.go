package main

import (
	"fmt"
	"learnGo/queueinterface"
)

func main() {
	q := queueinterface.Queue{1}

	//queueinterface.Queue改为泛型之后,这里就可以push和pop字符串了
	q.Push(2)
	q.Push(3)
	q.Push("abc")
	q.Push("def")
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

}

package queue

import "fmt"

//Queue内的Pop的示例代码
func ExampleQueue_Pop() {

	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//下面的Output:内容是作为测试时,我们期望的返回值,而且是必须要有的,而且这两行的注释也会被放到示例代码中展示
	//示例代码不仅仅是示例代码,而且会检查Output的值是不是对的, 下面这个是故意写错的
	//Output:
	//2
}

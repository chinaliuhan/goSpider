package queue

import "fmt"

//Queue内的Pop的示例代码
func ExampleQueue_Pop() {

/**
//下面的Output:内容是作为测试时,我们期望的返回值,而且是必须要有的,而且这两行的注释也会被放到示例代码中展示
//示例代码不仅仅是示例代码,而且会检查Output的值是不是对的, 下面这个是故意写错的
//而且这里的注释不能够乱写, 这里写的注释都会被认为是文档的一部分
//注释如果写在Output:上面这个函数将不不能在IDE上点击左侧的箭头运行,箭头直接消失了
//所以我把 注释写在了函数内的第一行这里
//Output:下面正确的写法是应该写正确的结果,如下
1
2
false
3
true
 */


	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//Output:
	//1
	//2
	//false
	//3
	//true
}

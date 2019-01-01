package queue

//通过定义别名来扩展其他的包
type Queue []int

//Pushes the element into the queue
//	e.g q.Push(123)
func (q *Queue) Push(v int) {
	//这里的*q是一个指针接受者, 指针接受者时可以改变里面的值的
	// 后面使用*q 前面也要加一个*
	*q = append(*q, v)
}

//Pops element from head
func (q *Queue) Pop() int {
	//这里的括号是定界符
	//这里是取出第0个元素进行返回
	head := (*q)[0]
	//将q的内存地址中的数据, 修改为将q的下标为1开始剪切到最后一个的数据
	// 后面使用*q 前面也要加一个*
	*q = (*q)[1:]
	return head
}

//Returns if the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0

}

//文档查看与生成
//在函数的上边写的注释,生成文档

//命令行版查看文档
//在当前目录下使用go doc 可以查看的到Queue函数
//如 go doc 如果要查看该实体的函数go doc 函数名,如: go doc push 和go doc Queue
//使用go help doc 还可以看到go的包里面的函数说明,比如go doc fmt.Println
//通过起一个网络服务还查看文档
//godoc -http :6060 会起一个服务,然后通过localhost:6060可以查看一个页面,和官网差不多,有文档和包的说明
//该页面还可以看到我们在项目中写的函数的文档, 不过我们的函数需要有注释才会有文档,否则即只有函数
//函数上面的注释同一行后面的注释,就会被作为文档展示,注释没有指定格式,任意写
//如果在注释的下一行,在开一行注释,并且开头4个空格,会被认为是代码案例展示
//示例代码
//示例代码在queue_test中
//在我们补充完示例代码之后,在王网络服务查看文档时,该函数会有一个Example点开之后是示例代码,里面含有Code和Output
//但是如果//Output:下面的返回内容写的不对的话, 文档汇总是不会显示Output的
//如果要测试queue_test的话可以用IDE,点击函数名左侧的箭头,也可以在当前目录下使用go test来做

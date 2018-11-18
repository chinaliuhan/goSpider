package queue

//通过定义别名来扩展其他的包
type Queue []int

func (q *Queue) Push(v int) {
	//这里的*q是一个指针接受者, 指针接受者时可以改变里面的值的
	// 后面使用*q 前面也要加一个*
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	//这里的括号是定界符
	//这里是取出第0个元素进行返回
	head := (*q)[0]
	//将q的内存地址中的数据, 修改为将q的下标为1开始剪切到最后一个的数据
	// 后面使用*q 前面也要加一个*
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0

}

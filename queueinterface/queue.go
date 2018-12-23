package queueinterface

//通过定义别名来扩展其他的包
//这里使用Interface 表示支持任意类型
type Queue []interface{}

//这里使用Interface 表示支持任意类型
//如果要坚持使用int 这里还用int就行
func (q *Queue) Push(v interface{}) {
	//这里的*q是一个指针接受者, 指针接受者时可以改变里面的值的
	// 后面使用*q 前面也要加一个*
	//如果这里要类型限制的话, 也可以通过v.(int)来做转换
	*q = append(*q, v)
}

//如果这里还要用int 不仅需要把这里的interface{}改为int 还需要返回值改为return head.(int)
func (q *Queue) Pop() interface{} {
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

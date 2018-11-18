package container

import "fmt"

func printSlice(s []int) {
	fmt.Println("%v,len=%d, cap=%d\n", s, len(s), cap(s))

}

func main() {

	fmt.Println("creating slice")
	//声明一个slice,并不设置初始值
	// go语言有一个 zero value for slice is nill
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	//定义slice,长度为16
	s2 := make([]int, 16)
	//定义slice长度为16,cap为32
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
	//
	fmt.Println("Coping slice")
	//s2为复制的目标 s1为数据来源切片。下面就是将s1拷贝到s2
	copy(s2, s1)
	printSlice(s2)

	//删除中间的元素
	fmt.Println("Deleting elements form slice")
	//go语言没有专门的删除,这里以删除s2中的下标为3的值为例
	//这里利用append会对slice生成一个新的slice, 可以先取原s2的0-3 在把4-最后的拼在一起即可
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	//删除头部第一个元素
	fmt.Println("Popping form front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	//删除尾部第一个元素
	fmt.Println("Popping form back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)

}

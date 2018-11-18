package container

import "fmt"

func updateSlice(s []int) {

	s[0] = 100
}

//在另一个地方我们是使用指针操作的 这里可以改成传入一个slice
func printArr2(arr []int) {
	//这里也可以这么用, k值直接用 _ 占位符省略变量, 后面就可以不输出了下标, 值打印变量了
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}

//slice go语言的所有类型都是一个值类型, 但是slice不是一个值类型,slice本身是没有数据的, 他的肚子里面有一个数据结构
//slice是数组的一个view(视图)
func main() {
	//go语言的切换, 从X开始到Y,值得是下标, 且不包括Y
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//切片 2-6
	fmt.Println("arr[2:6]=", arr[2:6])
	//从开始 到6
	fmt.Println("arr[:6]=", arr[:6])
	//从2到最后
	fmt.Println("arr[2:]=", arr[2:])
	//取所有
	fmt.Println("arr[:]=", arr[:])

	//
	s1 := arr[2:]
	fmt.Println("arr[2:]", s1)
	s2 := arr[:]
	fmt.Println("arr[2:]", s2)

	//下面的输出结果是 ,可以看到s1里面的第一个值, 即arr中的下标为2的已经被修改为了100
	//我们在输出s1的时候会发现s1的第一个值已经被修改了,而arr中的小标为2的也被修改了,这里的s1是arr的一个slice
	//这里就可以证明slice是数组的一个view,对新数组的操作会影响到原来的数组
	// after update slice
	//[100 3 4 5 6 7]
	//[0 1 100 3 4 5 6 7]
	fmt.Println("after update slice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	// 下面的结果会输出,和上面的操作的理论上是一样的
	//after update slice(s2)
	//[100 3 4 5 6 7]
	//[0 1 100 3 4 5 6 7]
	fmt.Println("after update slice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	//这里我们本来要用指针来做到类似于引用传值的效果
	//而这里我们可以直接通过slice而不用指针
	fmt.Println("after update slice(s3)")
	printArr2(arr[:])

	//
	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[5:]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//slice的扩展,
	fmt.Println("Extending slice")
	fmt.Println(arr)
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	//这里是会直接报错, 因为slice是数组的一个view,在view的时候会将后面的值打印,直接用[]取值是取不到的但是再次通过slice取是可以取到的
	//在对数组进行, slice后会,有一个截取后开头的叫ptr,本身数组的len,还有一个cap包括从截取后的ptr到原数组的最后,
	//在对新的slice再次进行slice时不超过ptr的范围都可以去得到,也就是说可以向后扩展, 但是不能向前扩展
	//s[i]不可超过len(s),向后扩展不可超过底层数组cap(s)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) //取出对应的长度
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2)) //取出对应的长度
	//fmt.Println("s1[4]=", s1[4])                                        //取不到,报错
	fmt.Println("s1=", s1) //s1= [2 3 4 5]
	fmt.Println("s2=", s2) //s2= [5 6]

	fmt.Println("append\n")
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5=", s3, s4, s5)
	//s4 and s5 no longer view arr
	//s4 s4 不再是对arr的一个view,这时候系统会分配一个更加长的array来添加我们这么多的数据
	// 添加数据时如果超过cap,系统会重新分配更大的底层数组,并将原来的元素拷贝过去,go是有垃圾回收机制的,如果原来的数组还有人用 则保留,否则就会被垃圾回收掉
	//由于值传递的关系,必须接受append的返回值,append的时候slice的len可以会改变,他的cap,ptr也可能会改变,所有我们要用新的slice来接这个返回值
	fmt.Println("arr=", arr)
}

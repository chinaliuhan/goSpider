package main

import (
	"fmt"
)

//这里我们直接用channel来通信,告诉外面我们这里做完了
func doWorker(id int, c chan int, done chan bool) {
	for {
		n := <-c
		fmt.Printf("worker %d received %c \n", id, n)
		//通知打印完成,外面再用这个参数来判断里面运行完毕了
		//发送者角度：对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的。如果chan中的数据无人接收，就无法再给通道传入其他数据。因为新的输入无法在通道非空的情况下传入。所以发送操作会等待 chan 再次变为可用状态：就是通道值被接收时（可以传入变量）。
		//接收者角度：对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
		//go func() {done <- true}()
		done <- true
	}
}

//声明一个结构体,结构体里两个channel类型的成员属性, 将in作为传值,done作为内外通信
type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func chanDemo() {
	//这里我们使用channel通信的方式,在内部完毕的时候通过外面,外面在继续运行,使main函数不退出,不用再用sleep让main函数等待

	//下面是一种阻塞形式,两种并行方式进行

	//通过worker实例,来做通信
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//这里的读取没有用,只是为了把doWorker里面写入的值拿出来,然后进入下一个循环
		//在worker的成员属性写入a+i之后,马上把doWorker里面写入的done的数据读取出来
		//这样的形式确实是不用sleep了,但是编程阻塞运行了...一进一出,channel的通信是block阻塞的
		//所以我们把读取数据单独放到另外一个for循环里面
		//<-workers[i].done
	}

	//下面那种把两个一起取出来的的并发比较负责
	//这里的这种比较简单的形式,让大A和小a分开并发执行
	//先并发执行小a,再并发执行大A,里面的done <- true就不用改成改成go func() {done <- true}()
	for _, worker := range workers {
		<-worker.done
	}

	//也可以这么写
	for i, work := range workers {
		work.in <- 'A' + i
		//这里的读取没有用,只是为了把doWorker里面写入的值拿出来,然后进入下一个循环
		//在worker的成员属性写入a+i之后,马上把doWorker里面写入的done的数据读取出来
		//这样的形式确实是不用sleep了,但是编程阻塞运行了...一进一出,channel的通信是block阻塞的
		//所以我们把读取数据单独放到另外一个for循环里面
		//<-work.done
	}

	//下面那种把两个一起取出来的的并发比较负责
	//这里的这种比较简单的形式,让大A和小a分开并发执行
	//先并发执行小a,再并发执行大A,里面的done <- true就不用改成改成go func() {done <- true}()
	for _, worker := range workers {
		<-worker.done
	}


	//这样就可以并发执行了
	//但是会有一个问题,里面的done在运行的时候,channel是阻塞的,正在循环等待外面的人接收
	//外面的小a还没收到,大A就开始发送了,他发的这个大写的A的task就回到了下一个for里面因此就卡在了那个
	//我们有一个快速的方法就是把里面的done <- true改成go func() {done <- true}()再开一个goroutine让他去并行的发送数据,我们就不会卡住了
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//}

}

func main() {
	chanDemo()
}

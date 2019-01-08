package main

import (
	"fmt"
	"sync"
)

//这里我们直接用sync.WaitGroup来通信,告诉外面我们这里做完了
func doWorker(id int, w worker) {

	//for n := range w.in {
	//	fmt.Printf("worker %d received %c \n", id, n)
	//	w.done()
	//}
	for {
		n := <-w.in
		fmt.Printf("worker %d received %c \n", id, n)
		w.done()
	}
}

//声明一个结构体,结构体里两个channel类型的成员属性, 将in作为传值,done作为内外通信
type worker struct {
	in chan int
	//这里要用指针,因为他是一个引用,而因为我们要用外面的wg,不能说我们来拷一份
	//其他地方也只能传一个指针进去,wg只能用同一个
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {

	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func chanDemo() {
	//sync.WaitGroup的库
	//go语言的并发执行,等待多人来完成任务这个go语言的库提供了一个方法,叫做sync.WaitGroup的库
	//sync.WaitGroup.Add(n)添加计数,有几个协程n就写几
	//sync.WaitGroup.Wait()阻塞直到计数为零,让进程等待协程执行完毕
	//sync.WaitGroup.Done()减掉计数，等价于Add(-1)，这样sync.WaitGroup只有两个API了, 在协程内部使用,运行完毕时用欧冠这个方法,告诉外面该方法运行完毕了

	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	//这里这个也可以放到下面的for循环中每次wg.Add(1)
	//add是说我们有多少个任务要并行
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//也可以这样,但是既然我们知道一共是20个,所以不如直接在上面写20个算了
		//wg.Add(1)
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	//上面上面的任务全部做完
	wg.Wait()
}

func main() {
	//针对刚才的channelWaitgroup进行函数式编程的一点优化
	//将sync.WaitGroup.done封装取来,doWorker在运行的结束的时候,直接调用worker.done()即可
	//也只是让代码的抽象程度更高一些而已
	chanDemo()
}

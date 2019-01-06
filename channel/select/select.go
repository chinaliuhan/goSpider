package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//sleep 1500毫秒以内的时间,然后读取将0的自增写入到out中
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func test1() {
	var c1, c2 chan int //c1 and c2 =nil 初始化定义还没有赋值,所以会是默认值 nil channel是可以在select中运行,但是无法被select到,永远阻塞住的

	//channel里面不管是发数据还是收数据,都是阻塞模式,如果想非阻塞就用select 调度,default获取
	//select 作为channel的调度器,个人感觉和switch差不多
	select {
	case n := <-c1:
		fmt.Println("Received from c1", n)
	case n := <-c2:
		fmt.Println("Received from c2", n)
	default:
		//如果c1和c2都没有数据,则默认走的区间,如果这里没有default,就会报错,我们想从c1和c2发数据但是里面没有,所以就会报错
		fmt.Println("No value received")
	}
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second * 2)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func test2() {
	var c1, c2 = generator(), generator() //声明时直接赋值
	var worker = createWorker(0)
	n := 0
	hasValue := false

	for {
		//nil channel是可以在select中运行的,虽然无法被select到,形成阻塞
		//所以这里我们利用这个特性,做一个判断
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false

			//default:
			//如果这里有default的话,会陷入到default的死循环当中
			//原因就是,上面那两个case的运行还需要一点时间,才发数据,所以就会陷入这个死循环中
			//	fmt.Println("No value received")
		}
	}
}

func test3() {
	var c1, c2 = generator(), generator() //声明时直接赋值
	var worker = createWorker(0)
	//因为供需关系的原因,c1,c2提供的数据比较快,而worker打印的比较慢,会导致中间漏掉很多多数据
	//所以这里我们必须缓存起来
	var values [] int
	//计时器,在指定时间后发送一个数据,返回一个channel,下面我们用case来取,命中之后推出程序
	tm := time.After(time.Second * 10)
	//计时器,每秒发一次信息,返回一个channel,下面我们用case来取,命中之后提示当前队列有多长
	tick := time.Tick(time.Second * 1)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			//每次循环的时候判断, 八百毫秒取不到数据则提示
			fmt.Println("time out")
		case <-tick:
			//每秒统计一下,当前队列有多长
			fmt.Println("queue len=", len(values))
		case <-tm:
			//限制程序只运行十秒钟
			fmt.Println("bye")
			return
		}
	}
}
func main() {
	//test1()
	//test2()
	test3()

}

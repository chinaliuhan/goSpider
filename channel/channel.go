package main

import (
	"fmt"
	"time"
)

//channel
//我们可以开很多goroutine,那么goroutine和goroutine之间的双向的通道就是channel

func chanDemo() {
	//定义channel
	//chan 代表channel 然后是int类型
	//只是定义了c是一个channel 变量类型是一个int,但是他里面的channel并没有帮我们做出来,此时的c 的值是nil,nil的channel我们没办法用
	//nil channel 以后我们学select的时候会用到
	//var c chan int

	//做一个channel出来, 这时候这个channel是可以用的
	//我们创建完channel之后就可以往这里面发数据
	//函数是一等公民,可以作为参数也可以作为返回值,我们的channel也是一等公民
	c := make(chan int)
	//channel的数据必须在goroutine中接收,否则,在channel第一次写入数据时会报错,死锁
	go func() {
		for {
			n := <-c
			//这里打印的时候只能打印出1,打印不出2
			//因为,如果main中没有加sleep的时候,该打印2的时候main就退出了
			//所以,我们在main中加了一个sleep,以后的话会有怎么进行协作,到时候就不用sleep了
			fmt.Println(n)
		}
	}()
	//todo 这时候运行的时候在这里是会报错的,提示发送1时死锁 fatal error: all goroutines are asleep - deadlock!
	//死锁报错, 是因为channel是goroutine和goroutine之间的一个交互,我们必须采用另外一个goroutine去接收
	//也就是说发了一个数据没有任何接收的话,是会出现deadlock的
	//因此这上面我们写一个goroutine的协程来做
	c <- 1 //把1发进去
	c <- 2 //把2发进去
	//从channel中收数据
	//直接接收数据打印是会报错的,理由就是上面我们写的TODO中
	//n := <-c //把channel中的数据读出来赋值到n当中
	//打印我们读出来的数据
	//fmt.Println(n)

	time.Sleep(time.Millisecond)
}

//channel作为参数,也可以添加别的参数
func worker(id int, c chan int) {
	for {
		n := <-c
		//这里的n也可以直接写<-c,上面不用 n:=<-c,不过我依然保留了, 便于以后自己理解
		//这里打印的时候可以发现数据的结果是乱序的
		//这是因为,虽然worker接到的时按照顺序接受的,但是在下面打印的时候,是一个IO操作,goroutine会进行调度,所以会乱序,但是都会打出来
		fmt.Printf("worker %d received %c \n", id, n)
	}
}

func chanDemo2() {
	//这时候channel再次变成了一个一等公民
	//为每个人建立一个channel,因为这类定义一个channel的数组
	//数组里每一个人都是一个channel,然后我们在for循环中把这10个channel分发给这10个worker
	//让后我们再次使用for循环给这10个人发送一些数据
	var channels [10]chan int
	//开多个worker
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		//把数组中定义的10个channel分发给10个worker
		go worker(i, channels[i])
	}
	//然后我们给这10个人分发一些数据
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	//如果觉得不够可以继续打
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

//这里这个chan如果在后面有<-代表这个channel只能用来send数据,如果<-在chan前面,表示只能从channel中取数据,不能往里面存数据
//收到这个channel的人,只能给他们发数据,既然外面的人只能发数据,那我们里面的那个人只能用来<-c收数据
func createWorker(id int) chan<- int {
	//把channel作为返回值,该函数建了一个channel,开了一个goroutine,立刻就返回了,真正做事情是在这个goroutine里面
	//这里自己建立一个channel
	c := make(chan int)
	go func() {
		//这里这一块要分发给一个goroutine去做,不然这里就会死循环去收
		//收的时候,这里还没有人拿到这个c,没人给我发数据,这里就会死掉了
		for {
			fmt.Printf("Worker %d received %c\n", id, <-c)
		}
	}()
	//自己建立玩channel,要把他return出去
	return c
}

func chanDemo3() {
	//如果下面的createWorker的返回channel后面有<-,这里也必须有一个<-,如果前面有一个,这里的前面也要有意给
	//但是如果加了<-之后,那这里的channel里拿出数据来就不可以了,只能往里面写
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		//这里调用createWorker建立10个worker,每个worker建立完之后就会返回一个channel
		//把返回的channel存起来在,事先声明的数组中
		//存起来之后就可以给他们分发数据
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

}

//channel作为参数,也可以添加别的参数
func worker2(id int, c chan int) {
	//如果一旦外面的channel调用了close()关闭了channel之后,这里接收的那个人,还是会在1毫秒之内收到数据的不断的打印,但是,他收到的就是他的具体类型的默认值 int 0 string ''
	//如果这里使用range(遍历完就会退出)就不需要下面的if判断了
	//如果这里不加判断,close之后还在传具体类型的默认值,所以这里永远也不会退出,但是因为外面的main只能运行1毫秒,一毫秒后main函数退出,这里自然也就不存在了
	//只有发送方才可以close
	for n := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}

		//这里的n也可以直接写<-c,上面不用 n:=<-c,不过我依然保留了, 便于以后自己理解
		//这里打印的时候可以发现数据的结果是乱序的
		//这是因为,虽然worker接到的时按照顺序接受的,但是在下面打印的时候,是一个IO操作,goroutine会进行调度,所以会乱序,但是都会打出来
		fmt.Printf("worker %d received %d \n", id, n)
	}
}

//channel缓冲区
func bufferedChannel() {
	//之前我们说过,这里make完之后,下面我们往这个channel中发送数据,发完之后这个程序就会死掉,因为他没有人来收
	//我们一旦发送数据,就必须要有人来收数据
	//但是我们一旦发送了数据就要用协程来接收的话,也是比较耗费资源的,虽然协程是轻量级的
	//这时候我们可以 加入一个缓冲区,比如我们缓冲区的大小是3,一旦设置缓冲区,数据大小不可大于缓冲区,否则会报错,但是如果有人接收的话,超过也没问题
	//跟缓冲区的话,对性能的提升是有一定的优势的
	//c := make(chan int)
	c := make(chan int, 3)
	go worker2(0, c)
	c <- 1
	c <- 2
	c <- 3
	//这里写入超过缓冲区,会出现deadlock的死锁
	c <- 4
	time.Sleep(time.Millisecond)
}

//关闭channel
func channelClose() {
	//只有发送方才可以关闭channel
	c := make(chan int)
	//调用函数处理channel
	go worker2(0, c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	//这时候我们close了这个channel,channel的关闭只能由发送方来做
	//告诉接收方,我们发完了
	//这时候在worker2中打印的内容就不像bufferedChannel的4个也不想其他的有多个,而是在打完1,2,3,4之后还在不断的打印,如果是int 则打印0,string则打印空
	//也就是说一旦外面的channel关闭了之后,里面的接收的那个人,还是会收到数据的,但是一旦外面的关闭了之后,他收到的就是他的具体类型的默认值 int 0 string ''
	//会在这1毫秒之内依然不断的,继续打印,如果想要屏蔽这一块,可以在worker2中打印的时候做判断,如果没有取到值就break即可
	//如果里面不加判断,close之后还在传具体类型的默认值,所以这里永远也不会退出,但是因为外面的main只能运行1毫秒,一毫秒后main函数退出,这里自然也就不存在了
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("channel as first-class citizen 作为一等公民,可以作为参数传来传去")
	//chanDemo()
	//chanDemo2()
	//chanDemo3()
	fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()

	/**
	channel 为什么要做成这个样子
	理论基础: Communication Sequence Process (CSP)的理论
	go语言的并发就是基于这个CSP模型做出来的
	学完了go语言的channel接下来就进行应用
	实践应用有一句话,是go语言的创作者他所说的一句话
	Don't communicate by sharing memory,share memory by communication;
	意思就是 不要通过共享内存来通信,要通过通信来共享内存
	 */

}

package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	//如果相对一块代码区进入一个Lock,而出代码区则释放这个lock,可以向下面这样,使用匿名函数即可
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	//传统同步模式的互斥量的使用,达到线程安全机制
	//这里我们做了一个atomicInt来实现该功能,实际上用的时候不要这么用
	//GO语言给我们的有一个包atomic.AddInt32/64/Uint32/64之类的

	var a atomicInt

	//因为这里increment是对数据的++操作,对内存的写入功能,下面的fmt.Println(a.get())同时是在读取内存中的数据,
	//我们使用go run -race ./atomic 可以看到输出的有数据冲突,就是提示的increment和fmt.Println(a.get()),冲突,一个人在写的时候,我正好在读
	//所以我们在上面使用sync.Mutex对数据的读写加了锁,来对数据的读写进行保护, 达到线程安全的目的
	//加完锁之后我们在使用go run -race ./atomic就不会看到有数据冲突的提示了

	//go语言的传统的同步机制一般我们很少使用,一般我们尽量使用channel来进行通信,因为传统的同步机制他们都是用内存共享进行通信,所以需要用Mutex和Cond来保护

	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())

}

package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int //worker数量
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

//引擎启动
func (e *ConcurrentEngine) Run(seeds ...Request) {
	//所有的worker共用一个输入输出
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in) //把输入送进去,里面的要用来做判断
	for i := 0; i < e.WorkerCount; i++ {
		//创建worker
		createWorker(in, out)
	}

	//遍历得到的URL链接
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itermCount := 0
	//获取out,因为out是worker输出的一个结果
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got tem #%d: %v", itermCount, item)
			itermCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			//发送者角度：对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的。如果chan中的数据无人接收，就无法再给通道传入其他数据。因为新的输入无法在通道非空的情况下传入。所以发送操作会等待 chan 再次变为可用状态：就是通道值被接收时（可以传入变量）。
			//接收者角度：对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
			//因为所有的worker都在忙,没人接收这里的out,会在外面的e.Scheduler.Submit()操作时卡死,所以在哪里加了一个gorou
			out <- result
		}
	}()
}

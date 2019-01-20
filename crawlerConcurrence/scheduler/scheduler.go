package scheduler

import (
	"learnGo/crawlerConcurrence/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//send request down to worker chan
	//s.workerChan <- request

	//发送者角度：对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的。如果chan中的数据无人接收，就无法再给通道传入其他数据。因为新的输入无法在通道非空的情况下传入。所以发送操作会等待 chan 再次变为可用状态：就是通道值被接收时（可以传入变量）。
	//接收者角度：对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。
	//因为所有的worker都在忙,没人接收外面的out,会在这里死掉,所以这里再加有一个goroutine,让他很快的过掉这一行进行下一个操作
	go func() {
		s.workerChan <- request
	}()
}

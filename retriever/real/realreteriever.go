package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

//Retriever这里实现的接口和main.go中的无关,只要实现了GET方法方法就可以了
func (r *Retriever) Get(url string) string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(response, true)
	response.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}

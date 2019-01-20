package engine

import (
	"learnGo/crawler/fetcher"
	"log"
)

//运行爬虫
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//如果队列中确实有数据,则进行遍历
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching url %s", r.Url)
		//调用fetch下载页面
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s:-%v", r.Url, err)
			continue
		}
		//将下载好的页面body,调用传入的函数parser.ParseCityList处理页面URL
		parseResult := r.ParserFunc(body)
		//这里的...代表把parseResult.Requests里面的内容展开一个个加进去,和下面这种方式是一样的结果
		//append(requests, parseResult.Requests[0], parseResult.Requests[1])
		requests = append(requests, parseResult.Requests...)

		//逐个打印item内容
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}

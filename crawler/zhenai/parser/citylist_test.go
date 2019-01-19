package parser

import (
	"io/ioutil"
	"testing"
)

//测试
func TestParseCityList(t *testing.T) {
	//测试的时候,可能会有各种原因,到网络上拿的时候没有,所以下面我们从文件中拿
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("./citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	//将得到的页面解析为城市列表
	result := ParseCityList(contents)
	//
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	const resultSize = 470
	//判断获取的url列表是否符合预期
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d:%s; but was %s", i, url, result.Requests[i].Url)
		}
	}
	//判断获取的城市名称列表是否符合预期
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d:%s; but was %s", i, city, result.Items[i].(string))
		}
	}
}

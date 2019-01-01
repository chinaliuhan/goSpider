package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}
func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Println(writer, "no errror")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	//如果把这个函数写在里面显得太多了,所以提到外面去了
	//而下面的errWrapper()函数,只接收一个
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Not Found"},
	{noError, 200, "no error"},
}

//对errWrapper的函数进行测试
func TestErrWrapper(t *testing.T) {
	//把这个挪到外面去,因为下面的测试server也要用
	//tests := []struct {
	//	h       appHandler
	//	code    int
	//	message string
	//}{
	//	//如果把这个函数写在里面显得太多了,所以提到外面去了
	//	//而下面的errWrapper()函数,只接收一个
	//	{errPanic, 500, "Internal Server Error"},
	//	{errUserError, 400, "user error"},
	//	{errNotFound, 404, "Not Found"},
	//	{errNoPermission, 403, "Forbidden"},
	//	{errUnknown, 500, "Not Found"},
	//	{noError, 200, "no error"},
	//}

	for _, tt := range tests {
		//遍历tests,循环把内容填入到errWrapper中去调用
		f := errWrapper(tt.h)
		//发送网络测试请求信息
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com",
			nil,
		)
		////将网路返回信息填入到errWrapper中
		f(response, request)
		//将下面的部分进行单独的封装了,所以这里注册地
		////读取网络请求的返回内容
		//b, _ := ioutil.ReadAll(response.Body) //测试中不管返回的错误信息
		////b返回的是一个byte[]因为需要转成一个string
		//body := strings.Trim(string(b), "\n")
		//if response.Code != tt.code || body != tt.message {
		//	t.Errorf("expect(%d,%s);"+"got(%d,%s)", tt.code, tt.message, response.Code, body)
		//}

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body) //测试中不管返回的错误信息
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect(%d,%s);"+"got(%d,%s)", expectedCode, expectedMsg, response.StatusCode, body)
	}
}

//起一个server来做测试
//这个测试和上面的有什么区别
//上面的只是拿个假的参数去调用,这次是我们起一个服务器,真的去运行
func TestErrWrapperInServer(t *testing.T) {

	for _, tt := range tests {
		f := errWrapper(tt.h)
		//因为NewServer是一个interface可以用HandlerFunc把f转成一个server的interface
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		//下面单独封装了,所以进行注释
		//b, _ := ioutil.ReadAll(response.Body) //测试中不管返回的错误信息
		//body := strings.Trim(string(b), "\n")
		//if response.StatusCode != tt.code || body != tt.message {
		//	t.Errorf("expect(%d,%s);"+"got(%d,%s)", tt.code, tt.message, response.StatusCode, body)
		//}

		verifyResponse(response, tt.code, tt.message, t)
	}

}

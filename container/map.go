package container

import "fmt"

func main() {
	//定义map [sting]代表键的类型 string代表值的类型 注意map是无序的
	m := map[string]string{
		"name":    "ccc",
		"course":  "golang",
		"site":    "mooc",
		"quality": "notbad",
	}
	//定义空的map
	m2 := make(map[string]int) //m2 == empty map 可以和nil运算混用

	//定义空的Map
	var m3 map[string]int //m3 == nil 同上
	fmt.Println(m, m2, m3)

	//遍历map
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	//通过下标取值, 如果下标不存在, 则会返回空字符串或zero value
	fmt.Println("Getting values")
	courseName := m["course"]
	fmt.Println(courseName)

	//判断是否存在,如果该key存在于map中, OK则会返回true, 否则返回false
	courseName1, ok := m["course"]
	fmt.Println(courseName1, ok)

	if cause, ok := m["cause"]; ok {
		fmt.Println(cause)
	} else {
		fmt.Println("key does not exist")
	}

	//删除元素
	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	//删除后再输出, name返回空字符串, ok返回false
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}

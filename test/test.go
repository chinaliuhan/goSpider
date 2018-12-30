package main

import (
	"fmt"
	"os"
)

type Human int

func (h Human) String() string {
	return fmt.Sprintf("%s","String echo")
}

func (h Human) chinese(name string) (age int, sex string) {

	return 0, "nan"
}

func main() {

	if err := os.Chmod("nihao.log", 777); err != nil {
		fmt.Println(err)
	}

	var h Human

	fmt.Println(h.chinese("zemin"))
}

package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is terraplanets@gmail.com
email is aaa@aaa.com
email is ss@sss.com
email is dd@dd.com.cn
`

func main() {

	//这里填入一个正则表达式,返回一个正则表达式的匹配器,和错误信息
	//re, err := regexp.Compile("terraplanets@gmail.com")
	//这里和上面的一样,不一样的地方在于,如果表达式不对,会直接panic
	//re := regexp.MustCompile("terraplanets@gmail.com")
	//这里的如果要匹配'.',不把他当做一个正则表达式字符的话需要用\\.,否则会被认为是golang的转义字符,也可以向下面那样使用``包含字符串
	//中括号的中'.'不用转义之类的操作
	//re := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z0-9.]+")
	//re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`)
	//子匹配
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//输入原字符串,在源字符串中,通过正则表达式获取符合要求的字符串
	//只匹配第一个,返回一个string
	//match := re.FindString(text)
	//匹配所有,返回一个list
	//match := re.FindAllString(text, -1)
	//子匹配,匹配()中的内容,返回一个二维的list,第二维里面第一个是匹配到的整个字符串,第二个是第一个(),第三个是第二给(),以此类推
	match := re.FindAllStringSubmatch(text, -1)
	//fmt.Println(match)
	for _, m := range match {
		fmt.Println(m)
	}
}

package main

import "fmt"

//取出不重复,连续最长字符串的长度,不支持中文
func lengthOfNonRepeatingSubstr(s string) int {
	//字符串最后一次出现的位置
	lastOccurred := make(map[byte]int)
	//当前开始扫描的位置
	start := 0
	//连续不重复字符串的总长度
	maxLength := 0

	//将字符串s 转换为byte类型, 进行遍历
	for i, ch := range []byte(s) {
		//判断当前遍历的, 字符串ch 是否在map中
		// go原因的if是可以放表达式的,所以我们可以把这句话, 放在if中, 当然也可以单独一行
		// lastI, ok := lastOccurred[ch]
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			//如果以存在, 则将
			start = lastI + 1
		}

		//将当前遍历到的位置存储在maxLength中
		//这里为什么要+1 是因为,这里的start就是开始的位置, 我们要把start向后增加一位,所以+1
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		//储存本次遍历的字符串位置
		lastOccurred[ch] = i
	}
	return maxLength
}

//通过pprof的分析结果查看这里非常的占用性能,所以这里我们使用空间来换性能
//因为rune类型的map字符型的,都很费性能, 优化的方式是我们开一个比较大的slice
var lastOccurred = make([]int, 0xffff)
//支持中文
func lengthOfNonRepeatingSubstr1(s string) int {
	//字符串最后一次出现的位置
	//lastOccurred := make(map[rune]int)
	//性能优化添加的
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	//当前开始扫描的位置
	start := 0
	//连续不重复字符串的总长度
	maxLength := 0

	//将字符串s 转换为byte类型, 进行遍历
	for i, ch := range []rune(s) {
		//判断当前遍历的, 字符串ch 是否在map中
		// go原因的if是可以放表达式的,所以我们可以把这句话, 放在if中, 当然也可以单独一行
		// lastI, ok := lastOccurred[ch]
		//if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
		//如果以存在, 则将
		//start = lastI + 1
		//}
		//性能优化后添加
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1

		}

		//将当前遍历到的位置存储在maxLength中
		//这里为什么要+1 是因为,这里的start就是开始的位置, 我们要把start向后增加一位,所以+1
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		//储存本次遍历的字符串位置
		lastOccurred[ch] = i
	}
	return maxLength
}
func main() {

	//很显然这里就仅仅支持英文字母, 最后一个中文就不支持
	fmt.Println(
		lengthOfNonRepeatingSubstr("aaaasssbbsdafasdfasdf"),
		lengthOfNonRepeatingSubstr("sss"),
		lengthOfNonRepeatingSubstr("aaas"),
		lengthOfNonRepeatingSubstr("你好我是"),
		lengthOfNonRepeatingSubstr("一二三三二一"),
	)
	//下面这个将所有的字符串操作, 支持中文
	fmt.Println(
		lengthOfNonRepeatingSubstr1("aaaasssbbsdafasdfasdf"),
		lengthOfNonRepeatingSubstr1("sss"),
		lengthOfNonRepeatingSubstr1("aass"),
		lengthOfNonRepeatingSubstr1("你好我是"),
		lengthOfNonRepeatingSubstr1("一二三三二一"),
	)
}

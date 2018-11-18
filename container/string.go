package container

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "Yes我爱慕课网!"
	fmt.Println(len(s)) //这里输出19
	//输出字符串的utf8编码
	fmt.Printf("%X\n", []byte(s))

	fmt.Println("---")
	//通过遍历输出每个字的, utf8编码, 会发现输出, 英文为1字节, 汉字为3字节,
	//输出的E...的三组数字就是汉字的utf8编码, 这里的[]byte()应该是通过slice的切片,将字符串转换为byte类型得多所有字节内容,这样便于输出utf8的编码
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b) //输出的结果59 65 73 E6 88 91 E7 88 B1 E6 85 95 E8 AF BE E7 BD 91 21
	}
	fmt.Println()

	//通过遍历这个字符串, 输出他的%d下标, 和unicode编码
	for i, ch := range s { //	这里的ch点开看说他是一个int32的其实就是一个rune类型的
		//这里就会发现, 输出的中文的下标是不连续的,对应上面输出的utf8结果可以发现,59 65 73 E6 88 91 E7 88 B1 E6 85 95 E8 AF BE E7 BD 91 21
		//这里的s 被从utf8解码为unicode 之后又放在了这个rune类型里面
		fmt.Printf("(%d %X)", i, ch) //输出结果(0 59)(1 65)(2 73)(3 6211)(6 7231)(9 6155)(12 8BFE)(15 7F51)(18 21)
	}

	fmt.Println()

	//通过utf8库来做字符串操作
	fmt.Println(
		//这时候这里输出的字符串统计个数是对的
		"Rune count: ", utf8.RuneCountInString(s),
	)

	//将字符串s通过slice切片转换为bytes 获取所有的字节
	bytes := []byte(s)
	for len(bytes) > 0 {
		//该函数会返回一个ch字节, 和一个字节的长度size
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		//这里会将所有的字符一个一个的拿下来
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	//utf8.RuneCountInString(s) 获取含有字符串的字符数量
	//len() 获取字符串的字节长度
	//[]byte 获取字符串的字节内容
	//使用range直接遍历字符串, 获取到的下标是不连续的,因为这时候遍历的内容因该是unicode类型的,英文是连续的 ,中文每次下标都会+3
	//其他字符串的操作可以通过string.*的库来做, 不用个自己去写

	//通过将字符串转换为rune类型来做遍历, 这时候输出下标和字符串, 位置和字符串是匹配的
	//从上面的使用range遍历s的下标是不连续的, 因为这时候如果是中文的话,下标每次会加3
	//如果用len()得到的是字节数,不是字符数, 如果要得到含有中文的字符数需要用utf8.RuneCountInString(s)
	//注意, 之类的转换不是说用另一种方法来解释这块内存中的内容,而是说将取出来的东西, 开辟一块新的内存去解析,然后存储,另外开一个rune的数组
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}

	fmt.Println()
}

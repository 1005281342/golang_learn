package strings_operation

import (
	"fmt"
	"strconv"
	"strings"
)

// strings.HasPrefix(s string, prefix string) bool：判断字符串s是否以prefix开头
// 判断一个url是否以"http://"开头，如果不是，则加上"http://"
func HasPrefixTest(s, h string) string {
	if strings.HasPrefix(s, h) {
		return s
	}
	return h + s
}

// strings.HasSuffix(s string, suffix string) bool：判断字符串s是否以suffix结尾
// 判断一个路径是否以“/”结尾，如果不是，则加上/
func HasSuffixTest(s, e string) string {
	if strings.HasSuffix(s, e) {
		return s
	}
	return s + e
}

// strings.Index(s string, str string) int：判断str在s中首次出现的位置，如果没有出现，则返回-1
func IndexTest(s, str string) int {
	return strings.Index(s, str)
}

// strings.LastIndex(s string, str string) int：判断str在s中最后出现的位置，如果没有出现，则返回-1
func LastIndexTest(s, str string) int {
	return strings.LastIndex(s, str)
}

// 写一个函数返回一个字符串在另一个字符串的首次出现和最后出现位置
func StrIndex(s, str string) (int, int) {
	return IndexTest(s, str), LastIndexTest(s, str)
}

// 写一个程序，从终端读取输入，并转成整数，如果转成整数出错，则输出 “can not convert to int”，并返回。否则输出该整数。
func AtoiTest(a string) {
	_, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("can not convert to int")
	}
}

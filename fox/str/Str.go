package str

import (
	"fmt"
	"strings"
)

//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
	length := len(str)
	if length < 1 {
		return str
	}
	if start < 0 || start > length {
		fmt.Println("Substr error: start is wrong")
		return str
		//panic("start is wrong")
	}

	if end < 0 {
		fmt.Println("Substr error: end is wrong")
		return str
		//panic("end is wrong")
	}
	if end > length {
		end = length
	}

	return string(str[start:end])
}

//截取指定的字符串之前的字符,匹配第一个字符 srt 要截取的字符串 separator 指定字符串
func SubBefore(str string, separator string) string {
	if str == "" {
		return ""
	}
	if separator == "" {
		return str
	}
	if !strings.Contains(str, separator) {
		return str
	}
	index := strings.Index(str, separator)
	return string(str[:index])
}

//截取指定的字符串之前的字符,匹配第一个字符 srt 要截取的字符串 separator 指定字符串
func SubAfter(str string, separator string) string {
	if str == "" {
		return ""
	}
	if separator == "" {
		return str
	}
	if !strings.Contains(str, separator) {
		return str
	}
	index := strings.Index(str, separator)
	return string(str[index:])
}

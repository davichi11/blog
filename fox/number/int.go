package number

import (
	"blog/fox/str"
)

//此处只记录 转换方式，并不使用

//onj变成数字
func ObjToInt(i interface{}) int {
	n := 0
	switch i.(type) {
	case int:
		n = i.(int)
	case int32:
		n = int(i.(int32))
	case int64:
		n = int(i.(int64))
	case float32:
		n = int(i.(float32))
	case float64:
		n = int(i.(float64))
	case string:
		n = str.Int(i.(string))
	}
	return n
}

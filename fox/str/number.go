package str

import "strconv"

//字符串转换成  int
func Int(str string) int {
	return ToInt(str, 0)
}

//字符串转换成  int 出错返回默认值
func ToInt(string2 string, defaults int) int {
	i, e := strconv.Atoi(string2)
	if e != nil {
		return defaults
	}
	return i
}

//字符串转换成  int64,默认0
func Int64(str string) int64 {
	return ToInt64(str, 0)
}

//字符串转换成  int64 出错返回默认值
func ToInt64(str string, defaults int64) int64 {
	i, e := strconv.ParseInt(str, 10, 64)
	if e != nil {
		return defaults
	}
	return i
}

//字符串转换成  float64 没有错误返回
func Float64(str string) float64 {
	return ToFloat64(str, 0.00)
}

//字符串转 float64 错误返回默认值
func ToFloat64(str string, defaults float64) float64 {
	i, e := strconv.ParseFloat(str, 64)
	if e != nil {
		return defaults
	}
	return i
}

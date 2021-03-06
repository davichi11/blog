package datetime

import (
	"fmt"
	"time"
)

//go的格式化固定写法
const (
	Y_M                 = "2006-01"
	Y_M_D               = "2006-01-02"
	Y_M_D_2             = "2006年01月02日"
	Y_M_D_H_I_S         = "2006-01-02 15:04:05"
	Y_M_D_H_I_S_CST     = "2006-01-02 15:04:05 +0800 CST" //2016-12-04 15:39:06 +0800 CST
	Y_M_D_H_I_S_RFC3339 = "2006-01-02T15:04:05Z07:00"
	Y_M_D_H_I_S_2       = "2006年01月02日 15:04:05"
	H_I_S               = "15:04:05"
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.DateTime())
	return []byte(stamp), nil
}

//日期时间
func (t DateTime) DateTime() string {
	return t.Format(Y_M_D_H_I_S)
}

//日期
func (t DateTime) Date() string {
	return t.Format(Y_M_D)
}

//时间
func (t DateTime) Time() string {
	return t.Format(H_I_S)
}
func (t DateTime) Format(layout string) string {
	return Format(t, layout)
}

//格式
func Format(str interface{}, layout string) string {
	date := FormatTimeStruct(str, layout)

	if date.IsZero() {
		return ""
	}

	return date.Format(layout)
}

//转换泛型日期 错误返回当前日期
func FormatTimeStruct(str interface{}, layout string) time.Time {
	var date time.Time
	var err error
	//判断变量类型
	switch d := str.(type) {
	case time.Time:
		date = d
	case string:
		//如果是字符串则转换成 标准日期时间格式
		date, err = time.Parse(layout, d)
		if err != nil {
			return time.Now()
		}
	}

	return date
}

//使用本地时区转换泛型日期 错误返回当前日期
func FormatTimeStructLocation(str interface{}, layout string) time.Time {
	var date time.Time
	var err error
	//判断变量类型
	switch d := str.(type) {
	case time.Time:
		date = d
	case string:
		local, _ := time.LoadLocation("Local")
		//如果是字符串则转换成 标准日期时间格式
		date, err = time.ParseInLocation(layout, d, local)
		if err != nil {
			return time.Now()
		}
	}

	return date
}

//当前日期时间
func Now() string {
	return Format(time.Now(), Y_M_D_H_I_S)
}

//当前日期
func Date() string {
	return Format(time.Now(), Y_M_D)
}

//当前时间
func Time() string {
	return Format(time.Now(), H_I_S)
}

//当前年月
func YearMonth() string {
	return Format(time.Now(), Y_M)
}

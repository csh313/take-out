package times

import "time"

// 自定义时间类型
type CustomTime time.Time

// 实现一个自定义的解析方法，解析 "yyyy-MM-dd HH:mm:ss" 格式
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	// 去掉前后双引号
	str = str[1 : len(str)-1]

	// 使用自定义格式进行解析
	parsedTime, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return err
	}

	*ct = CustomTime(parsedTime)
	return nil
}

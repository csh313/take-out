package code

type ErrorCode int

// 错误码
// 通常是从JSON文件中读取
const (
	SettingsError ErrorCode = 1001 //系统错误
	ArgumentError ErrorCode = 1002 //参数错误
)

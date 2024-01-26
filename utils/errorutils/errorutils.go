package errorutils

import "log"

// ExitOnError 打印日志并中断程序
func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// LogError 仅打印错误日志
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

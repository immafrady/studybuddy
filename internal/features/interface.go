package features

import (
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
)

type ExamTaker interface {
	TakeExam() bool      // 理解,回答
	FormatLabel() string // 生成标签
	GetAnswer() string   // 获取答案
}

type Pipeline[T any] interface {
	ParseOption(str string) []promptuihelper.Option[T]                       // 解析问题
	DoTask([]promptuihelper.Option[T], ExamTaker) (ans string, correct bool) // 做题+判断
}

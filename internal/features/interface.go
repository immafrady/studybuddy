package features

import (
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
)

type ExamTaker interface {
	TakeExam() (correct bool) // 理解,回答
	GetLabel() string         // 生成标签
	GetAnswer() string        // 获取答案
	DisplayResult()           // 展示结果
}

type Pipeline[T any] interface {
	ParseOption(str string) []promptuihelper.Option[T]                       // 解析问题
	DoTask([]promptuihelper.Option[T], ExamTaker) (ans string, correct bool) // 做题+判断
}

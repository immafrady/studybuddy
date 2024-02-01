package marx

import (
	"encoding/json"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
)

// selection 下拉选项
type selection struct {
	A string `json:"A,omitempty"`
	B string `json:"B,omitempty"`
	C string `json:"C,omitempty"`
	D string `json:"D,omitempty"`
}

// jsonStringify 格式化为字符串
func (s selection) jsonStringify() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}

// parseSelection 从字符串解析
func parseSelection(str string) *selection {
	var s selection
	errorhelper.LogError(json.Unmarshal([]byte(str), &s))
	return &s
}

// fromRawQuiz 从原始数据中获取值
func fromRawQuiz(quiz rawQuiz) *selection {
	return &selection{
		A: quiz.A,
		B: quiz.B,
		C: quiz.C,
		D: quiz.D,
	}
}

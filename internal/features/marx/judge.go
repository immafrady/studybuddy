package marx

import (
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
)

type judgePipeline struct {
	features.Pipeline[string]
}

func (j judgePipeline) ParseOption(_ string) []promptuihelper.Option[string] {
	return []promptuihelper.Option[string]{
		{Label: "正确", Value: "Y"},
		{Label: "错误", Value: "X"},
	}
}

func (j judgePipeline) DoTask(options []promptuihelper.Option[string], et features.ExamTaker) (string, bool) {
	ans := promptuihelper.SingleChoiceSelect(options, promptuihelper.SelectConfig{Label: et.GetLabel()})
	return ans, ans == et.GetAnswer()
}

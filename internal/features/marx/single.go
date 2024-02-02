package marx

import (
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
)

type singlePipeline struct {
	features.Pipeline[string]
}

func (s singlePipeline) ParseOption(str string) []promptuihelper.Option[string] {
	selection := parseSelection(str)
	return selection.toOptions()
}

func (s singlePipeline) DoTask(options []promptuihelper.Option[string], et features.ExamTaker) (string, bool) {
	ans := promptuihelper.SingleChoiceSelect(options, promptuihelper.SelectConfig{Label: et.GetLabel()})
	return ans, ans == et.GetAnswer()
}

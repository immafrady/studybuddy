package marx

import (
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
	"strings"
)

type multiplePipeline struct {
	features.Pipeline[string]
}

func (m multiplePipeline) ParseOption(str string) []promptuihelper.Option[string] {
	selection := parseSelection(str)
	return selection.toOptions()
}

func (m multiplePipeline) DoTask(options []promptuihelper.Option[string], et features.ExamTaker) (string, bool) {
	ans := promptuihelper.MultipleChoiceSelect(options, promptuihelper.SelectConfig{Label: et.FormatLabel()})
	ansStr := strings.Join(ans, "")
	return ansStr, ansStr == et.GetAnswer()
}

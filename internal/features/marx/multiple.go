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

func (m multiplePipeline) DoTask(options []promptuihelper.Option[string], et features.ExamTaker) bool {
	ans := promptuihelper.MultipleChoiceSelect(options, promptuihelper.SelectConfig{Label: et.FormatLabel()})
	return strings.Join(ans, "") == et.GetAnswer()
}

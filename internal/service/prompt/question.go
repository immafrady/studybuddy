package prompt

import (
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
	"github.com/immafrady/studybuddy/internal/model"
)

func SelectQuestionType() []model.QuestionType {
	options := []promptuihelper.MultipleOption[model.QuestionType]{
		{Label: "单选题", Value: model.QuestionSingle, Checked: true},
		{Label: "多选题", Value: model.QuestionMultiple, Checked: true},
		{Label: "判断题", Value: model.QuestionJudge, Checked: true},
	}
	return promptuihelper.MultipleChoiceSelect(options, promptuihelper.SelectConfig{
		Label: "请选择答题类型",
	})
}

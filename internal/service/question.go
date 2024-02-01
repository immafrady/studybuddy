package service

import (
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
	"github.com/immafrady/studybuddy/internal/model"
)

func SelectQuestionType() []model.QuestionType {
	options := []promptuihelper.MultipleOption[model.QuestionType]{
		{Label: "单选题", Value: model.QuestionSingle},
		{Label: "多选题", Value: model.QuestionMultiple},
		{Label: "判断题", Value: model.QuestionJudge},
	}
	return promptuihelper.MultipleChoiceSelect(options, "请选择答题类型：")
}

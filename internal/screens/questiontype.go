package screens

import (
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
	"github.com/immafrady/studybuddy/internal/model"
)

func QuestionTypeScreen() []model.QuestionType {
	m := selection.Run(selection.Config{
		Options: []*selection.Option{
			{Label: "单选题", Value: model.QuestionSingle, IsChecked: true},
			{Label: "多选题", Value: model.QuestionMultiple, IsChecked: true},
			{Label: "判断题", Value: model.QuestionJudge, IsChecked: true},
		},
		Title:    "请选择答题类型",
		Multiple: true,
	})
	vs := m.GetSelectedValues()
	var types = make([]model.QuestionType, len(vs))
	for i, v := range vs {
		if t, ok := v.(model.QuestionType); ok {
			types[i] = t
		}
	}
	return types
}

package service

import (
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/manifoldco/promptui"
)

type questionOption struct {
	Label   string
	Name    model.QuestionType
	Checked bool
}

func (o *questionOption) onSelect() {
	o.Checked = !o.Checked
}

func SelectQuestionType() []model.QuestionType {
	options := []questionOption{
		{Label: "Confirm"},
		{Label: "单选题", Name: model.QuestionSingle},
		{Label: "多选题", Name: model.QuestionMultiple},
		{Label: "判断题", Name: model.QuestionJudge},
	}
	index := -1
	for index != 0 {
		prompt := promptui.Select{
			Label: "请选择答题类型：",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}：",
				Active:   "● {{ if .Name }}[{{ if .Checked }}x{{ else }} {{ end}}] {{ .Label }}{{ else }}{{ .Label }}{{ end }}",
				Inactive: "○ {{ if .Name }}[{{ if .Checked }}x{{ else }} {{ end}}] {{ .Label }}{{ else }}{{ .Label }}{{ end }}",
			},
			HideSelected: true,
		}
		index, _, _ = prompt.Run()
		if index != 0 && index < len(options) {
			options[index].onSelect()
		}
	}
	var types []model.QuestionType
	for i := 1; i < len(options); i++ {
		if options[i].Checked {
			types = append(types, options[i].Name)
		}
	}
	return types
}

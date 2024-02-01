package service

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/manifoldco/promptui"
)

// SelectClassifyId 选择类别
func SelectClassifyId() uint {
	db, _ := database.Get()
	var classifies []model.Classify
	db.Find(&classifies)
	classifies = append(classifies, model.Classify{Name: "全部"})

	prompt := promptui.Select{
		Label: "请选择题目类别",
		Items: classifies,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}：",
			Active:   "● {{ .Name }}",
			Inactive: "○ {{ .Name }}",
			Selected: "已选择： {{ .Name }}",
		},
		Size: 5,
	}

	i, _, err := prompt.Run()
	errorhelper.LogError(err)
	if i < len(classifies) {
		item := classifies[i]
		return item.ID
	} else {
		return 0
	}
}

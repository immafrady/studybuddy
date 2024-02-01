package service

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/helpers/promptuihelper"
	"github.com/immafrady/studybuddy/internal/model"
)

// SelectSingleClassify 单选类别
func SelectSingleClassify() model.Classify {
	db, _ := database.Get()
	var classifies []model.Classify
	db.Find(&classifies)
	var options = make([]promptuihelper.SingleOption[model.Classify], len(classifies))

	for i, c := range classifies {
		options[i] = promptuihelper.SingleOption[model.Classify]{
			Label: c.Name,
			Value: c,
		}
	}

	return promptuihelper.SingleChoiceSelect(options, "请选择题目类别")
}

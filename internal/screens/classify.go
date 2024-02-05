package screens

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
	"github.com/immafrady/studybuddy/internal/model"
)

func ClassifyRun() (classify model.Classify) {
	db, _ := database.Get()
	var classifies []model.Classify
	db.Find(&classifies)

	var options = make([]*selection.Option, len(classifies))
	for i, c := range classifies {
		options[i] = &selection.Option{
			Label: string(c.Name),
			Value: c,
		}
	}

	m := selection.Run(selection.Config{
		Options: options,
		Title:   "请选择题目类别",
	})

	vs := m.GetSelectedValues()
	if len(vs) == 1 {
		if c, ok := vs[0].(model.Classify); ok {
			classify = c
		}
	}
	return classify
}

package dispatcher

import "github.com/immafrady/studybuddy/internal/model"

type Context struct {
	classify *model.Classify
	record   *model.Record
	types    []model.QuestionType
	limit    int
}

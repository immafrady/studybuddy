package ctx

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
	"github.com/immafrady/studybuddy/internal/model"
	"gorm.io/gorm"
)

type Context struct {
	DB       *gorm.DB
	Classify *model.Classify
	Record   *model.Record
	Types    []model.QuestionType
	Limit    int
	History  []*selection.Model // 记录每一题的过程
}

func NewContext() *Context {
	db, _ := database.Get()
	return &Context{
		DB: db,
	}
}

func (c Context) DoOver(callback func(ctx *Context)) {
	record := &model.Record{ClassifyId: c.Classify.ID}

	callback(&Context{
		DB:       c.DB,
		Classify: c.Classify,
		Record:   record,
		Types:    c.Types,
		Limit:    c.Limit,
		History:  nil,
	})
}

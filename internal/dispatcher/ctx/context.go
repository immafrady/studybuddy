package ctx

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper"
	"github.com/immafrady/studybuddy/internal/model"
	"gorm.io/gorm"
	"os"
)

type Context struct {
	DB       *gorm.DB
	Classify *model.Classify
	Record   *model.Record
	Types    []model.QuestionType
	Limit    int
	History  []tuihelper.ResultOutput // 记录每一题的过程
	State
	DoOverFn func(ctx *Context)
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

func (c Context) HandleState() {
	switch c.State {
	case DoNothing:
		os.Exit(1)
	case DoAgain:
		c.DoOver(c.DoOverFn)
	case StartAllOver:
		// 不需要操作什么自动重来
	}
}

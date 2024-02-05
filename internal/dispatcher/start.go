package dispatcher

import (
	"github.com/immafrady/studybuddy/internal/features/examiner"
	"github.com/immafrady/studybuddy/internal/screens"
	"github.com/immafrady/studybuddy/internal/service/quiz"
	"gorm.io/gorm"
)

func start(ctx *Context) {
	if ctx.classify == nil {
		classify := screens.ClassifyRun()
		ctx.classify = &classify
	}
	if ctx.types == nil {
		ctx.types = screens.QuestionTypeScreen()
	}
	if ctx.limit == 0 {
		ctx.limit = screens.LimitScreen()
	}
	quiz.FetchQuestionList(ctx.classify, func(db *gorm.DB) *gorm.DB {
		return db.Where("type IN (?)", ctx.types).Order("count").Limit(ctx.limit)
	})
	exam := examiner.NewExamHolder(ctx.classify, ctx.types, ctx.limit)
	exam.Start()

	// todo 到时候判断一个布尔值，如果是true就调用
	doOver(ctx, start)
}

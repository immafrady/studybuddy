package controller

import (
	"github.com/immafrady/studybuddy/internal/dispatcher/ctx"
	"github.com/immafrady/studybuddy/internal/features/examiner"
	"github.com/immafrady/studybuddy/internal/screens"
)

func Start(ctx *ctx.Context) {
	if ctx.Classify == nil {
		classify := screens.ClassifyRun()
		ctx.Classify = &classify
	}
	if ctx.Types == nil {
		ctx.Types = screens.QuestionTypeScreen()
	}
	if ctx.Limit == 0 {
		ctx.Limit = screens.LimitScreen()
	}
	examHolder := examiner.Examiner{
		Ctx:        ctx,
		ShowResult: true,
	}
	examHolder.StartExam()
	// todo 到时候判断一个布尔值，如果是true就调用
	ctx.DoOver(Start)
}

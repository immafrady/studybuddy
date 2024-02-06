package controller

import (
	"github.com/immafrady/studybuddy/internal/dispatcher/ctx"
	"github.com/immafrady/studybuddy/internal/features/examiner"
	"github.com/immafrady/studybuddy/internal/screens"
)

func Start(c *ctx.Context) {
	c.DoOverFn = Start
	if c.Classify == nil {
		classify := screens.ClassifyRun()
		c.Classify = &classify
	}
	if c.Types == nil {
		c.Types = screens.QuestionTypeScreen()
	}
	if c.Limit == 0 {
		c.Limit = screens.LimitScreen()
	}
	examHolder := examiner.Examiner{
		Ctx:        c,
		ShowResult: true,
	}
	examHolder.StartExam()
	c.HandleState()
}

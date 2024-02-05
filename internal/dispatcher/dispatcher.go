package dispatcher

import (
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/immafrady/studybuddy/internal/screens"
)

func Dispatch() {
	token := screens.HomeRun()
	switch token {
	case screens.TokenStart:
		start(&Context{})
	case screens.TokenExam:
	case screens.TokenReview:
	case screens.TokenHistory:
	}
}

func doOver(ctx *Context, callback func(ctx *Context)) {
	ctx.record = &model.Record{ClassifyId: ctx.classify.ID}
	callback(ctx)
}

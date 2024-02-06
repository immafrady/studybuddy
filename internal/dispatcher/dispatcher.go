package dispatcher

import (
	"github.com/immafrady/studybuddy/internal/dispatcher/controller"
	"github.com/immafrady/studybuddy/internal/dispatcher/ctx"
	"github.com/immafrady/studybuddy/internal/screens"
)

func Dispatch() {
	token := screens.HomeRun()
	switch token {
	case screens.TokenStart:
		controller.Start(ctx.NewContext())
	case screens.TokenExam:
	case screens.TokenReview:
	case screens.TokenHistory:
	}
}

package dispatcher

import "github.com/immafrady/studybuddy/internal/screens"

func Dispatch() {
	token := screens.HomeRun()
	switch token {
	case screens.TokenStart:
		start()
	case screens.TokenExam:
	case screens.TokenReview:
	case screens.TokenHistory:
	}
}

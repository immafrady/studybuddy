package features

import (
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
)

type ExamTaker2 interface {
	DoExam(showResult bool) selection.Model
}

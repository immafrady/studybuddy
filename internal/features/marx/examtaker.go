package marx

import (
	"fmt"
	"github.com/immafrady/studybuddy/internal/features"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/immafrady/studybuddy/internal/screens"
	"strings"
)

type ExamTaker struct {
	features.ExamTaker2
	*model.Question
	index  int
	total  int
	screen screens.QuizScreen
}

func NewExamTaker2(question *model.Question, curr int, total int) features.ExamTaker2 {
	return ExamTaker{
		Question: question,
		index:    curr,
		total:    total,
	}
}

func (e ExamTaker) DoExam(showResult bool) selection.Model {
	title := fmt.Sprintf("(%v/%v:%v) %v ", e.index+1, e.total, model.QuestionTypeLabelMap[e.Type], e.Q)
	var options []*selection.Option
	if e.Type == model.QuestionJudge {
		options = []*selection.Option{
			{Label: "正确", Value: "Y", IsCorrect: e.A == "Y"},
			{Label: "错误", Value: "X", IsCorrect: e.A == "X"},
		}
	} else {
		s := parseSelection(e.Detail)
		options = []*selection.Option{
			{Label: "A." + s.A, Value: "A", IsCorrect: strings.Contains(e.A, "A")},
			{Label: "B." + s.B, Value: "B", IsCorrect: strings.Contains(e.A, "B")},
			{Label: "C." + s.C, Value: "C", IsCorrect: strings.Contains(e.A, "C")},
			{Label: "D." + s.D, Value: "D", IsCorrect: strings.Contains(e.A, "D")},
		}
	}
	return screens.QuizSelectionRun(screens.QuizSelectionRunArgs{
		Question:   e.Question,
		Options:    options,
		Title:      title,
		ShowResult: showResult,
		Multiple:   e.Type == model.QuestionMultiple,
	})
}

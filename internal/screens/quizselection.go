package screens

import (
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
	"github.com/immafrady/studybuddy/internal/model"
	"github.com/immafrady/studybuddy/internal/service/quiz"
)

type QuizSelectionRunArgs struct {
	Screen
	QuizScreen
	Question   *model.Question // 这里传指针，可以直接改变外面的值
	Options    []*selection.Option
	Title      string
	ShowResult bool
	Multiple   bool
}

// QuizSelectionRun 下拉选项的抽象页面；可以控制showResult
func QuizSelectionRun(args QuizSelectionRunArgs) selection.Model {
	return selection.Run(selection.Config{
		Options:  args.Options,
		Title:    args.Title,
		Multiple: args.Multiple,
		Liked:    args.Question.Like,
		ToggleLikeFn: func() bool {
			return quiz.ToggleQuestionLike(args.Question)
		},
		Index:      0,
		ShowResult: args.ShowResult,
	})
}

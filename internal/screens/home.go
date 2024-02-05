package screens

import (
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
)

type HomeActionToken string

const (
	TokenStart   HomeActionToken = "Start"
	TokenExam    HomeActionToken = "Exam"
	TokenReview  HomeActionToken = "Review"
	TokenHistory HomeActionToken = "History"
)

func HomeRun() (token HomeActionToken) {

	m := selection.Run(selection.Config{
		Options: []*selection.Option{
			{Label: "开始做题", Desc: "随便刷刷，从做的少的题目开始", Value: TokenStart},
			{Label: "考试模式", Desc: "来一场痛痛快快都决斗吧！", Value: TokenExam},
			{Label: "复习模式", Desc: "查漏补缺", Value: TokenReview},
			{Label: "历史记录", Desc: "看看曾经都做过些什么", Value: TokenHistory},
		},
		Title: "请选择模式",
	})
	vs := m.GetSelectedValues()
	if len(vs) == 1 {
		if t, ok := vs[0].(HomeActionToken); ok {
			token = t
		}
	}
	return
}

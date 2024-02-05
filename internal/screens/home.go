package screens

import (
	"github.com/immafrady/studybuddy/internal/dispatcher"
	"github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"
)

type wrapper struct {
	cb func()
}

func HomeRun() {

	m := selection.Run(selection.Config{
		Options: []*selection.Option{
			{Label: "开始做题", Desc: "随便刷刷，从做的少的题目开始", Value: wrapper{
				func() {
					dispatcher.Start()
				},
			}},
			{Label: "考试模式", Desc: "来一场痛痛快快都决斗吧！", Value: wrapper{
				func() {
					dispatcher.Start()
				},
			}},
			{Label: "复习模式", Desc: "查漏补缺", Value: wrapper{
				func() {
					dispatcher.Start()
				},
			}},
			{Label: "历史记录", Desc: "看看曾经都做过些什么", Value: wrapper{
				func() {
					dispatcher.Start()
				},
			}},
		},
		Title: "请选择模式",
	})
	vs := m.GetSelectedValues()
	if len(vs) == 1 {
		if w, ok := vs[0].(wrapper); ok {
			w.cb()
		}
	}
}

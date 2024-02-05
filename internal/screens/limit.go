package screens

import "github.com/immafrady/studybuddy/internal/helpers/tuihelper/selection"

func LimitScreen() int {
	m := selection.Run(selection.Config{
		Options: []*selection.Option{
			{Label: "10", Desc: "随便玩玩", Value: 10},
			{Label: "25", Desc: "上点强度", Value: 25},
			{Label: "50", Desc: "我认真了", Value: 50},
			{Label: "100", Desc: "挑战自己", Value: 100},
		},
		Title: "请选择做题数量",
	})
	vs := m.GetSelectedValues()
	if len(vs) == 1 {
		if i, ok := vs[0].(int); ok {
			return i
		}
	}
	return 0
}

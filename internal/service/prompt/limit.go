package prompt

import "github.com/immafrady/studybuddy/internal/helpers/promptuihelper"

func SelectLimit() int {
	options := []promptuihelper.Option[int]{
		{Label: "10  (随便玩玩)", Value: 10},
		{Label: "25  (上点强度)", Value: 25},
		{Label: "50  (我认真了)", Value: 50},
		{Label: "100 (挑战自己)", Value: 100},
	}
	return promptuihelper.SingleChoiceSelect(options, promptuihelper.SelectConfig{
		Label: "请选择做题数量",
	})
}

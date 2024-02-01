package controller

import "fmt"

var examEntry = entryItem{
	Label:  "考试模式",
	Detail: "来一场痛痛快快都决斗吧！",
	Callback: func() {
		fmt.Println("待实现")
	},
}

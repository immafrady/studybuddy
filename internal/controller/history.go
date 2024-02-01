package controller

import "fmt"

var historyEntry = entryItem{
	Label:  "历史记录",
	Detail: "看看曾经都做过些什么",
	Callback: func() {
		fmt.Println("待实现")
	},
}

package controller

import "fmt"

var reviewEntry = entryItem{
	Label:  "复习模式",
	Detail: "查漏补缺",
	Callback: func() {
		fmt.Println("待实现")
	},
}

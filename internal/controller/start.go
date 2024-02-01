package controller

import (
	"fmt"
	"github.com/immafrady/studybuddy/internal/service/prompt"
)

var startEntry = entryItem{
	Label:  "开始做题",
	Detail: "随便刷刷，从做的少的题目开始",
	Callback: func() {
		classify := prompt.SelectSingleClassify()
		fmt.Printf("您选择的类目是：%v\n", classify.Name)
		types := prompt.SelectQuestionType()
		fmt.Println(types)
	},
}

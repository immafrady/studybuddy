package controller

import (
	"fmt"
	"github.com/immafrady/studybuddy/internal/features/examiner"
	"github.com/immafrady/studybuddy/internal/service/prompt"
	"github.com/immafrady/studybuddy/internal/service/quiz"
	"gorm.io/gorm"
)

var startEntry = entryItem{
	Label:  "开始做题",
	Detail: "随便刷刷，从做的少的题目开始",
	Callback: func() {
		classify := prompt.SelectSingleClassify()
		fmt.Printf("您选择的类目是：%v\n", classify.Name)
		types := prompt.SelectQuestionType("下一步")
		limit := prompt.SelectLimit()
		quiz.FetchQuestionList(&classify, func(db *gorm.DB) *gorm.DB {
			return db.Where("type IN (?)", types).Limit(limit)
		})
		exam := examiner.NewExamHolder(&classify, types, limit)
		exam.Start()
	},
}

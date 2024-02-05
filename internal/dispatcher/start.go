package dispatcher

import (
	"fmt"
	"github.com/immafrady/studybuddy/internal/features/examiner"
	"github.com/immafrady/studybuddy/internal/screens"
	"github.com/immafrady/studybuddy/internal/service/quiz"
	"gorm.io/gorm"
)

func start() {
	classify := screens.ClassifyRun()
	fmt.Printf("您选择的类目是：%v\n", classify.Name)
	types := screens.QuestionTypeScreen()
	limit := screens.LimitScreen()
	quiz.FetchQuestionList(&classify, func(db *gorm.DB) *gorm.DB {
		return db.Where("type IN (?)", types).Order("count").Limit(limit)
	})
	exam := examiner.NewExamHolder(&classify, types, limit)
	exam.Start()
}

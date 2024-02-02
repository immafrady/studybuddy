package quiz

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/model"
)

// AddAnswerRecord 新增一条记录
func AddAnswerRecord(record *model.Record, id uint, correct bool) {
	db, _ := database.Get()
	if !correct {
		record.WrongQuestionIds = append(record.WrongQuestionIds, id)
	}
	record.QuestionIds = append(record.QuestionIds, id)
	db.Save(record)
}

package quiz

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/model"
	"gorm.io/gorm"
)

// FetchQuestionList 获取问题列表
func FetchQuestionList(classify *model.Classify, questionPreload func(db *gorm.DB) *gorm.DB) {
	db, _ := database.Get()
	db.Preload("Questions", func(db *gorm.DB) *gorm.DB {
		return questionPreload(db).Order("random()")
	}).Find(classify)
}

// MarkQuestionDone 标记题目做过
func MarkQuestionDone(question *model.Question, correct bool) {
	db, _ := database.Get()
	if !correct {
		question.WrongCount += 1
	}
	question.Count += 1
	db.Save(question)
}

// ToggleQuestionLike 标记/取消标记题目喜欢
func ToggleQuestionLike(question *model.Question) {
	db, _ := database.Get()
	question.Like = !question.Like
	db.Save(question)
}

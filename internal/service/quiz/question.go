package quiz

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/model"
	"gorm.io/gorm"
)

// FetchQuestionList 获取问题列表
func FetchQuestionList(classify *model.Classify, questionPreload func(db *gorm.DB) *gorm.DB) {
	db, _ := database.Get()
	db.Debug().Preload("Questions", func(db *gorm.DB) *gorm.DB {
		return questionPreload(db).Order("random()")
	}).Find(classify)
}

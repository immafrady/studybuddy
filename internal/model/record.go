package model

import "gorm.io/gorm"

// Record 答题记录
type Record struct {
	gorm.Model
	ClassifyId       uint   // 分类ID
	QuestionIds      []uint `gorm:"serializer:json"` // 关联问题
	WrongQuestionIds []uint `gorm:"serializer:json"` // 答错的问题
}

func (r Record) TableName() string {
	return "quiz_record"
}

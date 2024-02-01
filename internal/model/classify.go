package model

import "gorm.io/gorm"

// Classify 课题分类
type Classify struct {
	gorm.Model
	Name     string `gorm:"column:name"` // 问题分类
	Question []Question
}

func (c Classify) TableName() string {
	return "quiz_classify"
}

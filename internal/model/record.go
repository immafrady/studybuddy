package model

import "gorm.io/gorm"

// Record 答题记录
type Record struct {
	gorm.Model
	Classify       Classify   // 分类ID
	Questions      []Question // 关联问题
	WrongQuestions []Question // 答错的问题
}

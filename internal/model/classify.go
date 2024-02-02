package model

import "gorm.io/gorm"

type ClassName string

const (
	ClassMarx ClassName = "马克思主义基本原理"
)

// Classify 课题分类
type Classify struct {
	gorm.Model
	Name      ClassName // 问题分类
	Questions []Question
}

func (c Classify) TableName() string {
	return "quiz_classify"
}

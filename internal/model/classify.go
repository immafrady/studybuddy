package model

import "gorm.io/gorm"

// Classify 课题分类
type Classify struct {
	gorm.Model
	Name string // 问题分类
}

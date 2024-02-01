package model

import (
	"gorm.io/gorm"
)

type QuestionType string

const (
	QuestionMultiple QuestionType = "Multi"  // 多选
	QuestionSingle   QuestionType = "Single" // 单选
	QuestionJudge    QuestionType = "Judge"  // 判断
)

// Question 问题
type Question struct {
	gorm.Model
	ClassifyId uint         `gorm:"foreignKey:id"` // 关联的课题分类
	Q          string       `gorm:"column:q"`      // 问题
	A          string       `gorm:"column:a"`      // 答案
	Detail     string       `gorm:"column:detail"` // 问题补充
	Like       bool         `gorm:"column:like"`   // 是否标记（todo 不知道bool类型支不支持）
	Type       QuestionType `gorm:"column:type"`   // 问题类型
}

func (q Question) TableName() string {
	return "quiz_question"
}

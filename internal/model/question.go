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

var QuestionTypeLabelMap = map[QuestionType]string{
	QuestionMultiple: "多选题",
	QuestionSingle:   "单选题",
	QuestionJudge:    "判断题",
}

// Question 问题
type Question struct {
	gorm.Model
	ClassifyId uint         `gorm:"foreignKey:id"` // 关联的课题分类
	Q          string       // 问题
	A          string       // 答案
	Detail     string       // 问题补充
	Like       bool         // 是否标记
	Count      uint         // 做的次数
	WrongCount uint         // 错的次数
	Type       QuestionType // 问题类型
	UserAnswer string       `gorm:"-"` // 用户的回答，仅作记录用
}

func (q *Question) TableName() string {
	return "quiz_question"
}

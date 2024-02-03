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
	WrongRate  uint         `gorm:"-"`           // 计算一下失败率 todo 好像可以删
	Type       QuestionType `gorm:"column:type"` // 问题类型
}

func (q *Question) AfterFind(db *gorm.DB) error {
	if q.Count != 0 {
		q.WrongRate = uint(float64(q.WrongCount) / float64(q.Count) * 100)
	}
	return nil
}

func (q *Question) TableName() string {
	return "quiz_question"
}

package model

import (
	"gorm.io/gorm"
)

// Question 问题
type Question struct {
	gorm.Model
	Classify Classify // 关联的课题分类
	Q        string   // 问题
	A        string   // 答案
	Detail   string   // 问题补充
	Like     bool     // 是否标记（todo 不知道bool类型支不支持）
}

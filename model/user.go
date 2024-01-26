package model

import (
	"github.com/immafrady/studybuddy/utils/errorutils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func init() {
	// 生产是不是得去掉?
	errorutils.ExitOnError(db.AutoMigrate(&User{}))
}

func (u User) TableName() string {
	return "sys_users"
}

func (u User) Create() {
	db.Create(&u)
}

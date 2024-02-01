// Package startup 项目启动
package startup

import (
	"github.com/immafrady/studybuddy/internal/database"
	"github.com/immafrady/studybuddy/internal/features/marx"
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"github.com/immafrady/studybuddy/internal/model"
)

func Bootstrap() {
	db, initialized := database.Get()
	errorhelper.ExitOnError(db.AutoMigrate(&model.Classify{}, &model.Question{}, &model.Record{}))
	if !initialized {
		// 首次调用
		marx.LoadData()
	}
}

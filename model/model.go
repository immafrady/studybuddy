package model

import (
	"github.com/immafrady/studybuddy/utils/errorutils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
)

var db *gorm.DB

func init() {
	baseDir, _ := os.UserConfigDir()
	folderDir := path.Join(baseDir, "studybuddy")
	_, err := os.Stat(folderDir)
	if os.IsNotExist(err) {
		errorutils.ExitOnError(os.MkdirAll(folderDir, os.ModePerm))
	}
	db, err = gorm.Open(sqlite.Open(path.Join(folderDir, "studybuddy.sqlite3")), &gorm.Config{})
	errorutils.ExitOnError(err)
}

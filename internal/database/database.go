package database

import (
	"github.com/immafrady/studybuddy/internal/helpers/errorhelper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
)

var db *gorm.DB
var baseDir string
var folderDir string
var dbPath string

func init() {
	baseDir, _ = os.UserConfigDir()
	folderDir = path.Join(baseDir, "studybuddy")
	dbPath = path.Join(folderDir, "studybuddy.sqlite3")
}

// Get 连接最新的数据库连接实例
func Get() (*gorm.DB, bool) {
	initialized := true
	if db == nil {
		connect()
		initialized = false
	}
	return db, initialized
}

// Reset 重置并返回最新的数据库连接实例
func Reset() (*gorm.DB, bool) {
	if !isFolderDirNotExist() {
		errorhelper.ExitOnError(os.Remove(dbPath))
	}
	return Get()
}

// connect 链接数据库
func connect() {
	if isFolderDirNotExist() {
		errorhelper.ExitOnError(os.MkdirAll(folderDir, os.ModePerm))
	}
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	errorhelper.ExitOnError(err)
}

// isFolderDirNotExist 判断数据文件夹是否存在
func isFolderDirNotExist() bool {
	_, err := os.Stat(folderDir)
	return os.IsNotExist(err)
}

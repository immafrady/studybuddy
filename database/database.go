package database

import (
	"github.com/immafrady/studybuddy/utils/errorutils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path"
	"time"
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

// Get 获取最新的数据库连接实例
func Get() *gorm.DB {
	if db == nil {
		connect()
	}
	return db
}

// Reset 重置并返回最新的数据库连接实例
func Reset() *gorm.DB {
	if !isFolderDirNotExist() {
		errorutils.ExitOnError(os.Remove(dbPath))
	}
	time.Sleep(5 * time.Second)
	return Get()
}

// connect 链接数据库
func connect() {
	if isFolderDirNotExist() {
		errorutils.ExitOnError(os.MkdirAll(folderDir, os.ModePerm))
	}
	var err error
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	errorutils.ExitOnError(err)
}

// isFolderDirNotExist 判断数据文件夹是否存在
func isFolderDirNotExist() bool {
	_, err := os.Stat(folderDir)
	return os.IsNotExist(err)
}

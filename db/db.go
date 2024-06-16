package db

import (
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// db entities
var dbMap = map[string]*gorm.DB{}
var syncLock sync.Mutex

func init() {
	initDB("gochat")
}

func initDB(dbName string) {
	var err error
	realPath, _ := filepath.Abs("./")
	//主文件执行，需加db
	configFilePath := realPath + "/db/gochat.sqlite3"
	println("dbpath: %s\n", configFilePath)
	syncLock.Lock()
	dbMap[dbName], err = gorm.Open(sqlite.Open(configFilePath), &gorm.Config{})
	sqlDB, err := dbMap[dbName].DB()
	sqlDB.SetMaxIdleConns(4)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(8 * time.Second)
	syncLock.Unlock()
	if err != nil {
		logrus.Error("open db error")
	}
}

func GetDb(dbName string) (db *gorm.DB) {
	if db, ok := dbMap[dbName]; ok {
		return db
	} else {
		return nil
	}
}

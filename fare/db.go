package fare

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func getDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("open sqlite failed: " + err.Error())
	}
	return db
}

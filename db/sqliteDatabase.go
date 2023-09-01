package db

import (
	"Go_Web_Scrapper/logfuncs"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database = connectDB()
var Docs = DocumentsCount()

type FileDB struct {
	gorm.Model
	Url  string `gorm:"index;unique;not null"`
	Type string
}

func connectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{Logger: silentLogger})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&FileDB{})
	return db
}

func AddFileToDb(file *FileDB) error {
	retries := 0

	for {
		result := Database.Create(&file)

		if err := result.Error; err != nil {
			if strings.Contains(err.Error(), "SQLITE_BUSY") {
				time.Sleep(time.Second)
				retries += 1
				logfuncs.Logger.Warnln("[E] DataBase Connection Error: Retry: ", retries)
				continue
			} else if strings.Contains(err.Error(), "UNIQUE constraint failed:") {
				return nil
			}
			panic(err)
		}
		logfuncs.Logger.Infoln("[+] Item Added To DB \t ", Docs, " :", truncateTextFromEnd(file.Url, 75))
		Docs++
		return nil
	}

}

func truncateTextFromEnd(text string, maxLength int) string {
	if len(text) > maxLength {
		startIndex := len(text) - maxLength
		return "..." + text[startIndex:]
	}
	return text
}

var silentLogger = logger.New(
	nil, // Use your preferred io.Writer if you want to log somewhere
	logger.Config{
		LogLevel: logger.Silent,
	},
)

func DocumentsCount() int64 {
	var documentCount int64
	Database.Model(&FileDB{}).Count(&documentCount)
	return documentCount
}

package dbmodel

import (
	"fmt"

	"github.com/bihari123/terminal-based-game-in-go/config"
	"github.com/bihari123/terminal-based-game-in-go/loghelper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

type DB struct {
	DB *gorm.DB
}

func Init(config *config.Configuration) (*DB, error) {

	db, err := InitializeDatabase(config)

	if err != nil {
		loghelper.LogError("Error in InitializeDatabase")
		return nil, err
	}
	DBInstance = db
	return &DB{db}, nil
}

func InitializeDatabase(config *config.Configuration) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local",
		config.DBConfig.Username,
		config.DBConfig.Password,
		config.DBConfig.Databasehost,
		config.DBConfig.Schemaname)

	loghelper.Log("connection String", connectionString)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		loghelper.LogError("error opening connection", err.Error())
		return nil, err
	}

	db.AutoMigrate(&User{})

	loghelper.Log("Database Initialized")

	return db, nil
}

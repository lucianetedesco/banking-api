package core

import (
	"fmt"
	"github.com/lucianetedesco/banking-api/entities"
	"github.com/lucianetedesco/banking-api/settings"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var instance *DatabaseConnection
var once sync.Once

func GetDatabaseConnectionInstance() *DatabaseConnection {
	once.Do(func() {
		instance = NewDatabaseConnection(
			settings.Environment.Database.DBHost,
			settings.Environment.Database.DBPort,
			settings.Environment.Database.DBUser,
			settings.Environment.Database.DBPPassword,
			settings.Environment.Database.DBName,
		)
	})
	return instance
}

type DatabaseConnection struct {
	Db *gorm.DB
}

func NewDatabaseConnection(host string, port int, user string, pass string, dbName string) *DatabaseConnection {
	d := DatabaseConnection{Db: getDatabaseConnection(host, port, user, pass, dbName)}
	d.migrate()
	return &d
}

func (d DatabaseConnection) migrate() {
	err := d.Db.AutoMigrate(&entities.Account{})
	if err != nil {
		log.Panic("Failed to migrate database: ", err.Error())
	}
}

func getDatabaseConnection(host string, port int, user string, pass string, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=America/Sao_Paulo",
		host, user, pass, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect database: ", err.Error())
	}
	return db
}

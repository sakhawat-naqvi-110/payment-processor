package config

import (
	"fmt"
	"go/payment-processor/pkg/utils"
	"gorm.io/driver/postgres"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var db *gorm.DB

func ConnectDB() {

	log := GetLogger()
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf(utils.DB_CONNECTION_URL, dbHost, dbUser, dbPassword, dbName, dbPort)
	log.Info("DB connection URL: " + dsn)
	//dsn := fmt.Sprintf(utils.DB_CONNECTION_URL, dbHost, dbUser, dbPass, dbName, dbPort)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: ZapLogger{},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "public.",
			SingularTable: true,
		},
	})
	db = database
	if err != nil {
		log.Fatal("failed to connect database", zap.Error(err))
	}
	log.Info("Database connected successfully")
}

func GetDb() *gorm.DB {
	return db
}

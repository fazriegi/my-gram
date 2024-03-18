package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper) *gorm.DB {
	host := viper.GetString("db.host")
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	name := viper.GetString("db.name")
	port := viper.GetInt32("db.port")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		host,
		username,
		password,
		name,
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}
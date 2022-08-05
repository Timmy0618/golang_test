package gorm

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() (*gorm.DB, error) {
	const config string = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	sources := fmt.Sprintf(config,
		viper.GetString("database.DATABASE_USERNAME"),
		viper.GetString("database.DATABASE_PASSWORD"),
		viper.GetString("database.DATABASE_HOST"),
		viper.GetString("database.DATABASE_PORT"),
		viper.GetString("database.DATABASE_NAME"),
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 1000, // Slow SQL threshold (ms)
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.
		New(mysql.Config{
			DSN: sources,
		}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("Connecting to database: %s\n", sources)

	return db, nil
}

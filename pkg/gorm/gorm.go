package gorm

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	fmt.Printf("Connecting to database: %s\n", sources)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: sources,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

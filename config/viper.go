package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Default() {

	basePath := viper.GetString("application.root")

	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(basePath + "/config")
	// viper.SetDefault("application.port", 8080)
	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}
	fmt.Println("設定檔讀取成功")
	// fmt.Println("application port = " + viper.GetString("database.DATABASE_USERNAME"))
}

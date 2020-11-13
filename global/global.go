package global

import (
	"log"
	"os"

	"github.com/hunght/gqlgen-sqlc-example/config"
	"github.com/spf13/viper"
)

const appRetValue = -1

// InitConfig function
func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("../../configs/")
	viper.AddConfigPath("configs/")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		os.Exit(appRetValue)
	}
	err := viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		os.Exit(appRetValue)
	}

}

// Conf declare
var Conf config.Configuration

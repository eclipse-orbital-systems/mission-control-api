package config

import (
	"errors"
	"fmt"
	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
	"log"
)

func initViper() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Print("No configuration file loaded - using environment variables")
	}

	viper.SetDefault("GORM_LOG", false)
	viper.SetDefault("MYSQL_PORT", "3306")
	viper.SetDefault("MYSQL_DATABASE", "mission_control")

	requiredConfigVars := []string{
		"MYSQL_HOST",
		"MYSQL_USERNAME",
		"MYSQL_PASSWORD",
	}
	if !viper.GetBool("SECURITY_DISABLED") {
		requiredConfigVars = append(requiredConfigVars, "ALLOW_ORIGINS")
	}

	for _, requiredConfigVar := range requiredConfigVars {
		if !viper.IsSet(requiredConfigVar) {
			message := fmt.Sprintf("Required configuration variable \"%s\" is not set!", requiredConfigVar)
			err := errors.New(message)
			log.Fatal(stacktrace.Propagate(err, "Could not initialize app configuration"))
		}
	}
}

func init() {
	initViper()
	initCors()
	initMySql()
}

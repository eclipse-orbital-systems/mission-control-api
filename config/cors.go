package config

import (
	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
	"strings"
)

var (
	Cors cors.Config
)

func initCors() {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	if viper.GetBool("SECURITY_DISABLED") {
		config.AllowAllOrigins = true
	} else {
		origins := strings.Split(viper.GetString("ALLOW_ORIGINS"), ",")
		config.AllowOrigins = origins
	}
	Cors = config
}

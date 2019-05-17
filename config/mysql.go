package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/palantir/stacktrace"
	"github.com/spf13/viper"
	"log"
)

var (
	MySql *gorm.DB
)

func initMySql() {
	host := viper.GetString("MYSQL_HOST")
	username := viper.GetString("MYSQL_USERNAME")
	password := viper.GetString("MYSQL_PASSWORD")
	database := viper.GetString("MYSQL_DATABASE")
	gormLog := viper.GetBool("GORM_LOG")

	connectionString := username + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8&parseTime=True&loc=UTC"

	var err error
	MySql, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(stacktrace.Propagate(err, "failed to dial mysql"))
	}
	MySql.LogMode(gormLog)
	MySql.SingularTable(true)
}
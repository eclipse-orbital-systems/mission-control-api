module github.com/eclipse-orbital-systems/mission-control-api

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/jinzhu/gorm v1.9.8
	github.com/palantir/stacktrace v0.0.0-20161112013806-78658fd2d177
	github.com/spf13/viper v1.3.2
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	gopkg.in/go-playground/validator.v8 v8.18.2
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43

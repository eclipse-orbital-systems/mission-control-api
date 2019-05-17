package main

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"html"
	"reflect"
	"time"
)

func initValidators() {
	v, ok := binding.Validator.Engine().(*validator.Validate)

	if !ok {
		panic("failed to init validators")
	}

	if err := v.RegisterValidation("rfc3339", validRFC3339); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("rfc3339nano", validRFC3339); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("xss", validNotXss); err != nil {
		panic(err)
	}
}

func validRFC3339(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value,
	field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	if _, err := time.Parse(time.RFC3339, field.String()); err == nil {
		return true
	}

	return false
}

func validRFC3339Nano(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value,
	field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	if _, err := time.Parse(time.RFC3339Nano, field.String()); err == nil {
		return true
	}

	return false
}

func validNotXss(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value,
	field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	return field.String() == html.EscapeString(field.String())
}
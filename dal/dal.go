package dal

import (
	"github.com/jinzhu/gorm"
)

type Dal struct {
	mySql *gorm.DB
}

func New(mySql *gorm.DB) *Dal {
	obj := Dal{}
	obj.mySql = mySql

	return &obj
}

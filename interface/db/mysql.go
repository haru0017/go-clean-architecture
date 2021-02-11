package db

import (
	"github.com/jinzhu/gorm"
)

// SQLHandler はinfrastructureのメソッドをinterface層で使うためのinterface
type SQLHandler interface {
	Find(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
}

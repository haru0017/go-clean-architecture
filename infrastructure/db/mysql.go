package db

import (
	"fmt"
	"github.com/haru0017/go-clean-architecture/interface/db"
	"github.com/jinzhu/gorm"
	"time"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLHandler はgormを扱うための型
type SQLHandler struct {
	Conn *gorm.DB
}

// NewSQLHandler は新しいSQLHandlerを生成する
func NewSQLHandler() db.SQLHandler {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(mysql:3306)"
	DBNAME := "sample"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	WAIT := 30

	conn, err := open(DBMS, CONNECT, WAIT)
	if err != nil {
		panic(err)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

// open はMySQLの立ち上がりを確認し、接続する
func open(dbms string, connect string, count int) (*gorm.DB, error) {
	db, err := gorm.Open(dbms, connect)
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("Connection time error")
		}
		time.Sleep(1 * time.Second)
		count--
		return open(dbms, connect, count)
	}
	return db, nil
}

// Find は可変長引数whereの条件を元にデータを引っ張ってくるメソッド
func (handler *SQLHandler) Find(user interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(user, where...)
}

// Create は新規ユーザー作成メソッド
func (handler *SQLHandler) Create(user interface{}) *gorm.DB {
	return handler.Conn.Create(user)
}

// Save は変更を保存するメソッド
func (handler *SQLHandler) Save(user interface{}) *gorm.DB {
	return handler.Conn.Save(user)
}

// Delete はuserを削除するメソッド
func (handler *SQLHandler) Delete(user interface{}) *gorm.DB {
	return handler.Conn.Delete(user)
}

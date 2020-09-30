package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"


)

// SQLHandler はgormを扱うための型
type SQLHandler struct {
	Conn *gorm.DB
}

// NewSQLHandler は新しいSQLHandlerを生成する
func NewSQLHandler() *SQLHandler {
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
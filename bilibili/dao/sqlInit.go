package dao

import (
	"database/sql"
	"fmt"
	"time"
)

var DB *sql.DB
//
func MysqlInit() *sql.DB {
	db, err := sql.Open("winter", "root:root@tcp(127.0.0.1:3306)/messageBoard?charset=utf8")
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	//连接池
	//空闲
	db.SetMaxIdleConns(20)
	//打开
	db.SetMaxOpenConns(1000)
	//超时
	db.SetConnMaxLifetime(time.Second*20)
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	DB = db
	return DB
}
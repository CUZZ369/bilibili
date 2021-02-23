package models

import (
	"fmt"
	"test01/dao"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(username string, password string) bool {
	stmt, err := dao.DB.Prepare("insert into user(username,password) values (?,?)")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	return true
}

func Login(username string, password string) bool {
	stmt, err := dao.DB.Query("select * from user where username = ?", username)
	if err != nil {
		fmt.Printf("query failed:%v", err)
		return false
	}
	defer stmt.Close()
	var u user
	for stmt.Next() {
		err = stmt.Scan(&u.Username, &u.Password)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	if u.Password == password {
		return true
	}
	return false
}

type UserDetail struct {
	Username string
	Works string
	Liker int
	Vip bool
}

func ShowUser(userName string)UserDetail{
	sqlStr := "select works,likes,vip from user where username=?"
	var u UserDetail
	err := dao.DB.QueryRow(sqlStr, userName).Scan(&u.Works, &u.Liker, &u.Vip)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}
	return u
}

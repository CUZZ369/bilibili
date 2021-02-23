package main

import (
	"test01/cmd"
	"test01/dao"
)

func main() {
  dao.MysqlInit()//连接数据库
  cmd.Entrance()
}

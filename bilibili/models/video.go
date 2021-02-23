package models

import (
	"fmt"
	"test01/dao"
)

type Video struct {
	Id uint
	Title string
	Info  string
	CreateTime int64
}

type CreateVideoService struct {
	Title string`from:"Title" jason:"title" binding:"required,min=1,max=36"`
	Info string`from:"Info" jason:"Info" binding:"required,min=0,max=360"`
	Username string`from:"Info" jason:"Info" binding:"required,min=1,max=30"`
}

type VideoDetail struct {
	Id int
	Admin string
	Title string
	Info string//视频简介
	NumberPlays uint//播放次数
	LikeNumbers int //点赞次数
	Hot bool//是否热门
}
//上传视频
func CreateVideo(th *CreateVideoService)bool {
	stmt, err := dao.DB.Prepare("insert into video(th.title,th.info,th.username) values (?,?,?)")
	if err != nil {
		fmt.Printf("mysql prepare failed:%v", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(th.Title,th.Info)
	if err != nil {
		fmt.Printf("insert failed:%v", err)
		return false
	}
	return true
}
//查询视频
func ShowVideoId(id int) VideoDetail{
	sqlStr := "select Id, Admin, Title, Info,NumberPlays,LikeNumbers from user where id=?"
	var u VideoDetail
	err := dao.DB.QueryRow(sqlStr, id).Scan(&u.Admin, &u.Title, &u.Info,&u.LikeNumbers,&u.NumberPlays)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
	}
	return u
}
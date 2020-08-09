package services

import "github.com/gangjun06/gChat/client/lib/db"

func GetUserInfo() (data db.UserInfo) {
	db.DB().First(&data)
	return
}

func SetUserInfo(username, avatar string) {
	var data db.UserInfo
	db.DB().First(&data)
	data.Username = username
	data.Avatar = avatar
	db.DB().Save(&data)
	ServerView.Refresh()
}

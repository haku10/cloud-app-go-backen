package service

import (
	"restapi/app/model"
)

// Regist DBへの登録処理
func Regist(name string) model.User {
	connect := GormConnect()
	defer connect.Close()
	user := model.User{}
	user.Name =name
	connect.Create(&user)
	return user
}

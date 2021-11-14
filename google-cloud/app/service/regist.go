package service

import (
	"math/rand"
	"restapi/app/model"
	"strconv"
)

// Regist DBへの登録処理
func Regist() string {
	connect := GormConnect()
	defer connect.Close()
	user := model.User{}
	var num = strconv.Itoa(rand.Intn(100))
	user.Name = "regist test" + num
	connect.Create(&user)
	return user.Name
}

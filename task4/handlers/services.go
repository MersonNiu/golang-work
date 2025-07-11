package handlers

import (
	"GoBlogProject/config"
	"GoBlogProject/models"
	"fmt"
)

func AutoSetTable() {
	err := config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		panic("自动建表失败: " + err.Error())
	}
	fmt.Println("自动建表成功")
}

//插入数据的函数，入参是
// func CreateUser(name, email string) {
// 	user := models.User{Name: name, Email: email}
// 	result := config.DB.Create(&user)
// 	if result.Error != nil {
// 		fmt.Println("插入失败：", result.Error)
// 	} else {
// 		fmt.Printf("插入成功，ID=%d\n", user.ID)
// 	}
// }

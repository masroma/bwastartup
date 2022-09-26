package main

import (
	"bwastartup/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:@tcp(127.0.0.1:3306)/donasi?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("koneksi database berhasil")

	// var users []user.User
	// ini get awal masih kosong
	// length := len(users)
	// fmt.Println(length)

	// fmt.Println("====================")

	// db.Find(&users)
	// setelah di get dr db
	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.ID)
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// }

	router := gin.Default()
	router.GET("/users", handler)
	router.Run()

}

func handler(c *gin.Context) {
	dsn := "root:@tcp(127.0.0.1:3306)/donasi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)

}

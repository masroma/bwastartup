package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/donasi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepocitory := user.NewRepocitory(db)

	// tes
	user := user.User{
		Name:       "Test Simpan aja",
		Occupation: "IT Analyst",
		Email:      "roma@gmail.com",
		Password:   "password",
		Avatar:     "usera.jpg",
		Role:       "user",
	}

	userRepocitory.Save(user)

}

// proses nya
// input = user melakukan input
// handler = menangkap inputan dr user, di map ke sebuah object atau struct, buat struct baru yg di input sama user
// service = mengambil object mapping yg udah di buat d handler lalu di masukan k struct Induk yg ada di entity
//repository = save data ke db, ini untuk interaksi k db

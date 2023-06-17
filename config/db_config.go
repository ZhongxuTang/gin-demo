package config

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:lemonsoul@tcp(127.0.0.1:3306)/ry-config?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		logrus.Println("------------------- init db success! -------------------")
	} else {
		logrus.Warningln("------------------- init db fail! -------------------", err)
	}
}

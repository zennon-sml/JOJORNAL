package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func fazerCon() *gorm.DB { //*sql.DB

	banco := "root:asdf@tcp(172.17.0.2:3306)/jojornal?charset=utf8mb4&parseTime=True&loc=Local"
	bd, err := gorm.Open(mysql.Open(banco), &gorm.Config{})
	if err != nil {
		log.Fatal("não foi possivel estabelecer conexão ", err)
	}
	return bd
}

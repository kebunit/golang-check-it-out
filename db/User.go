package mydb

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	Id int `sql:"AUTO_INCREMENT"`
	NomorHP string `sql:"varchar(16); unique"`
	Username string `sql:varchar(50)`
	Address string `sql:type:varchar(300)`
}

func init() {
	GetDB().AutoMigrate(&User{})
}


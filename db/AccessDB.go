package mydb

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB
var err error

func GetDB() *gorm.DB {
	if database == nil {
		OpenDB()
	}
	return database
}

func OpenDB()  {
	fmt.Println("Opening DB")
	database, err = gorm.Open("postgres", "host=localhost user=postgres password=54bit dbname=TMN  sslmode=disable")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB Opened!")
	}
	database.DB()
	database.DB().Ping()
	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)
}


func InsertDB(value interface{})  {
	OpenDB()
	database.Create(value)
	err:= database.Close()
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println("DB Closed!")
	}

}
package mydb

import (
  "fmt"
  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func MyDBTest() {
  	var db *gorm.DB
	var err error
	db, err = gorm.Open("postgres", "host=localhost user=postgres password=54bit dbname=TMN  sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Masooook Pak Ekooooo");
	}
	defer db.Close()
}
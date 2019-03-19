package mydb

type Version struct {
	Id int `sql:"AUTO_INCREMENT"`
	VersionName string `sql:varchar(50)`
	Details string `sql:type:varchar(300)`
}

func init() {
	GetDB().AutoMigrate(&Version{})
}
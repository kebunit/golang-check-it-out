package network

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/kebunit/golang-check-it-out/db"
)

func UserController(response http.ResponseWriter, request *http.Request) {
	user :=mydb.User{}
	mydb.GetDB().First(&user)
	data := map[string]string{}

	data["name"]=user.Username
	data["noHP"]=user.NomorHP
	data["address"]=user.Address
	data["ACK"]="OK"
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(response, "%s",json)
}

func HomeController(response http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(response, "Welcome!!!")
	// http.Redirect(response, request, "http://localhost", 301)
}

func InsertVersion(context *gin.Context)  {
	version:=mydb.Version{VersionName:"4.35.35", Details:"Perubahan pada sesuatu yang buruk"}
	context.JSON(200, version)
}
package network

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

func UserController(response http.ResponseWriter, request *http.Request) {
	data := map[string]string{}

	data["name"]="Sabituddin Bigbang"
	data["noHP"]="0909090909090909"
	data["city"]="Jakarta"
	data["country"]="Indonesia"
	data["postal_code"]="121093"
	data["ACK"]="OK"
	json, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(response, "%s",json)
}

func HomeController(response http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(response, "Welcome!!!")
	 http.Redirect(response, request, "http://localhost", 301)
}
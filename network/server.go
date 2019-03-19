package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var ginEngine *gin.Engine

func init() {
	e := GetEngine()
	e.POST("/version/:version", InsertVersion)
}


func BasicAuth(handler http.HandlerFunc, username, password, realm string) http.HandlerFunc {
    return func(response http.ResponseWriter, request *http.Request) {
        user, pass, ok := request.BasicAuth()
        if !ok || (user != username) || (pass != password) {
            response.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
            response.WriteHeader(401)
            response.Write([]byte("Unauthorised.\n"))
            return
        }
        handler(response, request)
    }
}

func Router(pattern string, handler http.HandlerFunc){
    http.HandleFunc(pattern, func(response http.ResponseWriter, request *http.Request) {
    	fmt.Printf("[HOST] : %s%s\n",request.Host, request.RequestURI)
        handler(response, request)
    })
}

func ServerON() {
	Router("/", BasicAuth(HomeController, "sabit", "54bit", "Please enter your username and password for this site"))
	Router("/user", UserController)
	fmt.Printf("Server running on http://localhost:6969\n")
	log.Fatal(http.ListenAndServe(":6969", nil))
}


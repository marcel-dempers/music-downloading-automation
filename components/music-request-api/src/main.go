package main


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"fmt"
	"os"
)

var environment = os.Getenv("ENVIRONMENT")

func cors(writer http.ResponseWriter) () {
	if(environment == "DEBUG"){
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MY-API-Version")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}


func main() {
	config := GetConfiguration()
	router := httprouter.New()

	router.POST("/submit/*query", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		cors(w) 
		err := SubmitRequest(w,r,p,config)

		if(err != nil){
			http.Error(w, err.Error(), 500)
		}
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}

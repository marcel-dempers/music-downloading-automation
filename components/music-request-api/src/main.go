package main


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"fmt"
	//"app/models"
)
func main() {
	config := GetConfiguration()
	router := httprouter.New()

	router.POST("/submit/*query", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		 
		err := SubmitRequest(w,r,p,config)

		if(err != nil){
			http.Error(w, err.Error(), 500)
		}
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}

package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"app/models"
	"fmt"
)

var config *models.Configuration 

func main() {
	c := GetConfiguration()
	config = c

	router := httprouter.New()

	router.GET("/songs", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		songs, err := Songs(w,r,p)
		if(err != nil){
			http.Error(w, err.Error(), 500)
		}

		fmt.Fprintf(w, "%s", songs)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}
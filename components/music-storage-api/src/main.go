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
	InitStorage()
	router := httprouter.New()

	router.GET("/songs/all", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		songs, err := SongsAll(w,r,p)
		if(err != nil){
			http.Error(w, err.Error(), 500)
		}

		fmt.Fprintf(w, "%s", songs)
	})

	router.GET("/songs/search", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		songs, err := SongsSearch(w,r,p)
		if(err != nil){
			http.Error(w, err.Error(), 500)
		}

		fmt.Fprintf(w, "%s", songs)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}
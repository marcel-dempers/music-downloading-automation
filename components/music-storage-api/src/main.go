package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"app/models"
	"fmt"
	"os"
)

var config *models.Configuration 
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
	c := GetConfiguration()
	config = c
	InitStorage()
	router := httprouter.New()

	router.GET("/songs/all", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		songs, err := SongsAll(w,r,p)
		if(err != nil){
			http.Error(w, err.Error(), 500)
		}

		cors(w)
		fmt.Fprintf(w, "%s", songs)
	})

	router.GET("/songs/search", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		songs, err := SongsSearch(w,r,p)
		if(err != nil){
			http.Error(w, err.Error(), 500)
		}
		cors(w)
		fmt.Fprintf(w, "%s", songs)
	})

	router.GET("/song/byurl", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		songs, err := SongByUrl(w,r,p)
		if(err != nil){
			http.Error(w, err.Error(), 500)
		}
		cors(w)
		fmt.Fprintf(w, "%s", songs)
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}
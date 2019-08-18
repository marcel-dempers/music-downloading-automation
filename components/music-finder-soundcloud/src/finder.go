package main

import (
	"fmt"
	"net/http"
	"app/models"
	"github.com/julienschmidt/httprouter"
	"errors"
)

func FindSong(writer http.ResponseWriter, request *http.Request, p httprouter.Params, config models.Configuration) (err error) {
	fmt.Println("Submitting request to queue")

	count := p.ByName("count")
	depth := p.ByName("depth")
	query := p.ByName("query")
	
	query = query[1:]
	//query/http://blah.test.com/test/blah
	fmt.Println("Received: " + query)
	fmt.Println("Count: " + count)
	fmt.Println("Depth: " + depth)
	
	if query == "" {
		return errors.New("Expeced query parameter")
	}

	var trackID = "353521499"

	trackList , err := sc_client.GetRelatedTracks(trackID)

	if err != nil {
		panic(err)
	}

	fmt.Println(trackList)

	return err
}
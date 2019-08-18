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
	trackID := p.ByName("trackid")
	
	//trackID = trackID[1:]
	//trackID/http://blah.test.com/test/blah
	fmt.Println("Received: " + trackID)
	fmt.Println("Count: " + count)
	fmt.Println("Depth: " + depth)
	
	if trackID == "" {
		return errors.New("Expeced trackID parameter")
	}

	trackList , err := sc_client.GetRelatedTracks(trackID)

	if err != nil {
		panic(err)
	}

	fmt.Println(trackList)

	return err
}
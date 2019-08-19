package main

import (
	"fmt"
	"net/http"
	"app/models"
	"github.com/julienschmidt/httprouter"
	"errors"
)

func FindSong(writer http.ResponseWriter, request *http.Request, p httprouter.Params, config models.Configuration) (err error) {
	count := p.ByName("count")
	depth := p.ByName("depth")
	trackUrl := p.ByName("trackurl")
	
	//trackID = trackID[1:]
	//trackID/http://blah.test.com/test/blah
	fmt.Println("Received: " + trackUrl)
	fmt.Println("Count: " + count)
	fmt.Println("Depth: " + depth)
	
	if trackUrl == "" {
		return errors.New("Expeced trackUrl parameter")
	}

	//var url = "https%3A%2F%2Fsoundcloud.com%2Fmsmrsounds%2Fms-mr-hurricane-chvrches-remix"
	trackList , err := sc_client.GetRelatedTracksByUrl(trackUrl)
	// trackid , err := sc_client.GetTrackIdFromUrl(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(trackList)

	return err
}


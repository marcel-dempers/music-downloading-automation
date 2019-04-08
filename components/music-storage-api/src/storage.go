package main

import (
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"app/models"
	"io/ioutil"
)

var storageClient *http.Client
var storageUri string

func InitStorage() {

	storageClient = &http.Client{
		Timeout: 10 * time.Second, //10 sec
	}

	storageUri =  "http://" + config.CouchDB.Host + ":" +  strconv.Itoa(config.CouchDB.Port)

}

//TODO: implement query search
//Returns rows from songs view - can be queried, not used yet
func Storage_SongList_Search(query string) (songlist models.Songlist, err error) {
		
	req, err := http.NewRequest("GET", storageUri + "/mydatabase/_design/songlist_view/_view/main_songlist", nil)
	
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	resp, err := storageClient.Do(req)
    songlist = models.Songlist{}
	if err != nil {
		return songlist, err
	}

	if resp != nil {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return songlist, err
		}

		err = json.Unmarshal(bodyBytes, &songlist)
		
		if err != nil {
			return songlist, err
		}

		return songlist, nil
	}

	return songlist, err 
}

///Returns rows from songs view
func Storage_SongList_All() (songlist models.Songlist, err error) {
	
	req, err := http.NewRequest("GET", storageUri + "/mydatabase/_design/songlist_view/_view/main_songlist", nil)
	
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	resp, err := storageClient.Do(req)
    songlist = models.Songlist{}
	if err != nil {
		return songlist, err
	}

	if resp != nil {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return songlist, err
		}

		err = json.Unmarshal(bodyBytes, &songlist)
		
		if err != nil {
			return songlist, err
		}

		return songlist, nil
	}

	return songlist, err 
}
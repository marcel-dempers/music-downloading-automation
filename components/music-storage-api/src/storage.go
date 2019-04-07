package main

import (
	"net/http"
	"strconv"
	"io/ioutil"
	"time"
)

///Returns rows from songs view
func GetAllSongs() (responseJSON string, err error) {
	
	req, err := http.NewRequest("GET", "http://" + config.CouchDB.Host + ":" +  strconv.Itoa(config.CouchDB.Port) + "/mydatabase/_design/songlist_view/_view/main_songlist", nil)
	
	client := &http.Client{
		Timeout: 10 * time.Second, //10 sec
	}

	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp != nil {
		defer resp.Body.Close()

		if b, err := ioutil.ReadAll(resp.Body); err == nil {
			return string(b), nil
		}
	}

	return "", err 
}
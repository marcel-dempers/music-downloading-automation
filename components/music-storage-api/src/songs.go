package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

func Songs(writer http.ResponseWriter, request *http.Request, p httprouter.Params) (songsJSON string, err error) {
	
	// u, err := url.Parse(request.URL.String())
	// if err != nil {
	// 	return rows, err
	// }

	//query := u.Query()
	//fmt.Println(query)
    
	songlist, err := GetAllSongs()
	if err != nil {
		return "", err
	}

	songlistBytes, err := json.Marshal(songlist)
	if err != nil {
		return "", err
	}

	r := bytes.NewReader(songlistBytes)
	if b, err := ioutil.ReadAll(r); err == nil {
			return string(b), nil
	}

	//fmt.Printf("%+v\n",rows)
	return "", err
}

package main

import (
	//"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	 
)

func Songs(writer http.ResponseWriter, request *http.Request, p httprouter.Params) (songsJSON string, err error) {
	
	// u, err := url.Parse(request.URL.String())
	// if err != nil {
	// 	return rows, err
	// }

	//query := u.Query()
	//fmt.Println(query)

	return GetAllSongs()

	//fmt.Printf("%+v\n",rows)

}

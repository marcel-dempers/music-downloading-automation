package main

import (
	"fmt"
	"net/http"
	"app/models"
	"github.com/julienschmidt/httprouter"
	"errors"
)

func SubmitRequest(writer http.ResponseWriter, request *http.Request, p httprouter.Params, config models.Configuration) (err error) {
	fmt.Println("Submitting request to queue")

	query := p.ByName("query")
	///query/http://blah.test.com/test/blah

	fmt.Println("Received: " + query)
	message, err := ParseQueryToMessage(query)
	fmt.Print(message)
	if err != nil {
		return err
	}

	if query == "" {
		return errors.New("Expected query parameter")
	}

	produceMessage(message, config.RabbitMq)


	if err != nil {
		panic(err)
	}

	return err
}
//Converts incoming query to message for queue
func ParseQueryToMessage(queryfull string) (message models.Message , err error){
	runes := []rune(queryfull)
	query := string(runes[0:7])
	queryVal := queryfull[7:len(queryfull)]
	fmt.Println(query)
	fmt.Println(queryVal)

	if query != "/query/" {
		return models.Message{SongUri : "",}, errors.New("Expected query parameter")
	}
	if err != nil {
		panic(err)
	}
	msg := models.Message{SongUri : queryVal}
	 
	return msg, err
}
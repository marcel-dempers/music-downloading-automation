package main


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)
func main() {
	config := GetConfiguration()
	router := httprouter.New()

	router.POST("/deploy/:environment/:namespace/:servicename/:dockerimagetag", func(w http.ResponseWriter, r *http.Request, p httprouter.Params){
		err := Deploy(w,r,p,config.Environments)

		if(err != nil){
			http.Error(w, err.Error(), 500)
		}
	})

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":10010", router))
}

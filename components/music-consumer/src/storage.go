package main


import (
	"fmt"
	"github.com/flimzy/kivik"
    _ "github.com/go-kivik/couchdb"
	"time"
	"github.com/rs/xid"
	"bytes"
)

func ConnectAndSaveContent(filename string, content []byte) {

	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("music-storage",5984,timeout)
	auth := couchdb.BasicAuth{Username: "user", Password: "password" }
	db := conn.SelectDB("mydatabase", &auth)
	id := xid.New()
	reader := bytes.NewReader(content)

	rev, err := db.Save("empty", id.String(), "") 
	rev, err = db.SaveAttachment(id.String(), "", "test" , "Application/audio", reader )
	fmt.Println(rev)
	if err != nil {
		panic(err)
	}
}

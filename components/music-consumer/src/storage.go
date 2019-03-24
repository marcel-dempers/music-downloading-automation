package main


import (
	"context"
	"fmt"
	"github.com/go-kivik/kivik"
     _ "github.com/go-kivik/couchdb"
	"github.com/rs/xid"
	"bytes"
	"io"
)

type ClosingBuffer struct { 
	*bytes.Reader 
} 

func (cb *ClosingBuffer) Close() (err error) { 
	//we don't actually have to do anything here, since the buffer is 
	//just some data in memory 
	//and the error is initialized to no-error 
	return 
} 

func ConnectAndSaveContent(filename string, content []byte) {

	//var timeout = time.Duration(500 * time.Millisecond)
	//conn, err := couchdb.NewConnection("music-storage",5984,timeout)
	//auth := couchdb.BasicAuth{Username: "user", Password: "password" }
	//db := conn.SelectDB("mydatabase", &auth)
	
	client, err := kivik.New("couch", "http://music-storage:5984/")
	if err != nil {
        panic(err)
    }
	
	db := client.DB(context.TODO(), "mydatabase")
 
	id := xid.New()
	reader := bytes.NewReader(content)
	
	doc := map[string]interface{}{
        "_id":      id.String(),
        "filename":     filename,
	}
	
	cb := &ClosingBuffer{reader}
	var rc io.ReadCloser
	rc = cb 
	defer rc.Close()
	attachment := &kivik.Attachment{Filename : filename, Content : rc, ContentType : "audio/mpeg" }

	rev, err := db.Put(context.TODO(), id.String(), doc)
    if err != nil {
        panic(err)
	}

	rev, err = db.PutAttachment(context.TODO(),id.String(), rev, attachment)
	if err != nil {
        panic(err)
	}
	
	//rev, err := db.Save("empty", id.String(), "") 
	//rev, err = db.SaveAttachment(id.String(), "", "test" , "Application/audio", reader )
	fmt.Println(rev)
}

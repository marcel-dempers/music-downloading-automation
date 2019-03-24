package main


import (
	"context"
	"fmt"
	"github.com/go-kivik/kivik"
     _ "github.com/go-kivik/couchdb"
	"github.com/rs/xid"
	"bytes"
	"io"
	"app/models"
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

func ConnectAndSaveContent(document models.Document, content []byte, metadatafile []byte) {

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
	readerMeta := bytes.NewReader(metadatafile)
	
	doc := map[string]interface{}{
        "_id":      id.String(),
		"fileName":     document.FileName,
		"localFilePath" : document.LocalFilePath,
		"metadataFilePath" : document.MetadataFilePath,
		"metadataFileName" : document.MetadataFileName,
	}
	
	cb := &ClosingBuffer{reader}
	cbMeta := &ClosingBuffer{readerMeta}

	var rc io.ReadCloser
	var rcMeta io.ReadCloser
	rc = cb
	rcMeta = cbMeta 

	defer rcMeta.Close()
	defer rc.Close()

	attachment := &kivik.Attachment{Filename : document.FileName, Content : rc, ContentType : "audio/mpeg" }
	attachmentMeta := &kivik.Attachment{Filename : document.MetadataFileName, Content : rcMeta, ContentType : "application/json" }
	rev, err := db.Put(context.TODO(), id.String(), doc)
    if err != nil {
        panic(err)
	}

	rev, err = db.PutAttachment(context.TODO(),id.String(), rev, attachment)
	if err != nil {
        panic(err)
	}
	rev, err = db.PutAttachment(context.TODO(),id.String(), rev, attachmentMeta)
	if err != nil {
        panic(err)
	}
	
	//rev, err := db.Save("empty", id.String(), "") 
	//rev, err = db.SaveAttachment(id.String(), "", "test" , "Application/audio", reader )
	fmt.Println(rev)
}

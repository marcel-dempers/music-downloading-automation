package main


import (
	"context"
	"fmt"
	"github.com/go-kivik/kivik"
     _ "github.com/go-kivik/couchdb"
	"bytes"
	"io"
	"app/models"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
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

	client, err := kivik.New("couch", "http://music-storage:5984/")
	if err != nil {
        panic(err)
    }
	
	db := client.DB(context.TODO(), "mydatabase")
 
	reader := bytes.NewReader(content)
	readerMeta := bytes.NewReader(metadatafile)
	
	metadata, err := ioutil.ReadFile(document.MetadataFilePath)
	if err != nil {
		fmt.Printf("Read metadata file error: %v\n", err)
		panic(err)
	}
  
	var meta models.MetadataFile
	err = json.Unmarshal(metadata, &meta)
	  
	if err != nil {
	  fmt.Printf("Problem in metatdata file: %v\n", err.Error())
	  panic(err)
	}

	if appConfig.Metadata.License.NcsAutodetect == true {
		if meta.Uploader == "NoCopyrightSounds" {
			meta.License = "ncs"
		}
	}
     
	id := base64.StdEncoding.EncodeToString([]byte(meta.Uploader + meta.ID))

	doc := map[string]interface{}{
		"_id":      id,
		"fileName":     document.FileName,
		"enrichDate" : "",
		"localFilePath" : document.LocalFilePath,
		"metadataFilePath" : document.MetadataFilePath,
		"metadataFileName" : document.MetadataFileName,
		"metadata" : meta,
	}
	
	cb := &ClosingBuffer{reader}
	cbMeta := &ClosingBuffer{readerMeta}

	var rc io.ReadCloser
	var rcMeta io.ReadCloser
	rc = cb
	rcMeta = cbMeta 

	defer rcMeta.Close()
	defer rc.Close()

	attachment := &kivik.Attachment{Filename : document.FileName, Content : rc, ContentType : "audio/wav" }
	attachmentMeta := &kivik.Attachment{Filename : document.MetadataFileName, Content : rcMeta, ContentType : "application/json" }
	
	//see if doc exists first
	fmt.Println("Checking doc...")

	var rev string
	row := db.Get(context.TODO(), id)
	rev = row.Rev
	if rev == "" {
		fmt.Println("Document does not exist, creating...")
		rev, err = db.Put(context.TODO(), id, doc)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("Document exists, updating...")
		doc["_rev"] = row.Rev
		rev, err = db.Put(context.TODO(), id, doc)
		if err != nil {
			panic(err)
		}
	}
	 
	if rev != "" {

		fmt.Println("Adding attachments..")
		rev, err = db.PutAttachment(context.TODO(),id, rev, attachment)
		if err != nil {
			panic(err)
		}
		rev, err = db.PutAttachment(context.TODO(),id, rev, attachmentMeta)
		if err != nil {
			panic(err)
		}

		fmt.Println("Binary data stored!")
		
		//rev, err := db.Save("empty", id.String(), "") 
		//rev, err = db.SaveAttachment(id.String(), "", "test" , "Application/audio", reader )
		fmt.Println(rev)
	}
	


}

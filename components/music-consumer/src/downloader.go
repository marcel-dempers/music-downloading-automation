package main

import (
	"fmt"
	"os/exec"
	"app/models"
	"time"
	"errors"
	"bytes"
	"bufio"
	"strings"
	//"regexp"
	"io/ioutil"
	"path/filepath"
	"github.com/h2so5/goback/regexp"
)

func ProcessMessage(message models.Message) (err error) {
	fmt.Println("Processing message")
	fmt.Println(message)
	
	//download song
	dl := Downloader{}
	output ,err := dl.exec([]string{"--write-info-json", "-x", "--audio-format", "wav", message.SongUri })
	//output ,err := dl.exec([]string{"--write-info-json", "-x", message.SongUri })
	if err != nil {
		panic(err)
	}
	
	//Process the StdOut of the downloader and get the path of the file downloaded
	var downloadedFilePath string

	scanner := bufio.NewScanner(strings.NewReader(output.String()))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if strings.Contains(line, "[ffmpeg] Destination") && strings.Contains(line, ".wav") {
			fmt.Println("Finding downloaded file on outputline: " + line)

			re := regexp.MustCompile(`\w*\/.*(?=\.)`)
			filePath := re.FindAllString(line, -1)
			
			if filePath == nil {
				//problem getting the destination of downloaded file
				return errors.New("Problem retrieve destination path from buffer")
			}

			fmt.Println("File saved to: " + filePath[0] + ".wav")
			downloadedFilePath = filePath[0] + ".wav"
		}
	}
	
	//attempt to store in couchdb
	if downloadedFilePath != "" {
		content, err := ioutil.ReadFile(downloadedFilePath)
		if err != nil {
			return errors.New("Problem reading downloaded file: " + downloadedFilePath)
		}
		
		filenoext := strings.Replace(downloadedFilePath, filepath.Ext(downloadedFilePath), "", 1)
		metadatafilepath :=  filenoext + ".info.json"
		document := models.Document{ LocalFilePath : downloadedFilePath , FileName : filepath.Base(downloadedFilePath) , FileNoExt : filenoext, MetadataFilePath : metadatafilepath, MetadataFileName :  filepath.Base(metadatafilepath)}
		//metadatafilepath := document.FileNoExt + ".info.json"

		fmt.Println("Reading metadatafile: " + document.MetadataFilePath)
		metadatafile, err := ioutil.ReadFile(document.MetadataFilePath)
		if err != nil {
			return errors.New("Problem reading metadata file")
		}

		ConnectAndSaveContent(document,content,metadatafile)
	}

	return err
}

type Downloader struct {}
type Exe struct {}
func (d Downloader) exec(args []string) (output bytes.Buffer , errr error) {
	e:= Exe{}
	return e.exec("youtube-dl", args, 120)
}

func (e Exe) exec(program string, args []string, timeoutInSec time.Duration) (output bytes.Buffer, err error){
	
	cmd := exec.Command(program, args...)
    
	var outputbuf, errbuf bytes.Buffer
    cmd.Stdout = &outputbuf
	cmd.Stderr = &errbuf
	// stderr, err := cmd.StderrPipe()
	// if err != nil {
	// 	fmt.Println("Exe returning error from StderrPipe...")
	// 	fmt.Print(err)
	// 	return err
	// }
	if err := cmd.Start(); err != nil {
		fmt.Println("Exe returning error from Start...")
		fmt.Print(err)
		return outputbuf, err
	}
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(timeoutInSec * time.Second)
	select {
	case <-timeout:
		cmd.Process.Kill()
		// slurp, _ := ioutil.ReadAll(stderr)
		// fmt.Printf("%s\n", slurp)
		return outputbuf, errors.New("There is a problem with the request")
	case err := <-done:
		//slurp, _ := ioutil.ReadAll(stderr)
		fmt.Println("Exe done")
		//fmt.Printf("%s\n", slurp)
		if err != nil {
			fmt.Println("Exe returned error after completion", err)
			
			println(outputbuf.String())
			println(errbuf.String())
			return outputbuf, errors.New(errbuf.String())
		}
	}

	fmt.Println("Exe processing complete")
	return outputbuf,nil
}

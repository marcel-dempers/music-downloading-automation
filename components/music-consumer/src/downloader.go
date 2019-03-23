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
	"regexp"
	"io/ioutil"
)

func ProcessMessage(message models.Message) (err error) {
	fmt.Println("Processing message")
	fmt.Println(message)
	
	dl := Downloader{}
	dl.exec([]string{message.SongUri})
	if err != nil {
		panic(err)
	}

	return err
}

type Downloader struct {}
type Exe struct {}
func (d Downloader) exec(args []string) (errr error) {
	e:= Exe{}
	return e.exec("youtube-dl", args, 120)
}

func (e Exe) exec(program string, args []string, timeoutInSec time.Duration) (err error){
	
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
		return err
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(timeoutInSec * time.Second)

	select {
	case <-timeout:
		cmd.Process.Kill()
		// slurp, _ := ioutil.ReadAll(stderr)
		// fmt.Printf("%s\n", slurp)
		return errors.New("There is a problem with the request")
	case err := <-done:
		//slurp, _ := ioutil.ReadAll(stderr)
		fmt.Println("Exe done")
		//fmt.Printf("%s\n", slurp)

		if err != nil {
			fmt.Println("Exe returned error after completion", err)
			
			println(outputbuf.String())
			println(errbuf.String())
			return errors.New(errbuf.String())
		}
	}

	//println(outputbuf.String())
    
	scanner := bufio.NewScanner(strings.NewReader(outputbuf.String()))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if strings.Contains(line, "[download]") {
			fmt.Println("Finding downloaded file")
			re := regexp.MustCompile(`\w*\/[^.]+[.]\w{1,3}`)
			filePath := re.FindAllString(line, -1)

			if filePath == nil {
				//problem getting the destination of downloaded file
				return errors.New("Problem retrieve destination path from buffer")
			}

			fmt.Println("File saved to: " + filePath[0])

			//attempt to store in couchdb
			content, err := ioutil.ReadFile(filePath[0])
			if err != nil {
				return errors.New("Problem reading downloaded file")
			}

			ConnectAndSaveContent(filePath[0],content)


		}
    }
	fmt.Println("Exe processing complete")
	return nil
}

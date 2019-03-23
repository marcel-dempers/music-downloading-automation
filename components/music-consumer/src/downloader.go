package main

import (
	"fmt"
	"os/exec"
	"app/models"
	"time"
	"errors"
	"bytes"
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

	println(outputbuf.String())

	fmt.Println("Exe processing complete")
	return nil
}


// func (e Exe) exec(yaml string, context string) (err error){
// 	echo := exec.Command("echo", []string{yaml}...)

// 	kubectl := exec.Command("kubectl", []string{"apply", "--context="+context, "-f", "-"}...)
// 	echoOut, _ := echo.StdoutPipe()
// 	echo.Start()
// 	kubectl.Stdin = echoOut

// 	var outputbuf bytes.Buffer
//     kubectl.Stdout = &outputbuf

// 	stderr, err := kubectl.StderrPipe()
// 	if err != nil {
// 		fmt.Println("kubeapply StderrPipe returning error")
// 		return err
// 	}

// 	if err := kubectl.Start(); err != nil {
// 		fmt.Println("kubeapply Start returning error")
// 		return err
// 	}

// 	slurp, _ := ioutil.ReadAll(stderr)
// 	fmt.Printf("%s\n", slurp)

// 	if err := kubectl.Wait(); err != nil {
// 		fmt.Println("kubeapply Wait returning error")
// 		return errors.New(fmt.Sprintf("%s\n", slurp))
// 	}

// 	println(outputbuf.String())
// 	fmt.Println("kubeapply returning...")
// 	return nil
// }
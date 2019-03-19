package main

import (
	"fmt"
	"os/exec"
	"app/models"
)


func SubmitRequest(writer http.ResponseWriter, request *http.Request, p httprouter.Params, models.Configuration) (err error) {
	fmt.Println("Submitting request to queue")
}
// func (d Downloader) exec(args []string) (errr error) {
// 	e:= Exe{}
// 	return e.exec("youtube-dl", args, 120)
// }

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
package main

import (
	"os/exec"
	"os"
	"fmt"
	"bytes"
	"time"
)

func EchoStr(){
	output, err := exec.Command("echo", "Executing a command in Go").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))

}

func GoEnv(){
	cmd := exec.Command("go", "env")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))
}

func excuAsync(){
	cmd := exec.Command("cat", "/dev/random")
	randomBytes := &bytes.Buffer{}
	cmd.Stdout = randomBytes

	// Start command asynchronously
	err := cmd.Start()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	ticker := time.NewTicker(time.Second)
	go func(ticker *time.Ticker) {
		now := time.Now()
		for _ = range ticker.C {
			fmt.Printf("Ticker: %s\n", []byte(fmt.Sprintf("%s", time.Since(now))))
		}
	}(ticker)

	fmt.Println("-------------------------")
	// Kill the process after 4 seconds
	timer := time.NewTimer(time.Second * 4)
	go func(timer *time.Timer, ticker *time.Ticker, cmd *exec.Cmd) {
		for _ = range timer.C {
			err := cmd.Process.Signal(os.Kill)
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}
			ticker.Stop()
		}
	}(timer, ticker, cmd)

	// Wait for the command to finish
	cmd.Wait()
	fmt.Printf("Result: %d\n", []byte(fmt.Sprintf("%d bytes generated", len(randomBytes.Bytes()))))
}


func main(){
	//EchoStr()   // Only work on linux platform
	GoEnv()
	excuAsync()  // Only work on linux platform
}




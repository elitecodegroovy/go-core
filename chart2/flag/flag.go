package main

import (
	"flag"
	"fmt"
)

func doBasicFlag(){
	versionPtr := flag.String("v", "1.9.2", " v version info")
	numberPtr := flag.Int("n", 1, "n parameter the concurrent times")
	boolPtr := flag.Bool("log", false, "log whether or not print log info ")

	var logPath string
	flag.StringVar(&logPath, "logpath", "./", "logpath  log file location")
	flag.Parse()

	fmt.Println("-v=", *versionPtr)
	fmt.Println("-n=", *numberPtr)
	fmt.Println("-log=", *boolPtr)
	fmt.Println("-logpath=", logPath)
	
	for i, a := range flag.Args(){
		fmt.Printf(" arg[%d]: %s", i, a)
	}
}

func main(){
	doBasicFlag()
	//>go build -o flag.exe
	//>flag.exe -v=2.0 -n=10 -log=true -logpath="D:\githubRepo\go\goapp>" a1 a2 a3
	//-v= 2.0
	//-n= 10
	//-log= true
	//-logpath= D:\githubRepo\go\goapp>
	//	arg[0]: a1 arg[1]: a2 arg[2]: a3

}

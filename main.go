package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var protocol string = os.Args[1]
	var adress string = os.Args[2]
	_, errs := net.Dial(protocol, adress)
	if errs == nil {
		fmt.Println("connection succssesful")
	} else {
		fmt.Println(errs)
	}

}

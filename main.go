package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var (
		protocol string = os.Args[1]
		adress   string = os.Args[2]
	)

	for i := 1; i <= 65535; i++ {
		fmt.Printf("scanning port:%d", i)
		var destination string = fmt.Sprintf("%s:%d", adress, i)
		connection, errs := net.Dial(protocol, destination)
		if errs == nil {
			fmt.Println(" connection succssesful")
			connection.Close()
		} else {
			fmt.Println(" connection refused")
			continue
		}
	}

}

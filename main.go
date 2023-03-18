package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func workerPool(ports <-chan int, result chan<- int, adress string) {
	for port := range ports {
		var destination string = fmt.Sprintf("%s:%d", adress, port)
		connection, errs := net.Dial("tcp", destination)
		if errs != nil {
			continue
		}
		result <- port
		connection.Close()
	}
}

func main() {
	var (
		adress string   = os.Args[1]
		ports  chan int = make(chan int, 200)
		result chan int = make(chan int, 200)
	)

	defer close(ports)
	defer close(result)

	fmt.Println("Possible numbers workers is: ", runtime.NumCPU())

	for i := 0; i < runtime.NumCPU(); i++ {
		go workerPool(ports, result, adress)
	}

	go func() {
		for port := 1; port <= 65535; port++ {
			ports <- port
		}
	}()

	for openPort := range result {
		fmt.Printf("Port %d is open\n", openPort)
	}

}

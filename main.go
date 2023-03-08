package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	var (
		adress string = os.Args[1]
		wg     sync.WaitGroup
	)

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			var destination string = fmt.Sprintf("%s:%d", adress, j)
			connection, errs := net.Dial("tcp", destination)
			if errs == nil {
				fmt.Printf("port:%d connection succssesful\n", j)
				connection.Close()
			}
		}(i)

	}
	wg.Wait()

}

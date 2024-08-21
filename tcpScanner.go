package main

import (
	"fmt"
	"net"
	"os"
	"sort"
)

func worker(ports, results chan int) {
	arguments := os.Args
	if len(arguments) <= 1 {
		fmt.Println("[-] host must be provided ")
		os.Exit(1)
	}

	for p := range ports {
		address := fmt.Sprintf("%v:%d", arguments[1], p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int, 100)
	var openPorts []int

	for i := 1; i <= cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d port is open \n", port)
	}

}

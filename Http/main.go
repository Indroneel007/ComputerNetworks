package main

import (
	"fmt"
	//"io"
	//"io/ioutil"
	"log"
	//"os"
	"net"
)

func main() {
	//println("Hello, World!")
	/*file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	//tyjtyjty
	b, err := ioutil.ReadAll(file)
	fmt.Printf("Data as string: %s\n", b)*/

	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Connection from server connected\n")

		_, err = conn.Write([]byte("Hello from server!\n"))

		if err != nil {
			log.Fatal(err)
		}
		conn.Close()
		fmt.Printf("Connection from server closed\n")
	}
}

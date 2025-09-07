package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//println("Hello, World!")
	file, err := os.Open("message.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	fmt.Printf("Data as string: %s\n", b)
}

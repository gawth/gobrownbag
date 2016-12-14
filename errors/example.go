package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(f.Name())

}

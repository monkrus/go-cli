package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log")
	level := flag.String("level", "ERROR", "Log level to search for: DEBUG, INFO, CRITICAL")
	flag.Parse()
	//open respective file
	f, err := os.Open(*path)
	//what if we have an  error
	if err != nil {
		log.Fatal(err)
	}
	// close it!
	defer f.Close()
	// read the fil
	r := bufio.NewReader(f)
	//indefinite loop
	// parse it to string
	for {
		s, err := r.ReadString('\n')
		//what if we have an error again?
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}

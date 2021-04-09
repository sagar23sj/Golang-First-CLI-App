package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	path := flag.String("path", "sample.log", "The path to the log file that should be analysed")
	level := flag.String("level", "ERROR", "Log-Level to search, for eg. ERROR. Option : [ ERROR, INFO, FATAL, DEBUG, WARN, TRACE ]")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		fmt.Println("error opening log file : %+v", err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)

	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// define a flag for specifying input file
	fnameP := flag.String("file", "data/sut_phones.source", "-file path/to/file")

	// parse passed flags
	flag.Parse()

	// read file contents
	b, err := os.ReadFile(*fnameP)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}

	// regex string that matches folowing phone numbers:
	// with country code - +1 (234) 567-89-01
	// without country code - (234) 567-89-01
	// city - 567-89-01
	re := `\+?\d* ?(?:\([0-9]{3}\) )?[0-9]{3}-[0-9]{2}-[0-9]{2}`
	regex, err := regexp.Compile(re)
	if err != nil {
		log.Fatalln("failed to compile regular expresion:", err)
	}

	// return all matches (-1 for all matches)
	matches := regex.FindAllString(string(b), -1)

	// print all matches
	for i, m := range matches {
		fmt.Printf("Found regex match %d: %s\n", i, m)
	}

}

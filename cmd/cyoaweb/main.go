package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Omkar-Patil16/cyoa.git"
)

func main() {
	// creating a flag to pass the filename to be parsed

	fileName := flag.String("file", "gopher.json", "file with adventure story")
	flag.Parse()

	fmt.Printf("Name Of the file selected %s\n", *fileName)

	// open the file

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonDecode(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)

}

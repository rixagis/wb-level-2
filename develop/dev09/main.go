package main

import (
	"log"
	"os"

	"github.com/rixagis/wb-level-2/develop/dev09/wget"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: wget <URL>")
	}

	url := os.Args[1]
	folder, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filename, err := wget.MakeFileName(url)
	if err != nil {
		log.Fatal(err)
	}
	filepath := folder + string(os.PathSeparator) + filename
	log.Println(filepath)
	wget.Wget(filepath, url)
}
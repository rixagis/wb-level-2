package main

import (
	"fmt"
	"os"
	
	"github.com/rixagis/wb-level-2/develop/dev02/unpacker"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments: only one argument supported.")
		os.Exit(1)
	}

	var input = os.Args[1]

	var unpacker = unpacker.NewUnpacker()
	
	var result, err = unpacker.Unpack(input)
	
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}

	fmt.Println(result)
}

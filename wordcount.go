package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatal(errors.New("need exactly 1 command-line argument"))
	}

	words := strings.Fields(args[1])
	wordsCount := len(words)

	fmt.Println(wordsCount)
}

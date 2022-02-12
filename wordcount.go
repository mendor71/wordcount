package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sentence, err := readString()

	if err != nil {
		log.Fatal(err)
	}

	words := strings.Fields(sentence)
	wordsCount := len(words)

	fmt.Println(wordsCount)
}

func readString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return input, nil
}

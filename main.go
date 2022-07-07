package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter your name, please: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userInput = strings.Trim(userInput, "\n")
	fmt.Println("hello", userInput)
}

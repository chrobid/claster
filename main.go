package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello, wanderer. This is Claster.")
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		input = strings.TrimSuffix(input, "\n")
		if strings.HasPrefix(input, "/") {
			command(input)
		}
	}
}

func command(input string) {
	switch input {
	case "/quit":
		fmt.Println("Be well.")
		os.Exit(0)
	case "/exit":
		fmt.Println("Be well.")
		os.Exit(0)
	case "/help":
		fmt.Println("This will be a big old help menu.")
	case "/connect":
		fmt.Println("This is how you'll connect to a server.")
	default:
		fmt.Println("Sorry, I don't understand.")
	}
}

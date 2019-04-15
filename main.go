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
		fmt.Print(`> `)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Panic(err)
		}
		input = strings.TrimSuffix(input, "\n")
		if strings.HasPrefix(input, `/`) {
			command(input)
		}
	}
}

func command(input string) {
	switch input {
	case `/quit`:
		os.Exit(0)
	case `/exit`:
		os.Exit(0)
	default:
		fmt.Println("Sorry, I don't understand.")
	}
}

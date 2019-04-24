package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/chrobid/libclaster"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Hello, this is Claster.")
	for {
		fmt.Print("> ")
		scanner.Scan()
		if scanner.Err() != nil {
			log.Println(scanner.Err())
			continue
		}
		if strings.HasPrefix(scanner.Text(), "/") {
			command(scanner.Text())
		}
	}
}

func command(input string) {
	scanner := bufio.NewScanner(os.Stdin)
	var client net.Conn
	var err error

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
		fmt.Print("Username: ")
		username := scanner.Scan()
		if scanner.Err() != nil {
			log.Println(scanner.Err())
			break
		}
		fmt.Print("Password: ")
		password := scanner.Scan()
		if scanner.Err() != nil {
			log.Println(scanner.Err())
			break
		}
		fmt.Println("Attempting to connect...")
		client, err = libclaster.NewClient(scanner.Text())
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Connection established")
			fmt.Println("Connected to ", client.RemoteAddr())
		}

	default:
		fmt.Println("Sorry, I don't understand.")
	}
}

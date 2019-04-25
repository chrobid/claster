package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-xmpp"
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
		scanner.Scan()
		if scanner.Err() != nil {
			log.Println(scanner.Err())
			break
		}
		jid := scanner.Text()
		if strings.Count(jid, "@") != 1 {
			fmt.Println("Invalid username")
			break
		}
		host := getHostname(jid)

		fmt.Print("Password: ")
		scanner.Scan()
		if scanner.Err() != nil {
			log.Println(scanner.Err())
			break
		}
		// *at least* clear the damn screen, since the password's visible
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		pass := scanner.Text()
		fmt.Println("Attempting to connect...")
		client, err := xmpp.NewClient(host, jid, pass, false)
		if err != nil {
			fmt.Println("Attempt failed :(")
			log.Println(err)
		} else {
			fmt.Println("Connection established")
			client = client
		}

	default:
		fmt.Println("Sorry, I don't understand.")
	}
}

func getHostname(jid string) string {
	parts := strings.Split(jid, "@")
	hostname := parts[1]
	return hostname
}

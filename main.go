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
		if strings.HasPrefix(input, `/`) {
			fmt.Println("This might one day be a command.")
		} else {
			fmt.Print(input)
		}
	}
}

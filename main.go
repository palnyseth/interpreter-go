package main

import (
	"fmt"
	"kursiv/interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	if user.Username != "nyseth" {
		fmt.Println("Who are you????")
	}
	fmt.Printf("Hello %s!\n", user.Username)
	fmt.Printf("Please feel free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

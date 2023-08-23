package main

import (
	"os"

	"github.com/wolves/greet"
)

func main() {
	greet.GreetUser(os.Stdin, os.Stdout)
}

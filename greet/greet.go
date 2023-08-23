package greet

import (
	"bufio"
	"fmt"
	"io"
)

func GreetUser(in io.Reader, out io.Writer) {
	name := ""
	fmt.Fprintln(out, "What is your name?")
	input := bufio.NewScanner(in)
	if input.Scan() {
		name = input.Text()
		fmt.Fprintf(out, "Hello, %s.\n", name)
	}
}

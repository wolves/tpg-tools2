package greet_test

import (
	"bytes"
	"testing"

	"github.com/wolves/greet"
)

func TestGreetUser_PromptsUserForNameAndRendersGreeting(t *testing.T) {
	t.Parallel()

	input := bytes.NewBufferString("John Doe")
	output := new(bytes.Buffer)
	greet.GreetUser(input, output)
	got := output.String()
	want := "What is your name?\nHello, John Doe.\n"
	if want != got {
		t.Fatalf("wanted %q but got %q", want, got)
	}
}

package match

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type matcher struct {
	input  io.Reader
	output io.Writer
	text   string
}

type option func(m *matcher) error

func WithInput(input io.Reader) option {
	return func(m *matcher) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		m.input = input
		return nil
	}
}
func WithOutput(output io.Writer) option {
	return func(m *matcher) error {
		if output == nil {
			return errors.New("nil output writer")
		}
		m.output = output
		return nil
	}
}
func WithArgsSearchText(args []string) option {
	return func(m *matcher) error {
		if len(args) < 1 {
			return nil
		}
		m.text = args[0]
		return nil
	}
}

func NewMatcher(opts ...option) (*matcher, error) {
	m := &matcher{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *matcher) MatchingLines() {
	lines := bufio.NewScanner(m.input)
	for lines.Scan() {
		//Check Matches
		if strings.Contains(lines.Text(), m.text) {
			fmt.Fprintln(m.output, lines.Text())
		}
	}
}

func Main() int {
	m, err := NewMatcher(
		WithArgsSearchText(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	m.MatchingLines()
	return 0
}

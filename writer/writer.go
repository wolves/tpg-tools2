package writer

import (
	"flag"
	"fmt"
	"os"
)

const Usage = `Usage: writefile -size SIZE_BYTES PATH

Creates the file PATH, containing SIZE_BYTES bytes, all zero.

Example: writefile -size 1000 zeroes.dat`

func Main() int {
	size := flag.Int("size", 0, "Size of file write")
	flag.Parse()

	path := flag.Arg(0)

	switch {
	case *size > 0:
		data := make([]byte, *size)
		err := WriteToFile(path, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Something went wrong. Error: %q", err)
		}
	default:
		fmt.Println(Usage)
	}
	return 0
}

func WriteToFile(path string, data []byte) error {
	if err := os.WriteFile(path, data, 0600); err != nil {
		return err
	}
	return os.Chmod(path, 0600)
}

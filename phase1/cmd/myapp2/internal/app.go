package app

import (
	"fmt"
	"io"
)

func Run(args []string, stdout io.Writer) error {
	fmt.Println("We are in myapp2")
	return nil
}

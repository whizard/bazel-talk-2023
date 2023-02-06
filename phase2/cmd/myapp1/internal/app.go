package app

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}

func Run(args []string, stdout io.Writer) error {
	fmt.Println("We are in myapp1 - starting server")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	return err
}

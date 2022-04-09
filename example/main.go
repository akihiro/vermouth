package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/akihiro/vermouth"
)

var Message string = "Hello, world!"

func init() {
	flag.StringVar(&Message, "message", Message, "response message")
}

func main() {
	vermouth.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		fmt.Fprintln(w, Message)
	})
	if err := vermouth.Execute(mux); err != nil {
		os.Exit(1)
	}
}

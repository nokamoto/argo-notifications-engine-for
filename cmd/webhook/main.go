package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	addr = "HTTP_ADDR"
)

func main() {
	v := os.Getenv(addr)
	if v == "" {
		v = ":8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading body: %v", err)
		}
		id := fmt.Sprintf("[%s %s %s]", r.RemoteAddr, r.Method, r.URL)
		log.Printf("%s body=%s", id, body)
		w.Write([]byte("ok"))
	})
	log.Fatal(http.ListenAndServe(v, nil))
}

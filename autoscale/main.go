package autoscale

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/receive", receive)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, autoscale!")
}

func receive(w http.ResponseWriter, r *http.Request) {
	time.Sleep(120 * time.Second) // 120秒待機
	fmt.Fprint(w, "Hello, autoscale!")
}

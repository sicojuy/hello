package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	addr string
)

func usageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usage:\n")
	fmt.Fprintf(w, "  %-30s %s\n", "/hello", "just say hello")
	fmt.Fprintf(w, "  %-30s %s\n", "/info", "show request line and headers")
	fmt.Fprintf(w, "  %-30s %s\n", "/redirect?target=<url>", "response 302 to target")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s: %v\n", k, strings.Join(v, ","))
	}
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "client addr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "\n")

	args := r.URL.Query()
	t := args.Get("t")
	if len(t) > 0 {
		if v, err := strconv.Atoi(t); err == nil && v > 0 {
			time.Sleep(time.Duration(v) * time.Second)
		}
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	args := r.URL.Query()
	t := args.Get("target")
	if len(t) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("target not found\n"))
		return
	}
	w.Header().Set("Location", t)
	w.WriteHeader(http.StatusFound)
}

func init() {
	flag.StringVar(&addr, "addr", ":9000", "server listen on")
}

func main() {
	flag.Parse()

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/redirect", redirectHandler)
	http.HandleFunc("/", usageHandler)

	log.Printf("listen on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

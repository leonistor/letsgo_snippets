package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// change behaviour, match exact, see below
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display snippet with id: %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost) // call before Write or WriteHeader
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create new snippet"))
}

func main() {
	// make your own, don't use DefaultServeMux with Handle() or HandleFunc()
	mux := http.NewServeMux()
	// subtree path, match when starts with pattern, ends with "/"
	mux.HandleFunc("/", home)
	// fixed path: exact match, does not end with "/"
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

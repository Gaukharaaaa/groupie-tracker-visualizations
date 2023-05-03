package internal

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func Server() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", MainPage)
	mux.HandleFunc("/artist", ArtistsPage)
	mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))

	// fmt.Printf("Starting server on %s", *addr)
	fmt.Println("Starting server at port 8000 : http://localhost:8000")
	log.Fatal(http.ListenAndServe(*addr, mux))
}

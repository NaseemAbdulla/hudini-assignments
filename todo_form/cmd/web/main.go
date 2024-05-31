package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	// Set all my enviromental varible
	os.Setenv("SNIPPETBOX_ADDR", ":9080")

	//	if env is prod
		// port of :80
		// 
	// else if env is developement
		// port of :4000

}

func main() {
	// addr := flag.String("addr", ":4000", "HTTP network address")
	// flag.Parse()

	addr := os.Getenv("SNIPPETBOX_ADDR")

	mux := http.NewServeMux()
	mux.HandleFunc("/", getTask)
	mux.HandleFunc("/addtask", addTask)
	mux.HandleFunc("/deletetask", deleteTask)
	// CSS file fetching and running
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", addr)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetMux.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			r.UserAgent(),
		)
	})
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	html := "Hello World"
	w.Write([]byte(html))
}

func main() {
	fmt.Println("start main function")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloWorld)

	fmt.Println("Server address: 127.0.0.1:8080")
	http.ListenAndServe(":8080", RequestLogger(mux))
}

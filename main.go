//Any copyright is dedicated to the Public Domain.
//https://creativecommons.org/publicdomain/zero/1.0/

// Greetd serves greetings over HTTP.
// The Greetd program is a simple web app for getting started with Stacktide.
package main

import (
	"fmt"
	"html"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "world"
		}

		fmt.Fprintf(w, "Hello, %s!", html.EscapeString(name))
	})

	slog.Info("Starting greetd server", "addr", addr)

	log.Fatal(http.ListenAndServe(addr, nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

func gopher(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	hello := `
 _          _ _                            _     _ _ 
| |        | | |                          | |   | | |
| |__   ___| | | ___   __      _____  _ __| | __| | |
| '_ \ / _ \ | |/ _ \  \ \ /\ / / _ \| '__| |/ _' | |
| | | |  __/ | | (_) |  \ V  V / (_) | |  | | (_| |_|
|_| |_|\___|_|_|\___/    \_/\_/ \___/|_|  |_|\__,_(_)`
	fromTsuru := `
 _                   
(_ _   _  _/ _   _   
/ / ()//) /_) (// (/
`
	fmt.Fprintf(w, "%s%s", color.HiCyanString(hello), color.HiGreenString(fromTsuru))
}

func main() {
	handler := http.DefaultServeMux
	handler.HandleFunc("/", gopher)

	host := "0.0.0.0"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	srv := http.Server{
		Handler: handler,
		Addr:    host + ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	color.Green("Starting server at: %s\n", srv.Addr)

	log.Fatal(srv.ListenAndServe())
}

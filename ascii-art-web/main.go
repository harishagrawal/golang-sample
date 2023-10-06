package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web-stylize/asciiHandler"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/pictures/", http.StripPrefix("/pictures/", http.FileServer(http.Dir("pictures"))))

	http.HandleFunc("/", asciiHandler.ViewHandler)

	http.HandleFunc("/ascii-art", asciiHandler.AsciiText)

	http.HandleFunc("/export", asciiHandler.ExportFile)

	fmt.Println("Running server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	pkg "groupie/pkg"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./templates/static"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./templates/img"))))

	http.HandleFunc("/", pkg.ViewHandler)

	http.HandleFunc("/artist/", pkg.ViewArtist)

	fmt.Println("Running server on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println(err)
		return
	}
}

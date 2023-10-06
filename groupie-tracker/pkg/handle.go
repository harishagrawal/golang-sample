package pkg

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var artists, err = Unmarshal()

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println("Cant Execute data to send")
		return
	}

	if r.URL.Path != "/" {
		errorHandler(w, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		parseFile, err := template.ParseFiles("./templates/index.html")
		if err != nil {

			errorHandler(w, http.StatusInternalServerError)
			return
		}

		datalist := Suggestions(artists)

		err = parseFile.Execute(w, Search{artists, datalist})

		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		parseFile, err := template.ParseFiles("./templates/index.html")
		if err != nil {

			errorHandler(w, http.StatusInternalServerError)
			return
		}

		datalist := Suggestions(artists)
		searched := SearchOut(artists, r.FormValue("search"))

		err = parseFile.Execute(w, Search{searched, datalist})

		if err != nil {
			errorHandler(w, http.StatusInternalServerError)
			return
		}
	default:
		errorHandler(w, http.StatusMethodNotAllowed)
	}
}

func ViewArtist(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url[2])
	if err != nil {
		errorHandler(w, http.StatusNotFound)
		return
	}
	if len(url) > 3 {
		errorHandler(w, http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	parseFile, err := template.ParseFiles("./templates/artist.html")
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
	if id <= 0 || id > len(artists) {
		errorHandler(w, http.StatusNotFound)
		return
	}
	err = parseFile.Execute(w, artists[id-1])
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
		return
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	type ErrorStruct struct {
		ErrorNum     int
		ErrorMessage string
	}
	errm := ErrorStruct{ErrorNum: status, ErrorMessage: http.StatusText(status)}
	fmt.Println(errm)

	temp, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, errm)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

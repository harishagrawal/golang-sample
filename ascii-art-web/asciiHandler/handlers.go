package asciiHandler

import (
	"html/template"
	"io"
	"net/http"
	"strconv"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	type ErrorStruct struct {
		ErrorNum     int
		ErrorMessage string
	}

	errm := ErrorStruct{ErrorNum: status, ErrorMessage: http.StatusText(status)}

	w.WriteHeader(status)

	temp, err := template.ParseFiles("./templates/errors.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// temp.Execute(w, errm)

	err = temp.Execute(w, errm)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	parseFile, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	err = parseFile.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func AsciiText(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	var output string

	InputName := r.FormValue("inputValue")
	if !IsPrintable(InputName) {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")

	bannerUsage := "./banners/" + banner + ".txt"

	var err error

	output, err = Mapa(bannerUsage, InputName)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// os.WriteFile("output.txt", []byte([]byte(output)), 0o644)

	parseFile, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	err = parseFile.Execute(w, output)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// http.Redirect(w, r, "/", http.StatusFound)
}

func ExportFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	var output string

	InputName := r.FormValue("inputValue")
	if !IsPrintable(InputName) {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	banner := r.FormValue("banner")

	bannerUsage := "./banners/" + banner + ".txt"

	var err error

	output, err = Mapa(bannerUsage, InputName)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(output)))
	w.Header().Set("Content-Disposition", "attachment; filename=export_file.txt")

	if _, err := io.WriteString(w, output); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

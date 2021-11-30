package index

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Holla dunya...")
}

func Else(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	notFoundMessage := []byte("Not found!")

	w.Write(notFoundMessage)
}

func Quda(w http.ResponseWriter, r *http.Request) {

}

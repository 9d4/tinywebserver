package index

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Holla dunya...")	
}

func Else(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Cari apa?")
	fmt.Fprintf(w, "Kamu sedang di %s", "asd")
}
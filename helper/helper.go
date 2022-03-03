package helper

import "net/http"

func IsEmptyString(s string) bool {
	return len(s) < 1
}

func RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user/login", http.StatusFound)
}

func RedirectToUser(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user", http.StatusFound)
}

package indexController

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/9d4/tinywebserver/helper"
	todo_service "github.com/9d4/tinywebserver/services/todo"
	"github.com/9d4/tinywebserver/view"
)

type PageData map[string]interface{}

type Todo struct {
	Content   string
	CreatedAt string
}

func Index(w http.ResponseWriter, r *http.Request) {
	todos := todo_service.GetList()

	data := map[string]interface{}{
		"Todos": todos,
	}

	view.Render(w, data, "index.html")
}

func Store(w http.ResponseWriter, r *http.Request) {
	// if the method is http.methodPost, we can use r.FormValue(key string)
	content := r.FormValue("todo")

	if helper.IsEmptyString(content) {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Payload"))
		return
	}

	_, err := todo_service.New(content)
	if err != nil {
		fmt.Fprintln(w, "error")
		fmt.Fprintln(w, err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
	// fmt.Fprintf(w, "%s", body)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, err := strconv.Atoi(vars["id"]); err == nil {
		err := todo_service.Delete(id)
		if err != nil {
			fmt.Println(err)
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func Done(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, err := strconv.Atoi(vars["id"]); err == nil {
		err := todo_service.SetStatus(id, "DONE")
		if err != nil {
			fmt.Println(err)
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func Undone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, err := strconv.Atoi(vars["id"]); err == nil {
		err := todo_service.SetStatus(id, "UNDONE")
		if err != nil {
			fmt.Println(err)
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func Else(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	notFoundMessage := []byte("Not found!")

	w.Write(notFoundMessage)
}

func LoginFake(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Preparing cookie to write")

	cookies := []http.Cookie{
		{
			Name:   "username",
			Value:  "traper",
			MaxAge: 600,
			Path:   "/",
		},
		{
			Name:   "loggedIn",
			Value:  "1",
			MaxAge: 600,
			Path:   "/",
		},
	}

	for _, cookie := range cookies {
		http.SetCookie(w, &cookie)
	}

	fmt.Println("Written!")
	helper.RedirectToUser(w, r)
}
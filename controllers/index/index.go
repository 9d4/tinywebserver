package index

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/traperwaze/tinywebserver/helper"
	todo_service "github.com/traperwaze/tinywebserver/services/todo"
	"github.com/traperwaze/tinywebserver/view"
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

package view

import (
	"errors"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/traperwaze/tinywebserver/config"
)

var partialsPath = "views/partials/"

func Partials() (Template *template.Template) {
	if _, err := os.Stat(partialsPath); err == os.ErrNotExist {
		return
	}

	filepath.WalkDir(partialsPath, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.Type().IsDir() {
			Template, _ = Template.ParseFiles(path)
		}

		return nil
	})

	return
}

// Returns real path of the view
func View(view string) string {
	viewPath := config.GetViewsPath()
	fullPath := path.Join(viewPath, view)

	if _, err := os.Stat(fullPath); errors.Is(err, os.ErrNotExist) {
		return ""
	}

	return fullPath
}

func Render(w http.ResponseWriter, data interface{}, html string) error {
	rootTmpl := Partials()

	var tmpl *template.Template

	_tmpl, err := rootTmpl.ParseFiles(View(html))
	if err != nil {
		return err
	}

	tmpl = _tmpl

	name := ""
	names := strings.Split(html, ".html")
	if len(names) > 0 {
		name = names[0]
	}

	tmpl.ExecuteTemplate(w, name, data)

	return nil
}

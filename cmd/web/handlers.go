package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/dianamatkava/snippetbox/cmd/internal/models"
)


func (app *application) home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/pages/base.html",
		"./ui/html/pages/home.html",
		"./ui/html/partials/nav.html",
	}
	file, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = file.ExecuteTemplate(w, "base", nil)  // write headers and body
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getSnippets(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Display all snippets")
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
	}
	fmt.Fprintf(w, "Display all snippets %+v", snippets)
}

func (app *application) getSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.clientError(w, r, http.StatusNotFound)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "Snippet: %+v", snippet)
}

func (app *application) getSnippetCreateForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create form")
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
    expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

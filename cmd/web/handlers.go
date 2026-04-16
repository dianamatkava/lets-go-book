package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dianamatkava/snippetbox/cmd/internal/models"
	"github.com/dianamatkava/snippetbox/cmd/internal/validator"
)


func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	templateData := SnippetsTemplate{
		Snippets: snippets, 
		CommonTemplateData: CommonTemplateData{
			CurrentYear: time.Now().Year(),
		},
	}
	app.render("home.html", templateData, w, r, http.StatusOK)
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
	templateData := SnippetTemplate{Snippet: snippet,
		CommonTemplateData: CommonTemplateData{
			CurrentYear: time.Now().Year(),
		},
	}
	app.render("view.html", templateData, w, r, http.StatusOK)
}

func (app *application) getCreateFormSnippet(w http.ResponseWriter, r *http.Request) {
	app.render("create.html", nil, w, r, http.StatusOK)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	var form SnippetCreateFormTemplate
	err := app.parseFormData(&form, r)
	if err != nil {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	form.checkField(validator.NoBlank(form.Title), "title", "This field cannot be blank")
	form.checkField(validator.MaxChar(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.checkField(validator.NoBlank(form.Content), "content", "This field cannot be blank")
	form.checkField(validator.Contains(form.Expires, []int{1, 7, 365}), "expires", "This field must equal 1, 7 or 365")
	
	if !form.Valid() {
		app.render("create.html", form, w, r, http.StatusUnprocessableEntity)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

package main

import (
	"html/template"
	"path/filepath"

	"github.com/dianamatkava/snippetbox/cmd/internal/models"
)

type Form struct {
    Errors map[string]string
}

func (form *Form) Valid() bool {
    return len(form.Errors) == 0
}

func (form *Form) checkField(ok bool, field string, err string) {
    if !ok {
        if form.Errors == nil {
            form.Errors = make(map[string]string)
        }

        if _, exists := form.Errors[field]; !exists {
            form.Errors[field] = err
        }
    }
}

type CommonTemplateData struct {
	CurrentYear int
} 

type SnippetTemplate struct {
	CommonTemplateData
	Snippet models.Snippet
}

type SnippetsTemplate struct {
	CommonTemplateData
	Snippets []models.Snippet
}

type SnippetCreateFormTemplate struct {
    Form                `form:"-"`
    CommonTemplateData  `form:"-"`
    Title       string  `form:"title"`
    Content     string  `form:"content"`
    Expires     int     `form:"expires"`
}


var functions = template.FuncMap{
    "HumanizeDate": HumanizeDate,
}


func cacheTemplates() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

    pages, err := filepath.Glob("./ui/html/pages/*.html")
    if err != nil {
        return nil, err
    }

    for _, page := range pages {
        name := filepath.Base(page)

        ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
        if err != nil {
            return nil, err
        }

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
        if err != nil {
            return nil, err
        }

        ts, err = ts.ParseFiles(page)
        if err != nil {
            return nil, err
        }

		cache[name] = ts
    }

    return cache, nil
}
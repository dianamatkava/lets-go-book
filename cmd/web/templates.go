package main

import (
	"html/template"
	"path/filepath"

	"github.com/dianamatkava/snippetbox/cmd/internal/models"
)


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
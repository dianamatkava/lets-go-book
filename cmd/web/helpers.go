package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)


func (app *application) render(key string, data any, w http.ResponseWriter, r *http.Request, status int) {
	template, ok := app.templates[key]
	if !ok {
		err := fmt.Errorf("the template does not exists")
		app.serverError(w, r, err)
		return
	}
	buffer := new(bytes.Buffer)
	err := template.ExecuteTemplate(buffer, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	buffer.WriteTo(w)
}

func GetCurrentYear() int {
	return time.Now().Year()
}
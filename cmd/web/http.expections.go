package main

import (
	"net/http"
	"runtime/debug"
)


func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri	   = r.URL 
	)
	app.logger.Error(err.Error(), "method", method, "URI", uri, "trace", string(debug.Stack()))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, r *http.Request, status int) {
	var (
		method = r.Method
		uri	   = r.URL 
	)
	app.logger.Error(http.StatusText(http.StatusBadRequest), "method", method, "URI", uri)
	http.Error(w, http.StatusText(status), status)
}
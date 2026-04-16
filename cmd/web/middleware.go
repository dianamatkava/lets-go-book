package main

import (
	"fmt"
	"net/http"
)


func (app *application) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
            ip     = r.RemoteAddr
            proto  = r.Proto
            method = r.Method
            uri    = r.URL.RequestURI()
        )

        app.logger.Info("REQ", "Method", method, "uri", uri, "ip", ip, "proto", proto)
		next.ServeHTTP(w, r) // call actual handler
		app.logger.Info("RES", "Method", r.Method, "URL", r.URL.Path)
	})
}


// ListenAndServe (TCP) -> SecureHeadersMiddleware -> ServeMux(next).ServeHTTP(w, r) -> ApiHandler.ServeHTTP(w, r) -> RES
func SecureHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
            "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

        w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "deny")
        w.Header().Set("X-XSS-Protection", "0")

        w.Header().Set("Server", "Go")
        next.ServeHTTP(w, r)
	})
}


func (app *application) panicRecoverMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        defer func() {
            pv := recover()
            if pv != nil {
                w.Header().Set("Connection", "close")
                app.serverError(w, r, fmt.Errorf("%v", pv))
            } 
        }()

        next.ServeHTTP(w, r)
    })
}
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerOptions struct {
	Bind  string
	Log   string
	Debug bool
}

func Start(opts ServerOptions) error {
	fmt.Println("Starting server on", opts.Bind)
	m := mux.NewRouter()
	m.Use(loginMiddleware)
	m.HandleFunc(`/v1/waf`, getAllWafHandler).Methods("GET")
	m.HandleFunc(`/v1/waf/{waf_tag}`, wafHandler).Methods("GET", "POST")
	return http.ListenAndServe(opts.Bind, m)
}
func wafHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)
	if r.Method == "GET" {
		getWafHandler(w, r)
	} else if r.Method == "POST" {
		postWafHandler(w, r)
	}
}

func httpError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(err.Error())); err != nil {
		fmt.Println(err)
	}
}

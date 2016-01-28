package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/pistarlabs/configs"
)

// Global variable configuration
var cfg *configs.Config
var token string

func init() {
	var err error

	cfg, err = configs.Load("./config.json")
	if err != nil {
		panic(err)
	}

	token = cfg.UString("token")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		// Invalid request method
		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusNotFound)
			return
		}

		// Invalid token key
		if r.Header.Get("X-Token-Key") != token {
			http.Error(w, "Request unauthorized", http.StatusUnauthorized)
			return
		}

		// Invalid parameter
		url := r.FormValue("url")
		if url == "" {
			http.Error(w, "Parameter required", http.StatusBadRequest)
			return
		}

		// Pinging URL
		result, err := ping(url)
		if err != nil {
			http.Error(w, fmt.Sprintf("Cannot ping to %s", url), http.StatusInternalServerError)
		}

		// Result
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result.toJSON())
	})

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":" + cfg.UString("server.port", "8000"))

}

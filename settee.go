package main

import (
	"encoding/json"
	"log"
	"net/http"

	badger "github.com/dgraph-io/badger/v3"
	chi "github.com/go-chi/chi/v5"
	middleware "github.com/go-chi/chi/v5/middleware"
)

type ServerVendor struct {
	Name string `json:"name"`
}

type ServerInfo struct {
	CouchDB  string       `json:"couchdb"`
	Version  string       `json:"version"`
	SHA      string       `json:"git_sha"`
	UUID     string       `json:"uuid"`
	Features []string     `json:"features"`
	Vendor   ServerVendor `json:"vendor"`
}

var (
	info = ServerInfo{CouchDB: "Welcome", Version: "4.0.0", SHA: "", UUID: "",
		Features: []string{"access-ready", "partitioned",
			"pluggable-storage-engines", "reshard", "scheduler"},
		Vendor: ServerVendor{Name: "Hexten"}}
)

func serve() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(info); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":5984", r)
}

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	serve()
	// Your code here…
}

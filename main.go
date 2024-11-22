package main

import (
	"db-album/db"
	"db-album/handlers"
	"net/http"
)

func main() {
	db.InitDB()

	http.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAlbums(w, r)
		case http.MethodPost:
			handlers.CreateAlbum(w, r)
		}
	})

	http.HandleFunc("/album", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetAlbum(w, r)
		case http.MethodPut:
			handlers.UpdateAlbum(w, r)
		case http.MethodDelete:
			handlers.DeleteAlbum(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}

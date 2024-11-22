package handlers

import (
	"database/sql"
	"db-album/db"
	"db-album/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, artist, price FROM album")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		albums = append(albums, album)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var album models.Album
	err := db.DB.QueryRow("SELECT id, title, artist, price FROM album WHERE id = ?", id).Scan(
		&album.ID, &album.Title, &album.Artist, &album.Price,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "Album not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		album.Title, album.Artist, album.Price,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	album.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(album)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	var album models.Album
	if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?",
		album.Title, album.Artist, album.Price, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Album with ID %s updated successfully", id)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Album with ID %s deleted successfully", id)
}

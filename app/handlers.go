package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type StatsResponse struct {
	OriginalURL string `json:"original_url"`
	Clicks      int    `json:"clicks"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "templates/index.html")
		return
	}

	id := r.URL.Path[1:]
	var originalURL string

	err := db.QueryRow(context.Background(), `UPDATE urls SET clicks = clicks + 1 WHERE id = $1 RETURNING original_url`, id).Scan(&originalURL)

	if err != nil {
		http.Redirect(w, r, "/?error=notfound", http.StatusFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	baseURL := strings.TrimSuffix(cfg.BaseURL, "/")

	if r.Method != http.MethodPost {
		http.Error(w, "POST requests only", http.StatusMethodNotAllowed)
		return
	}

	originalURL := r.FormValue("url")
	if !isValidURL(originalURL) {
		writeJSONError(w, http.StatusBadRequest, "Incorrect link")
		return
	}

	var id string
	err := db.QueryRow(context.Background(), `SELECT id FROM urls WHERE original_url = $1`, originalURL).Scan(&id)
	if err == nil {
		resp := ShortenResponse{
			ShortURL: baseURL + "/" + id,
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	newID := generateID(6)
	_, err = db.Exec(context.Background(), `INSERT INTO urls (id, original_url) VALUES ($1, $2)`, newID, originalURL)
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), 500)
		return
	}
	resp := ShortenResponse{
		ShortURL: baseURL + "/" + newID,
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)

}

func handleStats(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[7:]

	var originalURL string
	var clicks int

	err := db.QueryRow(context.Background(), `SELECT original_url, clicks FROM urls WHERE id = $1`, id).Scan(&originalURL, &clicks)
	if err != nil {
		http.Error(w, "Link not found", 404)
		return
	}

	resp := StatsResponse{
		OriginalURL: originalURL,
		Clicks:      clicks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

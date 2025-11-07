package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gorilla/mux"
)

type shortenRequest struct {
    URL string `json:"url"`
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("ok")) })
    r.HandleFunc("/metrics", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("# metrics placeholder\n")) })

    r.HandleFunc("/v1/shorten", shortenHandler).Methods("POST")
    r.HandleFunc("/{id:[a-zA-Z0-9_-]{6,}"+"}", redirectHandler).Methods("GET")

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      r,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Printf("starting server on %s", srv.Addr)
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server failed: %v", err)
    }

    // graceful shutdown example (never reached in this simple main)
    _ = srv.Shutdown(context.Background())
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
    dec := json.NewDecoder(r.Body)
    var req shortenRequest
    if err := dec.Decode(&req); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }
    // NOTE: This is a placeholder. Replace with DB logic / Redis and proper ID generation.
    id := "abc123"
    resp := map[string]string{"id": id, "short_url": "http://short.example.com/" + id}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    // placeholder: lookup original URL from DB
    original := "https://example.com"
    if id == "" {
        http.NotFound(w, r)
        return
    }
    http.Redirect(w, r, original, http.StatusFound)
}
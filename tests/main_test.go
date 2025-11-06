package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gorilla/mux"
)

func TestHealthz(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("ok")) })

    rr := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/healthz", nil)
    r.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", rr.Code)
    }
}
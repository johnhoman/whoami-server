package main

import (
    "encoding/base64"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "net/url"
    "reflect"
    "testing"
)

func Test_whoami(t *testing.T) {
    h := http.HandlerFunc(whoami)
    w := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    req.Header.Set("Authorization", "Basic "+b64Encode("John", "Pa$$word"))
    h.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d\n", w.Code)
    }
    if hi := w.Header().Get("hello"); hi != "John" {
        t.Fatalf("expected %q, got %q\n", "John", hi)
    }
    if content := w.Header().Get("Content-Type"); content != "application/json" {
        t.Fatalf("expected content-type=%q, got %q", "application/json", content)
    }
    var got map[string]any
    if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
        t.Fatalf("expected response to be serialized JSON, but got error: %s", err)
    }
    want := map[string]any{"hello": "John!"}
    if !reflect.DeepEqual(got, want) {
        t.Fatalf("expected response body %v, got %v", want, got)
    }
}

func b64Encode(username, password string) string {
    up := url.UserPassword(username, password)
    return base64.StdEncoding.EncodeToString([]byte(up.String()))
}
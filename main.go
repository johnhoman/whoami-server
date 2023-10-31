package main

import (
    "encoding/base64"
    "fmt"
    "log"
    "net/http"
    "strings"
)

const (
    Key = "whoami"
)

var (
    logger *log.Logger
)

func main() {
    logger = log.Default()
    logger.SetPrefix("whoami: ")

    err := http.ListenAndServe(":8080", http.HandlerFunc(whoami))
    if err != nil {
        logger.Fatalf("failed to start server: %s", err)
    }
}

func whoami(w http.ResponseWriter, req *http.Request) {
    log.Printf("%s %s %d %s\n", req.Method, req.URL.Path, http.StatusOK, http.StatusText(http.StatusOK))

    basic := req.Header.Get("Authorization")
    if basic == "" && !strings.HasPrefix(basic, "Basic ") {
        code := http.StatusUnauthorized
        http.Error(w, http.StatusText(code), code)
        return
    }
    basic = strings.TrimPrefix(basic, "Basic ")
    decoded, err := base64.StdEncoding.DecodeString(basic)
    if err != nil {
        code := http.StatusBadRequest
        http.Error(w, http.StatusText(code), code)
        return
    }

    who, _, _ := strings.Cut(string(decoded), ":")
    w.Header().Set("hello", who)
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"hello": "%s!"}`, who)
}
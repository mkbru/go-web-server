package main

import (
        "fmt"
        "log"
        "math/rand"
        "net/http"
        "os"
        "strconv"
        "time"
)

var serverID string

func main() {
        rand.Seed(time.Now().UTC().UnixNano())
        serverID = strconv.Itoa(rand.Intn(10) + 100)
        logger := log.New(os.Stdout, "", log.Lmicroseconds)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        logger.Printf("server %s starting on port %s ...", serverID, port)
        http.HandleFunc("/health", Middle(logger, Root))
        if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
                logger.Panicln(err)
        }
}

func Middle(l *log.Logger, f http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                l.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
                f(w, r)
        }
}

func Root(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "%s: ok", serverID)
}
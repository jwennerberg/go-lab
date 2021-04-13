package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "go api testing", r.URL.Path[1:])
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there!")
}

func main() {
  http.Handle("/metrics", promhttp.Handler())
  http.HandleFunc("/", HelloHandler)
  http.HandleFunc("/hello", HelloHandler)
  http.HandleFunc("/api", ApiHandler)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

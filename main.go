package main

import (
  "log"
  "net/http"
  "github.com/prometheus/client_golang/prometheus/promhttp"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("go api testing\n"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hi there!\n"))
}

func main() {
  http.Handle("/metrics", promhttp.Handler())
  http.HandleFunc("/hello", HelloHandler)
  http.HandleFunc("/api", ApiHandler)

  log.Fatal(http.ListenAndServe(":8000", r))
}

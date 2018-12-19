package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("."))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe("127.0.0.1:8008", nil)
}

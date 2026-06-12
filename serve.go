package main

import (
  "flag"
  "fmt"
  "log"
  "net/http"
  "os"
)

const (
  host = "127.0.0.1"
  port = "8008"
)

func main() {
  p := flag.String("p", port, "port to listen on")
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "usage: %s [-p port] <directory>\n", os.Args[0])
  }
  flag.Parse()

  if flag.NArg() != 1 {
    flag.Usage()
    os.Exit(1)
  }

  dir := flag.Arg(0)
  fs := http.FileServer(http.Dir(dir))
  http.Handle("/", fs)

  addr := host + ":" + *p
  log.Println("Listening on http://" + addr)
  log.Fatal(http.ListenAndServe(addr, nil))
}

package main

import (
    "net/http"
    "io"
    "fmt"
)

func main() {
    http.HandleFunc("/", handleXML)
    http.ListenAndServe(":8080", nil)
}

func handleXML(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("http://www.somesite.com/map.xml")
  if err != nil {
    w.WriteHeader(500)
    w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
    return
  }
  defer resp.Body.Close()

  w.Header().Add("content-type", resp.Header.Get("content-type"))
  w.Header().Set("content-length", resp.Header.Get("content-length"))
  w.Header().Add("Access-Control-Allow-Origin", "*")

  w.WriteHeader(200)
  io.Copy(w, resp.Body)
}

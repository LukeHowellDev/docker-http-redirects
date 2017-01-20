package main

import (
  "net/http"
  "strings"
)

func redirect(w http.ResponseWriter, req *http.Request) {
  if strings.Index(req.Host, "www.") != 0 {
    return
  }
  target := "//" + req.Host[4:] + req.URL.Path
  if len(req.URL.RawQuery) > 0 {
    target += "?" + req.URL.RawQuery
  }
  http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func main() {
  http.ListenAndServe(":80", http.HandlerFunc(redirect))
}

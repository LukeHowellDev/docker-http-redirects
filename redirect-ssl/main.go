package main

import (
  "net/http"
)

func redirect(w http.ResponseWriter, req *http.Request) {
  target := "https://" + req.Host + req.URL.Path
  if len(req.URL.RawQuery) > 0 {
    target += "?" + req.URL.RawQuery
  }
  http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

func main() {
  http.ListenAndServe(":80", http.HandlerFunc(redirect))
}

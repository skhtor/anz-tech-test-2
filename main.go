package main

import ("fmt"
        "net/http")

func index_handler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Hello, world!")
}

func healthcheck_handler(w http.ResponseWriter, r *http.Request)  {
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")

  payload:= `
  {
    "myapplication": [
      {
        "version": "1",
        "description": "pre-interview technical test",
        "lastcommitsha": "abc123"
      }
    ]
  }`

  fmt.Fprintf(w, payload)
}

func main()  {
  http.HandleFunc("/", index_handler)
  http.HandleFunc("/healthcheck/", healthcheck_handler)
  http.ListenAndServe(":8000", nil)
}

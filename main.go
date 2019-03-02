package main

import ("encoding/json"
        "fmt"
        "net/http")

// Structs

type healthcheckResponse struct {
  MyApplication []applicationData `json:"myapplication"`
}

type applicationData struct {
  Version       string `json:"version"`
  Description   string `json:"description"`
  LastCommitSHA string `json:"lastcommitsha"`
}


// Page handlers

func indexHandler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Hello, world!")
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request)  {

  var healthcheckResponse = healthcheckResponse {
    []applicationData { applicationData {
      Version: "1.0",
      Description: "pre-interview technical test",
      LastCommitSHA: "abc123",
    }},
  }

  response, err := json.Marshal(healthcheckResponse)
  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}


// Main

func main()  {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/healthcheck/", healthcheckHandler)
  http.ListenAndServe(":8000", nil)
}

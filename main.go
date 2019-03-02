package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

var filename = flag.String("config", "/app_metadata.json", "Location of the metadata file.")

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

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, world!")
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
  flag.Parse()
  jsonFile, err := os.Open(*filename)
  if err != nil {
  	panic(err)
  }
  defer jsonFile.Close()
  metadata, err := ioutil.ReadAll(jsonFile)

  var appData applicationData

  json.Unmarshal(metadata, &appData)

  var healthcheckResponse = healthcheckResponse {
    []applicationData { appData },
  }

  response, err := json.Marshal(healthcheckResponse)
  if err != nil {
    panic(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}


// Main

func main() {

  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/healthcheck/", healthcheckHandler)
  http.ListenAndServe(":8000", nil)
}

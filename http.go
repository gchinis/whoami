package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
  "time"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(10 * time.Second)
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
        fmt.Fprintf(w, "I'm %s\n", hostname)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}


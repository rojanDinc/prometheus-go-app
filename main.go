package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var port = envOrDefault("PORT", "8888")

func envOrDefault(key, defaultValue string) string {
	if val, found := os.LookupEnv(key); found {
		return val
	}

	return defaultValue
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello there\n")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

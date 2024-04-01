package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/ready", Ready)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, %s! You are %s years old.\n", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Family members: %s\n", data)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s, Password: %s\n", user, password)
}

func Healthz(writer http.ResponseWriter, request *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() > 25 {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Duration: %v\n", duration)))
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("OK: %v\n", duration)))
	}
}

func Ready(writer http.ResponseWriter, request *http.Request) {
	duration := time.Since(startedAt)
	if duration.Seconds() < 25 {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("Duration: %v\n", duration)))
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("OK: %v\n", duration)))
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	log.Print("starting server...")

	http.HandleFunc("/", handler)
	http.HandleFunc("/calc", calc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("setting default port to %s", port)
	}

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello WeShare!")
}

func calc(w http.ResponseWriter, r *http.Request) {
	type_operation := r.URL.Query().Get("op")
	x, _ := strconv.Atoi(r.URL.Query().Get("x"))
	y, _ := strconv.Atoi(r.URL.Query().Get("y"))

	var res int
	switch type_operation {
	case "add":
		log.Printf("addition operation, x=%d, y=%d", x, y)
		res = x + y
		break
	case "sub":
		log.Printf("substraction operation, x=%d, y=%d", x, y)
		res = x - y
		break
	case "mul":
		log.Printf("multiplication operation, x=%d, y=%d", x, y)
		res = x * y
		break
	}
	log.Printf("resultat: %d", res)
	fmt.Fprintf(w, "Resultat: %d!\n", res)
}

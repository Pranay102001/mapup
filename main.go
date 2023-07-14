package main

import (
	"log"
	"net/http"

	"github.com/your-username/project/api"
)

func main() {
	http.HandleFunc("/intersect", api.HandleIntersection)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

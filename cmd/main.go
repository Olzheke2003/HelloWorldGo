package main

import (
	"awesomeProject1/pkg/api"
	"log"
	"net/http"
)

func main() {
	api := api.New("localhost:8090", http.NewServeMux())
	api.FillEndPoints()
	log.Fatal(api.ListenAndServe())
}

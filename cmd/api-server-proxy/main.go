package main

import (
	"net/http"

	"github.com/api-server-proxy/internal/router"
)

func main() {
	router := router.NewRouter()

	

	http.ListenAndServe(":8080", router)
}

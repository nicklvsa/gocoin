package main

import (
	"fmt"
	"gocoin/routes/general"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/general/healthcheck", general.Healthcheck)

	fmt.Printf("Listening on :8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

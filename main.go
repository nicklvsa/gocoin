package main

import (
	"fmt"
	"gocoin/crypto/chain"
	bcr "gocoin/routes/blockchain"
	gcr "gocoin/routes/general"
	"gocoin/routes/sync"
	"gocoin/shared"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	blockchain, err := chain.New()
	if err != nil {
		panic(err.Error())
	}

	shared.GlobalChain = blockchain

	router.GET("/general/healthcheck", gcr.Healthcheck)
	router.POST("/transaction/new", bcr.CreateTransaction)
	router.POST("/sync/resolve", sync.ResolveTruths)

	fmt.Printf("Listening on :8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blockfrost/blockfrost-go"
)

func main() {
	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	addr := "addr1q85yx2w7ragn5sx6umgmtjpc3865s9sg59sz4rrh6f90kgwfwlzu3w8ttacqg89mkdgwshwnplj5c5n9f8dhp0h55q2q7qm63t"
	meta, err := api.Nutlink(context.TODO(), addr)
	if err != nil {
		log.Fatal(err)
	}

	tickers, err := api.Tickers(context.TODO(), addr, blockfrost.APIPagingParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Meta: \n\tAddress: %s\n\tURL: %s\n", meta.Address, meta.MetadataUrl)

	if len(tickers) == 0 {
		return
	}
	fmt.Printf("\nTickers(%d): \n\t%+v\n...", len(tickers), tickers[0])
}

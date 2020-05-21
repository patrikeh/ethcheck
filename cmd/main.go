package main

import (
	"flag"
	"log"

	"github.com/patrikeh/ethcheck"
)

func main() {
	address := flag.String("address", "", "ethereum address")
	privateKey := flag.String("private-key", "", "ethereum address")

	flag.Parse()

	if address == nil || *address == "" {
		log.Fatal("missing argument --address")
	}
	if privateKey == nil || *privateKey == "" {
		log.Fatal("missing argument --private-key")
	}

	isMatch, err := ethcheck.PrivateKeyMatchesAddress(*privateKey, *address)
	if err != nil {
		log.Fatalf("error checking private key against address - %s", err.Error())
	}

	if isMatch {
		log.Println("private key and address match")
	} else {
		log.Println("private key and address do not match")
	}
}

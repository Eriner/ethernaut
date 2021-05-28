package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

var password string

func init() {
	var ok bool
	password, ok = os.LookupEnv("ETHERNAUT_PASSWORD")
	if !ok {
		log.Fatalln("ETHERNAUT_PASSWORD is unset")
	}
}

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln("Error connecting to ETH network:", err)
	}

	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	keystoreDat, err := ioutil.ReadFile("./wallets/ethernaut.dat")
	if err != nil {
		log.Fatalln("Error reading wallet file:", err)
	}
	account, err := ks.Import(keystoreDat, password, password)
	if err != nil {
		log.Fatalln("Error importing wallet:", err)
	}

	address := account.Address.Hex()
	log.Println("address:", address)

	_ = client
}

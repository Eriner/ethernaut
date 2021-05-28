package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

const passwordEnvVar string = "ETHERNAUT_PASSWORD"

// it's either a global or an init() function. Pick your poision.
var (
	password string

	client *ethclient.Client
)

func main() {
	cmd := rootCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func rootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "ethernaut",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var ok bool
			password, ok = os.LookupEnv(passwordEnvVar)
			if !ok {
				return fmt.Errorf(passwordEnvVar + "is unset")
			}
			if client != nil {
				var err error
				client, err = ethclient.Dial("http://127.0.0.1:8545")
				if err != nil {
					return fmt.Errorf("Error connecting to the ETH network: %w", err)
				}
			}
			return nil
		},
	}
	// flags := rootCmd.PersistentFlags()
	return rootCmd
}

func getAccount() (*accounts.Account, error) {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	keystoreDat, err := ioutil.ReadFile("./wallets/ethernaut.dat")
	if err != nil {
		return nil, fmt.Errorf("Error reading wallet file: %w", err)
	}
	account, err := ks.Import(keystoreDat, password, password)
	if err != nil {
		return nil, fmt.Errorf("Error importing wallet: %w", err)
	}
	return &account, nil
}

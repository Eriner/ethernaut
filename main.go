package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// this isn't secure, I'm well aware. YOLO
const passwordEnvVar string = "ETHERNAUT_PASSWORD"

var (
	password string

	// it's either a global or an init() function. Pick your poision.
	client *ethclient.Client
)

func main() {
	cmd := rootCommand()
	cmd.AddCommand(
		createWallet(),
		mintWalletWithPrefix(),
	)
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func mintWalletWithPrefix() *cobra.Command {
	var prefix *string
	var threads *int
	cmd := &cobra.Command{
		Use: "mint-wallet-with-prefix",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if prefix == nil {
				log.Println("No prefix provided. Using \"0x6969\".")
				p := "0x6969"
				prefix = &p
			}
			if prefix != nil {
				if !strings.HasPrefix(*prefix, "0x") {
					return fmt.Errorf("invalid sting format. Should be in form of: \"0x6969\"")
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var attemptsMu sync.Mutex
			var attempts int
			fmt.Println("oh, he's tryyyin!")
			var wg sync.WaitGroup

			for i := 0; i <= *threads; i++ {
				wg.Add(1)
				go func() {
					var ks *keystore.KeyStore // re-use optimization
					for {
						attemptsMu.Lock()
						attempts++
						attemptsMu.Unlock()
						tmpDir, err := ioutil.TempDir("", "")
						if err != nil {
							log.Fatalln("unable to create temporary directory:", err.Error())
						}
						ks = keystore.NewKeyStore(tmpDir, keystore.LightScryptN, keystore.LightScryptP)
						account, err := ks.NewAccount(password)
						if err != nil {
							log.Fatalln("unable to create wallet:", err)
						}
						// log.Println("successfully created wallet:" + account.Address.Hex())
						if strings.HasPrefix(account.Address.Hex(), *prefix) {
							fmt.Println()
							log.Printf("found our address: %q after %d attempts\n", account.Address.Hex(), attempts)
							log.Printf("wallet is located in %q\n", tmpDir)
							os.Exit(0)
						}
						ks = nil // ensure old keystore and account is GC'd, otherwise there is a memory leak
						_ = os.Remove(tmpDir)
						fmt.Printf(".")
					}
				}()
			}
			wg.Wait() // this will never happen, but prevents termination.
			return nil
		},
	}
	cmd.PersistentFlags().StringVarP(prefix, "prefix", "p", "", "define a desired wallet address prefix. ex: 0x6969")
	cmd.PersistentFlags().IntVarP(threads, "threads", "t", 1, "threads")
	return cmd
}

func createWallet() *cobra.Command {
	return &cobra.Command{
		Use: "createWallet",
		RunE: func(cmd *cobra.Command, args []string) error {
			ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
			account, err := ks.NewAccount(password)
			if err != nil {
				log.Fatalln("unable to create wallet: %w", err)
			}
			log.Println("successfully created wallet:" + account.Address.Hex())
			return nil
		},
	}
}

func rootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "ethernaut",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var ok bool
			password, ok = os.LookupEnv(passwordEnvVar)
			if !ok {
				return fmt.Errorf(passwordEnvVar + " is unset")
			}
			if client == nil {
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

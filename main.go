package main

import (
	"log"

	vault "github.com/hashicorp/vault/api"
	"github.com/ragul28/hc-vault-client-basic/pkg/utils"
)

func main() {

	config := vault.DefaultConfig()

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(utils.GetEnv("VAULT_TOKEN", "DEV_TOKEN"))
}

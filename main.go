package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault/api"
	"github.com/ragul28/hc-vault-client-basic/pkg/utils"
)

func main() {

	config := vault.DefaultConfig()
	config.Address = utils.GetEnv("VAULT_ADDR", "http://localhost:8200")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(utils.GetEnv("VAULT_TOKEN", "DEV_TOKEN"))

	secretData := map[string]interface{}{
		"password": "TopSecret",
	}

	ctx := context.Background()

	// Write a secret
	_, err = client.KVv2("secret").Put(ctx, "my-secret-password", secretData)
	if err != nil {
		log.Fatalf("unable to write secret: %v", err)
	}

	log.Println("Secret written successfully.")
}

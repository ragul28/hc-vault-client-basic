package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault/api"
	"github.com/ragul28/hc-vault-client-basic/pkg/utils"
	"github.com/ragul28/hc-vault-client-basic/pkg/vaultSecret"
)

func main() {

	config := vault.DefaultConfig()
	config.Address = utils.GetEnv("VAULT_ADDR", "http://localhost:8200")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(utils.GetEnv("VAULT_TOKEN", "DEV_TOKEN"))

	secretPath := "my-secret-password"
	secretData := map[string]interface{}{
		"password": "TopSecret",
	}

	ctx := context.Background()

	// Write a secret
	vaultSecret.VaultWriteSecret(ctx, client, secretPath, secretData)

	// Read a secret
	vaultSecret.VaultReadSecret(ctx, client, secretPath, "password")

	// Get secret versions
	vaultSecret.VaultGetSecretVersions(ctx, client, secretPath, "password")

	// Delete Secret
	vaultSecret.VaultDeleteSecret(ctx, client, secretPath)
}

package main

import (
	"context"
	"log"

	vault "github.com/hashicorp/vault/api"
	"github.com/ragul28/hc-vault-client-basic/internal/config"
	"github.com/ragul28/hc-vault-client-basic/pkg/vaultSecret"
)

func main() {

	env := config.GetEnvVar()

	config := vault.DefaultConfig()
	config.Address = env.VaultAddr

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize a Vault client: %v", err)
	}

	client.SetToken(env.VaultToken)

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

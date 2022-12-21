package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Read a secret
	secret, err := client.KVv2("secret").Get(ctx, "my-secret-password")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		log.Fatalf("vaule type assetion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	log.Printf("Super secret password [%s] was retrieved.\n", value)

	// Get secret versions
	vaultGetSecretVersions(ctx, client, "my-secret-password", "password")
}

func vaultGetSecretVersions(ctx context.Context, client *vault.Client, secretPath, secretKey string) {
	versions, err := client.KVv2("secret").GetVersionsAsList(ctx, secretPath)
	if err != nil {
		log.Fatalf(
			"Unable to retrieve all versions of the super secret password from the vault. Reason: %v",
			err,
		)
	}

	fmt.Printf("Version\t Created at\t\t\t Deleted at\t Destroyed\t Value\n")

	for _, version := range versions {
		deleted := "Not deleted"
		if !version.DeletionTime.IsZero() {
			deleted = version.DeletionTime.Format(time.UnixDate)
		}

		secret, err := client.KVv2("secret").GetVersion(ctx, secretPath, version.Version)
		if err != nil {
			log.Fatalf(
				"Unable to retrieve version %d of the super secret password from the vault. Reason: %v",
				version.Version,
				err,
			)
		}
		value, ok := secret.Data[secretKey].(string)

		if ok {
			fmt.Printf("%d\t %s\t %s\t %t\t\t %s\n",
				version.Version,
				version.CreatedTime.Format(time.UnixDate),
				deleted,
				version.Destroyed,
				value,
			)
		}
	}
}

func vaultDeleteSecret(ctx context.Context, client *vault.Client, secretPath string) {
	// Delete: Deletes latest version
	// DeleteMetadata: Deletes all version & path
	// DeleteVersion: Deletes specfied versions of secret
	err := client.KVv2("secret").Delete(ctx, secretPath)
	if err != nil {
		log.Fatalf("Unable to delete the latest version of the secret: %v", err)
	}
	log.Println("Deleted latest version of Secret")
}

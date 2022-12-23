package config

import "github.com/ragul28/hc-vault-client-basic/pkg/utils"

type VaultVar struct {
	VaultAddr  string
	VaultToken string
}

func GetEnvVar() *VaultVar {
	var v VaultVar

	v.VaultAddr = utils.GetEnv("VAULT_ADDR", "http://localhost:8200")
	v.VaultToken = utils.GetEnv("VAULT_TOKEN", "DEV_TOKEN")

	return &v
}

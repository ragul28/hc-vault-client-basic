# Hashicorp Vault Client Basic

## Prerequisites

* Golang 1.19+
* Docker 20+ (compose extension)
* Vault

## Running

* Clone project & fetch go dependency
    ```sh
    go mod download
    ```

* Run the Vault using docker 
    ```sh
    docker compose pull
    docker compose up -d
    ```
# Vault Store

[![Tests Status](https://github.com/gouniverse/vaultstore/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/gouniverse/vaultstore/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/vaultstore)](https://goreportcard.com/report/github.com/gouniverse/vaultstore)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/vaultstore)](https://pkg.go.dev/github.com/gouniverse/vaultstore)

Vault - a secure value storage (data-at-rest) implementation for Go.

## Installation
```
go get -u github.com/gouniverse/valuestore
```

## Table Schema ##

The following schema is used for the database.

| vault       |                  |
|-------------|------------------|
| id          | String, UniqueId |
| vault_value | Long Text        |
| created_At  | DateTime         |
| updated_at  | DateTime         |
| deleted_at  | DateTime         |

## Setup

```golang
vault, err := NewStore(NewStoreOptions{
	VaultTableName:     "my_vault",
	DB:                 databaseInstance,
	AutomigrateEnabled: true,
})

```

## Usage

- Stores a value to the vault and return the ID

```golang
id, err = vault.ValueStore("my_value", "my_password")

if err != nil {
  t.Fatalf("Store Failure: [%v]", err.Error())
}
```

- Retrieve a value from vault by ID

```golang
val, err := vault.ValueRetrieve(id, "test_pass")

if err != nil {
  t.Fatalf("Retrieve Failure: [%v]", err.Error())
}

- Delete a value from vault by ID

```golang
err := vault.ValueDelete(id)
if err != nil {
  t.Fatalf("Delete Failed: " + errDelete.Error())
}
```

## Changelog

2022.12.26 - Updated ID to Human Friendly UUID

2021.12.12 - Added tests badge

2021.12.28 - Fixed bug where DB scanner was returning empty values

2021.12.28 - Removed GORM dependency and moved to the standard library

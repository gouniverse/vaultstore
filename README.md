# Vault Store <a href="https://gitpod.io/#https://github.com/gouniverse/vaultstore" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/vaultstore/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/gouniverse/vaultstore/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/vaultstore)](https://goreportcard.com/report/github.com/gouniverse/vaultstore)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/vaultstore)](https://pkg.go.dev/github.com/gouniverse/vaultstore)

Vault - a secure value storage (data-at-rest) implementation for Go.

## License

This project is licensed under the GNU General Public License version 3 (GPL-3.0). You can find a copy of the license at https://www.gnu.org/licenses/gpl-3.0.en.html

For commercial use, please use my [contact page](https://lesichkov.co.uk/contact) to obtain a commercial license.

## Installation
```
go get -u github.com/gouniverse/valuestore
```

## Table Schema ##

The following schema is used for the database.

| vault       |                  |
|-------------|------------------|
| id          | String, UniqueId |
| vault_token | Long Text, Unique|
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

- Stores a value to the vault and returns a token

```golang
token, err = vault.TokenCreate("my_value", "my_password", 20)

if err != nil {
  print("Create Failure: " + err.Error())
}
```

- Check a token exists

```golang
exists, err := vault.TokenExists(token)

if err != nil {
  print("Delete Failed: " + errDelete.Error())
}

if (!exists) {
    print("token does not exist")
}
```

- Retrieve a value from the vault by its token

```golang
val, err := vault.TokenRead(token, "test_pass")

if err != nil {
  print("Read failed:" + err.Error())
}
```

- Delete a value from the vault by its token

```golang
err := vault.TokenDelete(token)
if err != nil {
  print("Delete failed: " + err.Error())
}
```

## Changelog

2024.08.11 - Added token support, to allow (tokenization). More info at: https://en.wikipedia.org/wiki/Tokenization_(data_security)

2022.12.26 - Updated ID to Human Friendly UUID

2021.12.12 - Added tests badge

2021.12.28 - Fixed bug where DB scanner was returning empty values

2021.12.28 - Removed GORM dependency and moved to the standard library

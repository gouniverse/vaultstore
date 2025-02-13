# Vault Store

[![Tests Status](https://github.com/gouniverse/vaultstore/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/gouniverse/vaultstore/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/vaultstore)](https://goreportcard.com/report/github.com/gouniverse/vaultstore)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/vaultstore)](https://pkg.go.dev/github.com/gouniverse/vaultstore)

Vault - a secure value storage (data-at-rest) implementation for Go.

## License

This project is licensed under the GNU Affero General Public License v3.0 (AGPL-3.0). You can find a copy of the license at [https://www.gnu.org/licenses/agpl-3.0.en.html](https://www.gnu.org/licenses/agpl-3.0.txt)

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
| created_at  | DateTime         |
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


## 🌏  Development in the Cloud 

Click any of the buttons below to start a new development environment to demo or contribute to the codebase without having to install anything on your machine:

[![Open in VS Code](https://img.shields.io/badge/Open%20in-VS%20Code-blue?logo=visualstudiocode)](https://vscode.dev/github/gouniverse/vaultstore)
[![Open in Glitch](https://img.shields.io/badge/Open%20in-Glitch-blue?logo=glitch)](https://glitch.com/edit/#!/import/github/gouniverse/vaultstore)
[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/gouniverse/vaultstore)
[![Open in StackBlitz](https://developer.stackblitz.com/img/open_in_stackblitz.svg)](https://stackblitz.com/github/gouniverse/vaultstore)
[![Edit in Codesandbox](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/s/github/gouniverse/vaultstore)
[![Open in Repl.it](https://replit.com/badge/github/withastro/astro)](https://replit.com/github/gouniverse/vaultstore)
[![Open in Codeanywhere](https://codeanywhere.com/img/open-in-codeanywhere-btn.svg)](https://app.codeanywhere.com/#https://github.com/gouniverse/vaultstore)
[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/gouniverse/vaultstore)

## Changelog

2024.08.11 - Added token support, to allow (tokenization). More info at: https://en.wikipedia.org/wiki/Tokenization_(data_security)

2022.12.26 - Updated ID to Human Friendly UUID

2021.12.12 - Added tests badge

2021.12.28 - Fixed bug where DB scanner was returning empty values

2021.12.28 - Removed GORM dependency and moved to the standard library

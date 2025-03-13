# Vault Store

[![Tests Status](https://github.com/gouniverse/vaultstore/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/gouniverse/vaultstore/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/vaultstore)](https://goreportcard.com/report/github.com/gouniverse/vaultstore)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/vaultstore)](https://pkg.go.dev/github.com/gouniverse/vaultstore)

Vault - a secure value storage (data-at-rest) implementation for Go.

## Scope

VaultStore is specifically designed as a data store component for securely storing and retrieving secrets. It is **not** an API or a complete secrets management system. Features such as user management, access control, and API endpoints are intentionally beyond the scope of this project.

VaultStore is meant to be integrated into your application as a library, providing the data storage layer for your secrets management needs. The application using VaultStore is responsible for implementing any additional layers such as API endpoints, user management, or access control if needed.

## Documentation

- [Overview](/docs/overview.md) - General overview of the VaultStore library
- [Usage Guide](/docs/usage_guide.md) - Examples of how to use VaultStore
- [Technical Reference](/docs/technical_reference.md) - Detailed technical information
- [Query Interface](/docs/query_interface.md) - Documentation for the flexible query interface
- [Data Stores](/docs/data_stores.md) - Information about the data store implementation

## Features

- Secure storage of sensitive data
- Token-based access to secrets
- Password protection for stored values
- Flexible query interface for retrieving records
- Soft delete functionality for data recovery
- Support for multiple database backends

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
| soft_deleted_at | DateTime     |

## Setup

```golang
vault, err := NewStore(NewStoreOptions{
	VaultTableName:     "my_vault",
	DB:                 databaseInstance,
	AutomigrateEnabled: true,
})

```

## Usage

Here are some basic examples of using VaultStore. For comprehensive documentation, see the [Usage Guide](/docs/usage_guide.md).

```golang
// Create a token
token, err := vault.TokenCreate("my_value", "my_password", 20)

// Check if a token exists
exists, err := vault.TokenExists(token)

// Read a value using a token
value, err := vault.TokenRead(token, "my_password")

// Update a token's value
err := vault.TokenUpdate(token, "new_value", "my_password")

// Delete a token
err := vault.TokenDelete(token)

// Soft delete a token
err := vault.TokenSoftDelete(token)
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

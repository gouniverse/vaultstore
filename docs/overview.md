# VaultStore Documentation

## Overview

VaultStore is a Go implementation for securely storing secrets in a database. It provides a simple interface for creating, reading, updating, and deleting secrets using tokens.

## Scope

VaultStore is specifically designed as a data store component for securely storing and retrieving secrets. It is **not** an API or a complete secrets management system. Features such as user management, access control, and API endpoints are intentionally beyond the scope of this project.

VaultStore is meant to be integrated into your application as a library, providing the data storage layer for your secrets management needs. The application using VaultStore is responsible for implementing any additional layers such as API endpoints, user management, or access control if needed.

## Core Concepts

### Records

A `Record` is the fundamental data entity in VaultStore. It represents a stored secret and contains:

- **ID**: A unique identifier for the record
- **Token**: A string used to access the secret
- **Value**: The encrypted secret value
- **Metadata**: Creation, update, and deletion timestamps

Records are implemented as Go structs that embed `dataobject.DataObject` for data storage.

### Tokens

A token is a string identifier used to access and manipulate secrets. The VaultStore interface provides methods to interact with secrets using tokens:

- `TokenCreate`: Creates a new token with an associated secret value
- `TokenRead`: Retrieves the secret value associated with a token
- `TokenUpdate`: Updates the secret value associated with a token
- `TokenDelete`: Deletes a token and its associated secret
- `TokenExists`: Checks if a token exists

Tokens abstract away the direct manipulation of Record objects, providing a simpler interface for secret management.

## Store Interface

The `StoreInterface` defines the contract for interacting with the vault:

```go
type StoreInterface interface {
    AutoMigrate() error
    EnableDebug(debug bool)

    RecordCreate(ctx context.Context, record Record) error
    RecordFindByID(ctx context.Context, recordID string) (*Record, error)
    RecordFindByToken(ctx context.Context, token string) (*Record, error)
    RecordList(ctx context.Context, options RecordQueryOptions) ([]Record, error)
    RecordUpdate(ctx context.Context, record Record) error
    RecordDeleteByID(ctx context.Context, recordID string) error

    TokenCreate(ctx context.Context, value string, password string, tokenLength int) (token string, err error)
    TokenCreateCustom(ctx context.Context, token string, value string, password string) (err error)
    TokenDelete(ctx context.Context, token string) error
    TokenExists(ctx context.Context, token string) (bool, error)
    TokenRead(ctx context.Context, token string, password string) (string, error)
    TokenUpdate(ctx context.Context, token string, value string, password string) error
    TokensRead(ctx context.Context, tokens []string, password string) (map[string]string, error)
}
```

## Encryption

VaultStore encrypts secret values before storing them in the database. The encryption is password-based, meaning you need the correct password to decrypt and access the secret value.

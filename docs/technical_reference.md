# Technical Reference

This document provides a detailed technical reference for the VaultStore implementation.

## Database Schema

VaultStore uses a single table to store records. The table name is configurable when creating a new store.

The table has the following columns:

- `id`: Primary key, a unique identifier for the record
- `vault_token`: The token used to access the secret
- `vault_value`: The encrypted secret value
- `created_at`: Timestamp when the record was created
- `updated_at`: Timestamp when the record was last updated
- `deleted_at`: Timestamp when the record was deleted (for soft deletes)

## Record Structure

The `Record` struct is defined as follows:

```go
type Record struct {
    dataobject.DataObject
}
```

It embeds `dataobject.DataObject` for storing data. The struct provides methods for accessing and modifying the record's fields:

- `ID()` / `SetID(id string)`: Get/set the record's ID
- `Token()` / `SetToken(token string)`: Get/set the record's token
- `Value()` / `SetValue(value string)`: Get/set the record's encrypted value
- `CreatedAt()` / `SetCreatedAt(createdAt string)`: Get/set the record's creation timestamp
- `UpdatedAt()` / `SetUpdatedAt(updatedAt string)`: Get/set the record's update timestamp
- `DeletedAt()` / `SetDeletedAt(deletedAt string)`: Get/set the record's deletion timestamp

## Store Implementation

The `Store` struct implements the `StoreInterface` and provides methods for interacting with the database.

### Initialization

The store is initialized with the `NewStore` function, which takes a `NewStoreOptions` struct:

```go
type NewStoreOptions struct {
    VaultTableName     string
    DB                 *sql.DB
    DbDriverName       string
    AutomigrateEnabled bool
    DebugEnabled       bool
}
```

### Auto-Migration

If `AutomigrateEnabled` is set to `true`, the store will automatically create the necessary table in the database if it doesn't exist.

### Encryption and Decryption

VaultStore uses password-based encryption to protect secret values. The encryption and decryption functions are defined in `encdec.go`.

The `encode` function encrypts a value with a password:

```go
func encode(value string, password string) string
```

The `decode` function decrypts a value with a password:

```go
func decode(value string, password string) (string, error)
```

### Token Generation

Tokens are generated using the `GenerateToken` function, which creates a random string of the specified length:

```go
func GenerateToken(length int) string
```

## Error Handling

VaultStore returns errors for various scenarios:

- Database errors
- Token not found
- Decryption errors (e.g., wrong password)
- Invalid parameters

It's important to check for errors when calling VaultStore methods.

## Thread Safety

VaultStore is designed to be thread-safe. It uses database transactions to ensure data consistency when multiple goroutines access the store simultaneously.

## Performance Considerations

VaultStore is designed for secure storage of secrets, not for high-performance data access. The encryption and decryption operations can be CPU-intensive, especially for large values.

For better performance, consider:

- Caching frequently accessed secrets
- Using shorter tokens (but not too short to compromise security)
- Optimizing database access (e.g., using indexes)

## Security Considerations

- Store passwords securely; they are used to encrypt and decrypt secrets
- Use HTTPS when transmitting tokens and passwords
- Implement proper access control to restrict who can create, read, update, and delete secrets
- Consider using a secure key management system for storing encryption keys

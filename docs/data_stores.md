# Data Stores

The data stores in this vault implementation are located within the `vaultstore` package. These stores are responsible for managing the persistence of secrets in a database.

## Entities

Data entities are represented as `Record` objects. These objects encapsulate the secret data and associated metadata.

## Stores

The `Store` type provides the data access layer for interacting with the database. It offers methods for creating, reading, updating, and deleting records.

## Accessing Stores

The stores are accessed via public interfaces, ensuring a clear separation of concerns and allowing for potential future implementations or modifications without affecting the rest of the system.

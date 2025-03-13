# Password Hashing and Rekeying Functionality

*Proposal Date: March 12, 2025*

## Status: Needs Refinement
- Do we need an extra column? How much is the overhead?
- Why not separate meta table? How much is the overhead?
- Security?

## Overview

This proposal suggests implementing password hashing and rekeying functionality in VaultStore to enhance security and enable bulk operations on secrets that share the same password.

## Current Implementation

Currently, VaultStore uses passwords directly for encryption and decryption of secret values. There is no mechanism to:
1. Verify a password is correct before attempting decryption
2. Identify all secrets encrypted with the same password
3. Efficiently rekey multiple secrets when a password needs to be changed

## Proposed Changes

1. **Password Hash Storage**: Add a `password_hash` field to the Record structure to store a secure hash of the encryption password.

2. **Password Verification**: Implement a method to verify a provided password against the stored hash before attempting decryption.

3. **Record Identification by Password**: Create functionality to identify all records that share the same password hash.

4. **Bulk Rekeying**: Implement a method to rekey (re-encrypt) all secrets that share the same password with a new password.

5. **Password Strength Validation**: Add optional password strength validation during secret creation or update.

## Implementation Details

The implementation would require:

- Modifying the database schema to add the `password_hash` column
- Updating the Record struct to include the password hash
- Implementing secure password hashing using Argon2id or bcrypt
- Creating new methods for password verification and rekeying

Example new methods:

```go
// Verify a password matches the stored hash
VerifyPassword(ctx context.Context, token string, password string) (bool, error)

// Find all records with a matching password hash
FindRecordsByPasswordHash(ctx context.Context, password string) ([]Record, error)

// Rekey a single record with a new password
RekeyRecord(ctx context.Context, recordID string, oldPassword string, newPassword string) error

// Bulk rekey all records with a matching password
BulkRekey(ctx context.Context, oldPassword string, newPassword string) (count int, error error)
```

## Benefits

- **Enhanced Security**: Allows password verification without attempting decryption, preventing timing attacks
- **Operational Efficiency**: Enables bulk operations on secrets sharing the same password
- **Key Rotation**: Facilitates regular key rotation as a security best practice
- **Breach Response**: Provides a mechanism to quickly update all affected secrets if a password is compromised

## Risks and Mitigations

- **Storage of Password Hashes**: Even though hashes are one-way, they represent additional sensitive data. Mitigation: Use strong hashing algorithms with appropriate work factors.
- **Performance Impact**: Password hashing adds computational overhead. Mitigation: Cache verification results where appropriate.
- **Migration Complexity**: Existing records need to be updated with password hashes. Mitigation: Provide a migration tool and update records incrementally.

## Effort Estimation

- Development: 1-2 weeks
- Testing: 3-5 days
- Documentation: 1-2 days
- Migration Tools: 2-3 days

## Conclusion

Implementing password hashing and rekeying functionality would significantly enhance the security and manageability of VaultStore, particularly in scenarios where multiple secrets share the same password or where regular key rotation is required.

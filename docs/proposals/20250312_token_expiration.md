# Token Expiration Functionality

## Status: Requires Refinement
- Does it need another column? Why not separate meta table?
- Who will run the process? The client? The library?

## Overview

This proposal suggests adding expiration functionality to tokens in VaultStore, allowing secrets to automatically expire after a specified time period.

## Current Implementation

Currently, VaultStore tokens do not have an expiration mechanism. Once created, tokens remain valid indefinitely unless explicitly deleted.

## Proposed Changes

1. **Add Expiration Field**: Add an `expires_at` field to the Record structure to store the expiration timestamp.

2. **Expiration Parameter**: Extend the `TokenCreate` and `TokenCreateCustom` methods to accept an optional expiration duration parameter.

3. **Automatic Expiration Check**: Modify token retrieval methods to check if a token has expired before returning its value.

4. **Expired Token Cleanup**: Implement a background process or method to clean up expired tokens.

5. **Token Renewal**: Add a method to extend the expiration time of an existing token.

## Implementation Details

The implementation would require:

- Modifying the database schema to add the `expires_at` column
- Updating the Record struct and its methods
- Extending the StoreInterface with new methods
- Implementing the background cleanup process

Example API changes:

```go
// New method signatures
TokenCreate(ctx context.Context, value string, password string, tokenLength int, expiresIn time.Duration) (token string, err error)
TokenCreateCustom(ctx context.Context, token string, value string, password string, expiresIn time.Duration) (err error)
TokenRenew(ctx context.Context, token string, password string, expiresIn time.Duration) (err error)
CleanupExpiredTokens(ctx context.Context) (count int, err error)
```

## Benefits

- **Security Enhancement**: Automatically invalidates tokens after a certain period, reducing the risk of unauthorized access
- **Resource Management**: Helps keep the database clean by removing unused tokens
- **Compliance**: Supports compliance requirements that mandate credential rotation
- **Use Case Support**: Enables temporary access scenarios (e.g., one-time passwords, temporary API keys)

## Risks and Mitigations

- **Breaking Changes**: API changes might break existing code. Mitigation: Provide backward compatibility by making expiration optional.
- **Performance Impact**: Additional checks during token retrieval. Mitigation: Optimize database queries with proper indexing.
- **Clock Synchronization**: Reliance on system clock. Mitigation: Use server time consistently and document this dependency.

## Effort Estimation

- Development: 1-2 weeks
- Testing: 3-5 days
- Documentation: 1-2 days

## Conclusion

Adding token expiration functionality would significantly enhance the security and usability of VaultStore, making it more suitable for a wider range of use cases and security requirements.

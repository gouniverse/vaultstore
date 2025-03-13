# Secret Versioning System

## Status: Needs Refinement
- What about new passwords for new versions?
- How to handle password changes?

## Overview

This proposal suggests implementing a versioning system for secrets in VaultStore, allowing users to track changes, roll back to previous versions, and maintain a history of secret values.

## Current Implementation

Currently, VaultStore only maintains the current value of a secret. When a secret is updated, the previous value is permanently lost with no way to recover or track historical changes.

## Proposed Changes

1. **Version History**: Modify the database schema to store multiple versions of each secret.

2. **Version Metadata**: Add metadata for each version, including creation timestamp, version number, and optional comments.

3. **API Extensions**: Extend the API to support:
   - Retrieving a specific version of a secret
   - Listing all versions of a secret
   - Rolling back to a previous version
   - Adding comments when updating a secret

4. **Pruning Policies**: Implement configurable policies for pruning old versions to manage storage growth.

## Implementation Details

The implementation would require:

- Creating a new table for version history or modifying the existing table structure
- Updating the Record struct to support version information
- Extending the StoreInterface with new version-related methods
- Implementing pruning mechanisms

Example new methods:

```go
// New method signatures
TokenVersionRead(ctx context.Context, token string, password string, version int) (value string, err error)
TokenVersionList(ctx context.Context, token string, password string) ([]VersionInfo, error)
TokenVersionRollback(ctx context.Context, token string, password string, version int) error
TokenUpdateWithComment(ctx context.Context, token string, value string, password string, comment string) error
TokenVersionPrune(ctx context.Context, token string, password string, keepLatest int) error
```

## Benefits

- **Change Tracking**: Provides visibility into how and when secrets have changed
- **Recovery**: Allows recovery from accidental or malicious updates
- **Compliance**: Supports audit requirements for tracking changes to sensitive data
- **Operational Safety**: Reduces risk when updating secrets by preserving previous versions

## Risks and Mitigations

- **Storage Growth**: Version history can significantly increase storage requirements. Mitigation: Implement configurable pruning policies.
- **Performance Impact**: Retrieving and managing versions adds complexity. Mitigation: Optimize queries and consider caching strategies.
- **API Complexity**: More complex API may be harder to use. Mitigation: Provide clear documentation and helper functions.

## Effort Estimation

- Development: 2-3 weeks
- Testing: 1-2 weeks
- Documentation: 3-5 days

## Conclusion

Implementing a secret versioning system would significantly enhance the robustness and utility of VaultStore, making it more suitable for enterprise environments where change tracking and recovery capabilities are essential.

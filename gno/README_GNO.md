# ZK Package for Gno.land

This is a Gno.land-compatible version of the ZK package, structured as a **Pure Package** (`/p/`) for use in Gno.land smart contracts.

## Package Path

```
gno.land/p/low88/zk
```

## Key Differences from Standard Go

1. **No Random Number Generation**: Blockchains require deterministic execution. All randomness (nonces, blinding factors) must be generated **off-chain** and passed as parameters.

2. **Pure Package Structure**: This is a stateless library package (`/p/`) that can be imported by Realm packages (`/r/`) for stateful smart contracts.

3. **Gno Standard Library**: Uses `std` package instead of standard Go libraries where applicable.

4. **File Extension**: Uses `.gno` extension instead of `.go`.

## Usage in Gno.land

### As a Pure Package

Import in your Realm package:

```gno
import (
    "gno.land/p/low88/zk"
)
```

### Example: Verifying Proofs On-Chain

```gno
package main

import (
    "std"
    "gno.land/p/low88/zk"
)

// Realm package for storing verified proofs
type ZKRealm struct {
    verifiedProofs map[string]bool
}

func NewZKRealm() *ZKRealm {
    return &ZKRealm{
        verifiedProofs: make(map[string]bool),
    }
}

// VerifyAndStore verifies a ZK proof and stores it if valid
func (zr *ZKRealm) VerifyAndStore(proof *zk.Proof, statement *zk.Statement) bool {
    valid := zk.VerifyProof(proof, statement)
    if valid {
        // Store proof hash as verified
        proofHash := std.Hash(proof.Commitment)
        zr.verifiedProofs[string(proofHash)] = true
        return true
    }
    return false
}

// IsVerified checks if a proof has been verified
func (zr *ZKRealm) IsVerified(commitment []byte) bool {
    proofHash := std.Hash(commitment)
    return zr.verifiedProofs[string(proofHash)]
}
```

### Example: Merkle Tree Verification

```gno
import "gno.land/p/low88/zk"

// Verify membership in a Merkle tree
func VerifyMembership(leaf []byte, proof *zk.MerkleProof) bool {
    return zk.VerifyMerkleProof(proof)
}
```

## Workflow

1. **Off-Chain (Go)**: Generate proofs, commitments, and blinding factors using the standard Go package (`zk/` directory).

2. **On-Chain (Gno)**: Verify proofs using the Gno.land package (`gno/` directory).

## Deployment

To deploy this package to Gno.land:

```bash
# Make sure you have gnokey set up with your account
gnokey maketx addpkg \
    --pkgpath "gno.land/p/low88/zk" \
    --pkgdir "./gno" \
    --gas-fee "1000000ugnot" \
    --gas-wanted "2000000" \
    --broadcast \
    --chainid "dev"
```

## Important Notes

- **Proof Generation**: Must happen off-chain using the standard Go package
- **Proof Verification**: Happens on-chain using this Gno package
- **Gas Costs**: ZK verification can be expensive. Optimize for gas efficiency.
- **Storage**: Consider storage costs when storing proofs on-chain

## Integration with Zcash/Privacy Tech

Given your work with Zcash and privacy tech, this package can be used to:
- Verify zk-SNARK proofs on Gno.land
- Implement privacy-preserving smart contracts
- Bridge Zcash privacy features to Gno.land ecosystem
- Create shielded transaction verification

## Testing Locally

Use `gnodev` for local development:

```bash
gnodev test ./gno
```

## References

- [Gno.land Documentation](https://docs.gno.land)
- [Gno Package Structure](https://docs.gno.land/builders/anatomy-of-a-gno-package/)
- [Deploying Packages](https://docs.gno.land/builders/deploy-packages/)

# Setting Up ZK Package for Gno.land

## Overview

This package provides both:
1. **Standard Go package** (`zk/`) - For off-chain proof generation
2. **Gno.land package** (`gno/`) - For on-chain proof verification

## Package Structure

```
zk-go/
├── zk/              # Standard Go package (off-chain)
│   ├── proof.go
│   ├── commitment.go
│   ├── merkle.go
│   └── zk_test.go
├── gno/              # Gno.land package (on-chain)
│   ├── verifier.gno
│   ├── commitment.gno
│   ├── merkle.gno
│   └── example_realm.gno
├── go.mod            # Go module
├── gnomod.toml       # Gno module config
└── README_GNO.md     # Gno.land documentation
```

## Installation

### 1. Install Gno.land Tools

```bash
# Install gnodev (development tool)
go install github.com/gnolang/gno/gnodev@latest

# Install gnokey (key management)
go install github.com/gnolang/gno/gnokey@latest
```

### 2. Set Up Your Account

```bash
# Create a new key
gnokey add mykey

# Get your address
gnokey list
```

### 3. Update Package Path

Edit `gnomod.toml` and update the `pkgpath` with your Gno.land username:

```toml
[module]
pkgpath = "gno.land/p/YOUR_USERNAME/zk"
```

## Development Workflow

### Local Testing

```bash
# Test the Gno package locally
gnodev test ./gno

# Run the example realm
gnodev run ./gno/example_realm.gno
```

### Deploy to Devnet

```bash
gnokey maketx addpkg \
    --pkgpath "gno.land/p/low88/zk" \
    --pkgdir "./gno" \
    --gas-fee "1000000ugnot" \
    --gas-wanted "2000000" \
    --broadcast \
    --chainid "dev"
```

### Deploy to Mainnet

```bash
gnokey maketx addpkg \
    --pkgpath "gno.land/p/low88/zk" \
    --pkgdir "./gno" \
    --gas-fee "1000000ugnot" \
    --gas-wanted "2000000" \
    --broadcast \
    --chainid "mainnet"
```

## Usage Pattern

### Off-Chain (Go) - Generate Proofs

```go
package main

import "github.com/Lowo88/ZK-GO-g/zk"

func main() {
    // Generate proof off-chain
    statement := &zk.Statement{
        PublicValue: []byte("public statement"),
    }
    witness := &zk.Witness{
        Secret: []byte("secret"),
    }
    
    prover := zk.NewProver(statement, witness)
    proof, _ := prover.GenerateProof()
    
    // Send proof to blockchain for verification
    // (via transaction or call)
}
```

### On-Chain (Gno) - Verify Proofs

```gno
import "gno.land/p/low88/zk"

// In your realm package
func VerifyProof(proof *zk.Proof, statement *zk.Statement) bool {
    return zk.VerifyProof(proof, statement)
}
```

## Key Differences

| Feature | Go Package | Gno Package |
|---------|-----------|-------------|
| Random Generation | ✅ `crypto/rand` | ❌ Must be off-chain |
| Proof Generation | ✅ Full support | ❌ Off-chain only |
| Proof Verification | ✅ Supported | ✅ Supported |
| State Management | N/A | ✅ Realm packages |
| Gas Costs | N/A | ⚠️ Consider gas |

## Integration with Your Projects

Given your work with:
- **Zcash**: Use this for verifying zk-SNARK proofs on Gno.land
- **Nozy Wallet**: Bridge Zcash privacy to Gno.land
- **Leonine DAO**: Implement privacy-preserving smart contracts

## Next Steps

1. Test locally with `gnodev`
2. Deploy to devnet
3. Integrate with your Zcash/ZK projects
4. Build privacy-preserving smart contracts

## Resources

- [Gno.land Docs](https://docs.gno.land)
- [Gno.land GitHub](https://github.com/gnolang/gno)
- Your GitHub: [@Lowo88](https://github.com/Lowo88)

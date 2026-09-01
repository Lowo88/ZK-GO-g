# Setting Up ZK Package for Gno.land

## Overview

1. **Standard Go** (`zk/`) — off-chain Merkle trees, toy proofs, claim fixture helper
2. **Gno pure package** (`gno/`) — on-chain verify (`gno.land/p/low88/zk`)
3. **ZEC claim realm** (`realms/r/low88/zec_claim/`) — Merkle registry (`gno.land/r/low88/zec_claim`)

v1 verifies **Merkle membership** of Nozy-exported claim leaves. It does **not** verify Orchard/Halo2 or shielded ZEC proofs. See [`SPEC_ZEC_CLAIM_V1.md`](SPEC_ZEC_CLAIM_V1.md).

## Package structure

```
ZK-GO-g/
├── SPEC_ZEC_CLAIM_V1.md
├── zk/                    # Go off-chain
├── gno/                   # Pure package (p/low88/zk)
├── realms/r/low88/zec_claim/  # Claim registry realm
├── fixtures/
├── examples/
└── gnomod.toml            # pure package root
```

## Install Gno tools

```bash
go install github.com/gnolang/gno/gnodev@latest
go install github.com/gnolang/gno/gnokey@latest
```

## Local testing

```bash
go test ./zk/...
go run examples/zec_claim_fixture.go

gnodev test ./gno
```

## Deploy pure package

```bash
gnokey maketx addpkg \
    --pkgpath "gno.land/p/low88/zk" \
    --pkgdir "./gno" \
    --gas-fee "1000000ugnot" \
    --gas-wanted "2000000" \
    --broadcast \
    --chainid "dev"
```

## Deploy ZEC claim realm

```bash
gnokey maketx addpkg \
    --pkgpath "gno.land/r/low88/zec_claim" \
    --pkgdir "./realms/r/low88/zec_claim" \
    --gas-fee "1000000ugnot" \
    --gas-wanted "2000000" \
    --broadcast \
    --chainid "dev"
```

Operator workflow:

1. Build Merkle tree from allowed claim leaves (off-chain `zk.NewMerkleTree`)
2. `SetClaimsRoot(root)` on realm
3. Users `RegisterClaim(leaf, proof)` with proofs from Nozy / `examples/zec_claim_fixture.go`

## Nozy integration

- Nozy: `vote-export-notes` → `nozy zk-gno claim-draft` (NozyWallet repo)
- Public Gno RPC: treat like remote submit (mixnet/local egress policy in Nozy)

Do **not** claim “shielded ZEC verified on Gno” until [`PHASE2_GATE.md`](PHASE2_GATE.md) is satisfied.

## Resources

- [Gno.land Docs](https://docs.gno.land)
- [Gno GitHub](https://github.com/gnolang/gno)

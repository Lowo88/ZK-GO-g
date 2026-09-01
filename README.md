# ZK-Go: Zero Knowledge Primitives for Go and Gno.land

Off-chain Go (`zk/`) and on-chain Gno (`gno/`, `realms/`) for ZK-style primitives and the **Nozy × Gno ZEC claim registry (v1)**.

**Honest v1 scope:** Merkle membership registry for Nozy-exported claim leaves — **not** Orchard/Halo2 SNARK verify. See [`SPEC_ZEC_CLAIM_V1.md`](SPEC_ZEC_CLAIM_V1.md).

## Features

- **Merkle trees** — off-chain build + on-chain `VerifyMerkleProof` (aligned algorithms)
- **Toy hash proofs** — educational only (`zk/proof.go`); not Zcash-grade
- **ZEC claim realm** — `gno.land/r/low88/zec_claim` (`realms/r/low88/zec_claim/`)
- **Nozy glue** — claim leaf spec + fixture for `nozy zk-gno` track (NozyWallet repo)

## Installation

```bash
go get github.com/Lowo88/ZK-GO-g
```

## Quick start

```bash
go test ./zk/...
go run examples/zec_claim_fixture.go
```

Gno local test (requires `gnodev`):

```bash
gnodev test ./gno
```

## Key paths

| Path | Role |
|------|------|
| `SPEC_ZEC_CLAIM_V1.md` | Frozen v1 claim leaf + realm API |
| `gno/` | Pure package `gno.land/p/low88/zk` |
| `realms/r/low88/zec_claim/` | Claim registry realm |
| `fixtures/zec_claim_v1.json` | Integration fixture |
| `PHASE2_GATE.md` | SNARK verify — do not start until API frozen |

## NozyWallet

Nozy exports Ironwood notes → claim draft; this repo verifies Merkle membership on Gno. Parallel track: zk-CosmWasm in [NozyWallet](https://github.com/LEONINE-DAO/Nozy-wallet).

## License

MIT License

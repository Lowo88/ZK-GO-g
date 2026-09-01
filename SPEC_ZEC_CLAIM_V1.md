# ZEC claim registry v1 — Nozy × Gno.land

**Status:** Frozen v1 (Merkle membership registry — not Orchard/Halo2 verify)  
**Nozy format:** `nozy-zk-gno-claim-draft-v1`  
**Gno pure package:** `gno.land/p/low88/zk`  
**Gno realm:** `gno.land/r/low88/zec_claim`

## Statement (v1)

> This leaf commits to Ironwood note claim material exported by Nozy at snapshot height H.

Gno verifies **Merkle membership** of the leaf under an operator-published `claimsRoot`. It does **not** verify Orchard proofs, Halo2, or `proof_instance` bytes.

**Honest claim string:** “Gno verifies Merkle membership of a Nozy-exported claim leaf; not Orchard proof_instance.”

## Operator role (v1)

- Operator (Nozy Labs / DAO) builds a Merkle tree of allowed claim leaves for an epoch.
- Operator calls `SetClaimsRoot(root []byte)` on the realm (admin-only in v1).
- Users call `RegisterClaim(leaf, proof)` with a Merkle proof against that root.
- v1 is a **registry**, not a permissionless Zcash light-client.

## Leaf bytes (frozen)

One leaf per `nozy-vote-notes-v1` export (whole wallet claim, not per-note).

Domain-separated SHA-256 (matches off-chain Nozy; Gno uses `std.Hash` on the same byte sequence only if Gno `std.Hash` is SHA-256 — Nozy always uses SHA-256 for the claim leaf):

```
leaf_input = concat(
  "nozy-zk-gno-leaf-v1\0",           // 20 bytes domain
  utf8(network), "\0",
  le_u32(snapshot_height_or_0),
  le_u64(note_count),
  le_u64(total_value_zatoshis),
  for each note sorted by commitment_hex ascending:
    hex_decode(commitment_hex)        // 32 bytes each
)
claim_leaf = SHA256(leaf_input)
```

Fields sourced from Nozy `VoteNoteExportFile` (`src/vote_export.rs`):

| Field | Source |
|-------|--------|
| `network` | `vote.network` |
| `snapshot_height` | `vote.snapshot_height.unwrap_or(0)` |
| `note_count` | `vote.notes.len()` |
| `total_value_zatoshis` | sum of `note.value` |
| commitments | sorted `note.commitment_hex` |

## Merkle tree (frozen)

Same algorithm as `gno/merkle.gno` (v1 hardened):

- Leaf hash: `Hash(0x00 || claim_leaf)`
- Pair hash: `Hash(0x01 || left || right)`
- Odd leaf count: duplicate last leaf before pairing
- Proof verification starts from `HashLeaf(raw_leaf)`, not the raw leaf
- Empty path is valid for a single-leaf tree (root == HashLeaf(leaf))

## On-chain API (v1)

Realm `gno.land/r/low88/zec_claim`:

| Function | Args | Returns | Notes |
|----------|------|---------|-------|
| `SetClaimsRoot` | `root []byte` | — | Admin only; sets epoch root |
| `RegisterClaim` | `leaf []byte`, `proof *MerkleProof` | `bool` | Verifies proof, checks root, anti-replay |
| `IsRegistered` | `leaf []byte` | `bool` | Query registered status |
| `ClaimsRoot` | — | `[]byte` | Current published root |

`RegisterClaim` logic:

1. `VerifyMerkleProof(proof)` must be true
2. `proof.Root` must equal realm `claimsRoot`
3. `proof.Leaf` must equal the claim `leaf` bytes (raw claim leaf)
4. Store `Hash(leaf)` in `registered` map; reject duplicates

## Off-chain workflow

1. Nozy: `vote-export-notes` → `nozy zk-gno claim-draft`
2. Operator: build tree from allowed leaves → `SetClaimsRoot`
3. Nozy/helper: compute Merkle proof for user's leaf → `RegisterClaim` via gnokey/RPC
4. Public Gno RPC: same mixnet/local egress policy as Zebrad LCD (Nozy `assess_gno_submit_egress`)

## Phase 2 (gated — not v1)

Real Zcash-grade verify (Groth16/BN254 or Halo2) requires a frozen SNARK verify API in this repo. Bump claim format to `nozy-zk-gno-claim-draft-v2` when available. Do **not** claim Orchard/Halo2 verify on Gno until that API exists.

## References

- Nozy demo: `docs/reference/ZK_GNO_NYM_DEMO.md` (NozyWallet repo)
- Parallel track: `docs/reference/ZK_COSMWASM_NYM_DEMO.md`
- Fixture: `fixtures/zec_claim_v1.json`

# Phase 2 gate — real Zcash-grade verify (not v1)

v1 (`SPEC_ZEC_CLAIM_V1.md`) is a **Merkle claim registry**. It does **not** verify Orchard, Halo2, or CosmWasm `proof_instance` bytes.

Do **not** start Phase 2 in NozyWallet until this repo freezes:

1. Circuit choice (Groth16/BN254 off-chain + pairing verify in Gno, or Halo2 research path)
2. Public-input layout and verify API on a Gno realm
3. Claim format bump: `nozy-zk-gno-claim-draft-v2`

Nozy maps `vote_export` → public inputs only after that API exists (same gap as zk-CosmWasm step D).

// ZEC claim registry realm (v1) — Merkle membership, not Orchard verify.
// Deploy: gno.land/r/low88/zec_claim
package zec_claim

import (
	"std"
	"gno.land/p/low88/zk"
)

var (
	claimsRoot []byte
	registered map[string]bool
)

func ensureRegistered() {
	if registered == nil {
		registered = make(map[string]bool)
	}
}

// SetClaimsRoot publishes the operator epoch root (admin gate in production).
func SetClaimsRoot(root []byte) {
	claimsRoot = root
}

// ClaimsRoot returns the current published root.
func ClaimsRoot() []byte {
	return claimsRoot
}

// RegisterClaim verifies Merkle membership and stores the claim leaf hash.
func RegisterClaim(leaf []byte, proof *zk.MerkleProof) bool {
	ensureRegistered()
	if len(claimsRoot) == 0 {
		return false
	}
	if proof == nil || len(leaf) == 0 {
		return false
	}
	if !zk.VerifyMerkleProof(proof) {
		return false
	}
	if !bytesEqual(proof.Root, claimsRoot) {
		return false
	}
	if !bytesEqual(proof.Leaf, leaf) {
		return false
	}
	leafKey := string(std.Hash(leaf))
	if registered[leafKey] {
		return false
	}
	registered[leafKey] = true
	return true
}

// IsRegistered reports whether a claim leaf was registered.
func IsRegistered(leaf []byte) bool {
	ensureRegistered()
	if len(leaf) == 0 {
		return false
	}
	leafKey := string(std.Hash(leaf))
	return registered[leafKey]
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

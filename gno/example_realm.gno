// Example Realm package that uses the ZK pure package
// This demonstrates how to use ZK verification in a stateful smart contract
package main

import (
	"std"
	"gno.land/p/low88/zk"
)

// ZKVerifierRealm is a stateful smart contract for verifying and storing ZK proofs
type ZKVerifierRealm struct {
	// Map of commitment hash -> verification status
	verifiedProofs map[string]bool
	// Map of commitment hash -> statement
	proofStatements map[string]*zk.Statement
}

// NewZKVerifierRealm creates a new ZK verifier realm
func NewZKVerifierRealm() *ZKVerifierRealm {
	return &ZKVerifierRealm{
		verifiedProofs:  make(map[string]bool),
		proofStatements: make(map[string]*zk.Statement),
	}
}

// VerifyAndStore verifies a ZK proof and stores it if valid
// Returns true if proof is valid and was stored, false otherwise
func (zvr *ZKVerifierRealm) VerifyAndStore(proof *zk.Proof, statement *zk.Statement) bool {
	// Verify the proof
	valid := zk.VerifyProof(proof, statement)
	if !valid {
		return false
	}

	// Store proof hash as verified
	proofHash := string(std.Hash(proof.Commitment))
	zvr.verifiedProofs[proofHash] = true
	zvr.proofStatements[proofHash] = statement

	return true
}

// IsVerified checks if a proof commitment has been verified
func (zvr *ZKVerifierRealm) IsVerified(commitment []byte) bool {
	proofHash := string(std.Hash(commitment))
	return zvr.verifiedProofs[proofHash]
}

// GetStatement returns the statement associated with a verified proof
func (zvr *ZKVerifierRealm) GetStatement(commitment []byte) *zk.Statement {
	proofHash := string(std.Hash(commitment))
	return zvr.proofStatements[proofHash]
}

// VerifyCommitment verifies a commitment scheme
// Useful for verifying commitments without full proofs
func (zvr *ZKVerifierRealm) VerifyCommitment(commitment, value, blinding []byte) bool {
	return zk.VerifyCommitment(commitment, value, blinding)
}

// VerifyMerkleMembership verifies membership in a Merkle tree
func (zvr *ZKVerifierRealm) VerifyMerkleMembership(proof *zk.MerkleProof) bool {
	return zk.VerifyMerkleProof(proof)
}

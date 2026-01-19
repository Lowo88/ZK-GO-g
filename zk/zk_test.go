package zk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProofGenerationAndVerification(t *testing.T) {
	statement := &Statement{
		PublicValue: []byte("public statement"),
	}
	witness := &Witness{
		Secret: []byte("secret knowledge"),
	}

	prover := NewProver(statement, witness)
	proof, err := prover.GenerateProof()
	require.NoError(t, err)
	assert.NotNil(t, proof)
	assert.NotNil(t, proof.Commitment)
	assert.NotNil(t, proof.Challenge)
	assert.NotNil(t, proof.Response)

	verifier := NewVerifier(statement)
	valid, err := verifier.Verify(proof)
	require.NoError(t, err)
	assert.True(t, valid)
}

func TestCommitmentScheme(t *testing.T) {
	value := []byte("secret value")

	commitment, err := NewPedersenCommitment(value)
	require.NoError(t, err)
	assert.NotNil(t, commitment)
	assert.NotNil(t, commitment.Commitment)
	assert.NotNil(t, commitment.Blinding)

	opened := commitment.Open(value, commitment.Blinding)
	assert.True(t, opened)

	opened = commitment.Open([]byte("wrong value"), commitment.Blinding)
	assert.False(t, opened)

	opened = commitment.Open(value, []byte("wrong blinding"))
	assert.False(t, opened)
}

func TestMerkleTree(t *testing.T) {
	leaves := [][]byte{
		[]byte("leaf1"),
		[]byte("leaf2"),
		[]byte("leaf3"),
		[]byte("leaf4"),
	}

	tree, err := NewMerkleTree(leaves)
	require.NoError(t, err)
	assert.NotNil(t, tree)
	assert.NotNil(t, tree.Root)
	assert.Equal(t, 4, len(tree.Leaves))

	proof, err := tree.GenerateProof(0)
	require.NoError(t, err)
	assert.NotNil(t, proof)
	assert.NotNil(t, proof.Path)

	valid := VerifyMerkleProof(proof)
	assert.True(t, valid)

	proof, err = tree.GenerateProof(3)
	require.NoError(t, err)
	valid = VerifyMerkleProof(proof)
	assert.True(t, valid)
}

func TestMerkleTreeSingleLeaf(t *testing.T) {
	leaves := [][]byte{
		[]byte("single leaf"),
	}

	tree, err := NewMerkleTree(leaves)
	require.NoError(t, err)
	assert.NotNil(t, tree)

	proof, err := tree.GenerateProof(0)
	require.NoError(t, err)
	
	valid := VerifyMerkleProof(proof)
	assert.True(t, valid)
}

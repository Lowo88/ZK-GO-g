package zk

import (
	"std"
)

type PedersenCommitment struct {
	Commitment []byte
}

func NewPedersenCommitment(value, blinding []byte) (*PedersenCommitment, error) {
	if value == nil {
		return nil, std.NewError("value cannot be nil")
	}
	if blinding == nil {
		return nil, std.NewError("blinding cannot be nil")
	}

	commitment := createCommitmentHash(value, blinding)

	return &PedersenCommitment{
		Commitment: commitment,
	}, nil
}

func VerifyCommitment(commitment, value, blinding []byte) bool {
	expectedCommitment := createCommitmentHash(value, blinding)
	return verifyBytesEqual(commitment, expectedCommitment)
}

func createCommitmentHash(value, blinding []byte) []byte {
	return std.Hash(append(value, blinding...))
}

func verifyBytesEqual(a, b []byte) bool {
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

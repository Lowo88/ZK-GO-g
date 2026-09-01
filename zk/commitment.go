package zk

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
)

type PedersenCommitment struct {
	Commitment []byte
	Blinding   []byte
}

func NewPedersenCommitment(value []byte) (*PedersenCommitment, error) {
	if value == nil {
		return nil, errors.New("value cannot be nil")
	}
	blinding := make([]byte, 32)
	if _, err := rand.Read(blinding); err != nil {
		return nil, fmt.Errorf("failed to generate blinding: %w", err)
	}
	h := sha256.New()
	h.Write(value)
	h.Write(blinding)
	return &PedersenCommitment{
		Commitment: h.Sum(nil),
		Blinding:   blinding,
	}, nil
}

func (c *PedersenCommitment) Open(value, blinding []byte) bool {
	h := sha256.New()
	h.Write(value)
	h.Write(blinding)
	expected := h.Sum(nil)
	return verifyBytesEqual(c.Commitment, expected)
}

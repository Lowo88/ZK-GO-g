package zk

import (
	"std"
)

type Proof struct {
	Commitment []byte
	Challenge  []byte
	Response   []byte
}

type Statement struct {
	PublicValue []byte
}

type Verifier struct {
	statement *Statement
}

func NewVerifier(statement *Statement) *Verifier {
	return &Verifier{
		statement: statement,
	}
}

func (v *Verifier) Verify(proof *Proof) (bool, error) {
	if proof == nil {
		return false, std.NewError("proof cannot be nil")
	}

	if v.statement == nil {
		return false, std.NewError("verifier must have a statement")
	}

	expectedResponse := v.createResponseFromCommitment(proof.Commitment, proof.Challenge)
	
	return verifyBytesEqual(proof.Response, expectedResponse), nil
}

func VerifyProof(proof *Proof, statement *Statement) bool {
	verifier := NewVerifier(statement)
	valid, err := verifier.Verify(proof)
	if err != nil {
		return false
	}
	return valid
}

func (v *Verifier) createResponseFromCommitment(commitment, challenge []byte) []byte {
	return std.Hash(append(commitment, challenge...))
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

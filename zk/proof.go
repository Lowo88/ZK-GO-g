package zk

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
)

type Proof struct {
	Commitment []byte
	Challenge  []byte
	Response   []byte
}

type Statement struct {
	PublicValue []byte
}

type Witness struct {
	Secret []byte
}

type Prover struct {
	statement *Statement
	witness   *Witness
}

type Verifier struct {
	statement *Statement
}

func NewProver(statement *Statement, witness *Witness) *Prover {
	return &Prover{
		statement: statement,
		witness:   witness,
	}
}

func NewVerifier(statement *Statement) *Verifier {
	return &Verifier{
		statement: statement,
	}
}

func (p *Prover) GenerateProof() (*Proof, error) {
	if p.statement == nil || p.witness == nil {
		return nil, errors.New("prover must have both statement and witness")
	}

	nonce := make([]byte, 32)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	commitment := p.createCommitment(p.witness.Secret, nonce)

	challenge := make([]byte, 32)
	if _, err := rand.Read(challenge); err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	response := p.createResponse(p.witness.Secret, challenge, nonce)

	return &Proof{
		Commitment: commitment,
		Challenge:  challenge,
		Response:   response,
	}, nil
}

func (v *Verifier) Verify(proof *Proof) (bool, error) {
	if proof == nil {
		return false, errors.New("proof cannot be nil")
	}

	if v.statement == nil {
		return false, errors.New("verifier must have a statement")
	}

	expectedResponse := v.createResponseFromCommitment(proof.Commitment, proof.Challenge)
	
	return verifyBytesEqual(proof.Response, expectedResponse), nil
}

func (p *Prover) createCommitment(secret, nonce []byte) []byte {
	h := sha256.New()
	h.Write(secret)
	h.Write(nonce)
	return h.Sum(nil)
}

func (p *Prover) createResponse(secret, challenge, nonce []byte) []byte {
	h := sha256.New()
	h.Write(secret)
	h.Write(challenge)
	h.Write(nonce)
	return h.Sum(nil)
}

func (v *Verifier) createResponseFromCommitment(commitment, challenge []byte) []byte {
	h := sha256.New()
	h.Write(commitment)
	h.Write(challenge)
	return h.Sum(nil)
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

package main

import (
	"fmt"
	"log"

	"github.com/Lowo88/ZK-GO-g/zk"
)

func main() {
	fmt.Println("Zero Knowledge Proof Example")
	fmt.Println("============================")

	// Example 1: Basic Proof
	fmt.Println("\n1. Basic Zero-Knowledge Proof:")
	exampleBasicProof()

	// Example 2: Commitment Scheme
	fmt.Println("\n2. Commitment Scheme:")
	exampleCommitment()

	// Example 3: Merkle Tree
	fmt.Println("\n3. Merkle Tree Proof:")
	exampleMerkleTree()
}

func exampleBasicProof() {
	statement := &zk.Statement{
		PublicValue: []byte("I know a secret"),
	}

	witness := &zk.Witness{
		Secret: []byte("my secret password"),
	}

	prover := zk.NewProver(statement, witness)

	proof, err := prover.GenerateProof()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("  Generated proof with commitment: %x...\n", proof.Commitment[:8])

	verifier := zk.NewVerifier(statement)

	valid, err := verifier.Verify(proof)
	if err != nil {
		log.Fatal(err)
	}

	if valid {
		fmt.Println("  ✓ Proof verified successfully!")
		fmt.Println("  Note: Verifier confirmed knowledge without seeing the secret")
	} else {
		fmt.Println("  ✗ Proof verification failed")
	}
}

func exampleCommitment() {
	secretValue := []byte("I commit to this value")

	commitment, err := zk.NewPedersenCommitment(secretValue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("  Committed to value (hash): %x\n", commitment.Commitment[:8])
	fmt.Println("  Commitment created - value is hidden")

	opened := commitment.Open(secretValue, commitment.Blinding)
	if opened {
		fmt.Println("  ✓ Commitment opened successfully")
		fmt.Println("  Original value revealed and verified")
	}
}

func exampleMerkleTree() {
	items := [][]byte{
		[]byte("item1"),
		[]byte("item2"),
		[]byte("item3"),
		[]byte("item4"),
	}

	tree, err := zk.NewMerkleTree(items)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("  Merkle tree root: %x\n", tree.Root[:8])
	fmt.Printf("  Tree depth: %d\n", tree.Depth)

	proof, err := tree.GenerateProof(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("  Generated proof for item: %s\n", string(proof.Leaf))
	fmt.Printf("  Proof path length: %d\n", len(proof.Path))

	valid := zk.VerifyMerkleProof(proof)
	if valid {
		fmt.Println("  ✓ Merkle proof verified - item is in the tree")
	} else {
		fmt.Println("  ✗ Merkle proof verification failed")
	}
}

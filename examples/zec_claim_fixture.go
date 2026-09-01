//go:build ignore

// Run: go run examples/zec_claim_fixture.go
// Prints fixture JSON fields for fixtures/zec_claim_v1.json and gnokey call hints.
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/Lowo88/ZK-GO-g/zk"
)

func main() {
	leaves := [][]byte{
		[]byte("claim-a"),
		[]byte("claim-b"),
	}
	tree, err := zk.NewMerkleTree(leaves)
	if err != nil {
		panic(err)
	}
	proof, err := tree.GenerateProof(0)
	if err != nil {
		panic(err)
	}

	out := map[string]interface{}{
		"claims_root_hex": hex.EncodeToString(tree.Root),
		"leaf_hex":        hex.EncodeToString(proof.Leaf),
		"proof_leaf_hex":  hex.EncodeToString(proof.Leaf),
		"path_hex":        hexSlice(proof.Path),
		"indices":         proof.Indices,
		"root_hex":        hex.EncodeToString(proof.Root),
		"gnokey_hint":     "gnokey maketx call -pkgpath gno.land/r/low88/zec_claim -func RegisterClaim -args <leaf_hex> ...",
	}
	b, _ := json.MarshalIndent(out, "", "  ")
	fmt.Println(string(b))
}

func hexSlice(b [][]byte) []string {
	out := make([]string, len(b))
	for i, x := range b {
		out[i] = hex.EncodeToString(x)
	}
	return out
}

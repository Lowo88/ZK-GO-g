package zk

import (
	"crypto/sha256"
	"errors"
)

type MerkleTree struct {
	Root   []byte
	Leaves [][]byte
	Depth  int
}

type MerkleProof struct {
	Leaf    []byte
	Path    [][]byte
	Indices []int
	Root    []byte
}

func NewMerkleTree(leaves [][]byte) (*MerkleTree, error) {
	if len(leaves) == 0 {
		return nil, errors.New("cannot create tree with no leaves")
	}

	depth := calculateDepth(len(leaves))
	root := buildMerkleTree(leaves)

	return &MerkleTree{
		Root:   root,
		Leaves: leaves,
		Depth:  depth,
	}, nil
}

func (mt *MerkleTree) GenerateProof(leafIndex int) (*MerkleProof, error) {
	if leafIndex < 0 || leafIndex >= len(mt.Leaves) {
		return nil, errors.New("invalid leaf index")
	}

	path, indices := mt.getMerklePath(leafIndex)

	return &MerkleProof{
		Leaf:    mt.Leaves[leafIndex],
		Path:    path,
		Indices: indices,
		Root:    mt.Root,
	}, nil
}

// VerifyMerkleProof checks membership. Starts from HashLeaf(raw leaf).
func VerifyMerkleProof(proof *MerkleProof) bool {
	if proof == nil || len(proof.Leaf) == 0 || len(proof.Root) == 0 {
		return false
	}
	if len(proof.Path) != len(proof.Indices) {
		return false
	}

	current := HashLeaf(proof.Leaf)
	for i, sibling := range proof.Path {
		if proof.Indices[i] == 0 {
			current = HashPair(current, sibling)
		} else {
			current = HashPair(sibling, current)
		}
	}

	return verifyBytesEqual(current, proof.Root)
}

func buildMerkleTree(leaves [][]byte) []byte {
	level := make([][]byte, len(leaves))
	for i, leaf := range leaves {
		level[i] = HashLeaf(leaf)
	}

	for len(level) > 1 {
		if len(level)%2 != 0 {
			level = append(level, level[len(level)-1])
		}
		next := make([][]byte, len(level)/2)
		for i := 0; i < len(level)/2; i++ {
			next[i] = HashPair(level[i*2], level[i*2+1])
		}
		level = next
	}
	return level[0]
}

func (mt *MerkleTree) getMerklePath(leafIndex int) ([][]byte, []int) {
	var path [][]byte
	var indices []int

	level := make([][]byte, len(mt.Leaves))
	for i, leaf := range mt.Leaves {
		level[i] = HashLeaf(leaf)
	}
	currentIndex := leafIndex

	for len(level) > 1 {
		if len(level)%2 != 0 {
			level = append(level, level[len(level)-1])
		}
		siblingIndex := currentIndex ^ 1
		path = append(path, level[siblingIndex])
		indices = append(indices, currentIndex%2)

		next := make([][]byte, len(level)/2)
		for i := 0; i < len(level)/2; i++ {
			next[i] = HashPair(level[i*2], level[i*2+1])
		}
		level = next
		currentIndex = currentIndex / 2
	}

	return path, indices
}

func HashLeaf(leaf []byte) []byte {
	h := sha256.New()
	h.Write([]byte{0x00})
	h.Write(leaf)
	return h.Sum(nil)
}

func HashPair(left, right []byte) []byte {
	h := sha256.New()
	h.Write([]byte{0x01})
	h.Write(left)
	h.Write(right)
	return h.Sum(nil)
}

func calculateDepth(leafCount int) int {
	depth := 0
	for leafCount > 1 {
		leafCount = (leafCount + 1) / 2
		depth++
	}
	return depth
}

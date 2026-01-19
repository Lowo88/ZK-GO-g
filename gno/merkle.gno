package zk

import (
	"std"
)

type MerkleTree struct {
	Root   []byte
	Leaves [][]byte
	Depth  int
}

type MerkleProof struct {
	Leaf      []byte
	Path      [][]byte
	Indices   []int
	Root      []byte
}

func NewMerkleTree(leaves [][]byte) (*MerkleTree, error) {
	if len(leaves) == 0 {
		return nil, std.NewError("cannot create tree with no leaves")
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
		return nil, std.NewError("invalid leaf index")
	}

	path, indices := mt.getMerklePath(leafIndex)

	return &MerkleProof{
		Leaf:    mt.Leaves[leafIndex],
		Path:    path,
		Indices: indices,
		Root:    mt.Root,
	}, nil
}

func VerifyMerkleProof(proof *MerkleProof) bool {
	if proof == nil || len(proof.Path) == 0 {
		return false
	}

	current := proof.Leaf
	for i, sibling := range proof.Path {
		if proof.Indices[i] == 0 {
			current = hashPair(current, sibling)
		} else {
			current = hashPair(sibling, current)
		}
	}

	return verifyBytesEqual(current, proof.Root)
}

func buildMerkleTree(leaves [][]byte) []byte {
	if len(leaves) == 1 {
		return hashLeaf(leaves[0])
	}

	if len(leaves)%2 != 0 {
		leaves = append(leaves, leaves[len(leaves)-1])
	}

	nextLevel := make([][]byte, len(leaves)/2)
	for i := 0; i < len(leaves)/2; i++ {
		nextLevel[i] = hashPair(leaves[i*2], leaves[i*2+1])
	}

	return buildMerkleTree(nextLevel)
}

func (mt *MerkleTree) getMerklePath(leafIndex int) ([][]byte, []int) {
	var path [][]byte
	var indices []int

	currentIndex := leafIndex
	leaves := mt.Leaves

	for len(leaves) > 1 {
		siblingIndex := currentIndex ^ 1
		if siblingIndex < len(leaves) {
			path = append(path, leaves[siblingIndex])
			indices = append(indices, currentIndex%2)
		}

		currentIndex = currentIndex / 2
		if len(leaves)%2 != 0 {
			leaves = append(leaves, leaves[len(leaves)-1])
		}
		nextLevel := make([][]byte, len(leaves)/2)
		for i := 0; i < len(leaves)/2; i++ {
			nextLevel[i] = hashPair(leaves[i*2], leaves[i*2+1])
		}
		leaves = nextLevel
	}

	return path, indices
}

func hashLeaf(leaf []byte) []byte {
	return std.Hash(append([]byte{0x00}, leaf...))
}

func hashPair(left, right []byte) []byte {
	return std.Hash(append(append([]byte{0x01}, left...), right...))
}

func calculateDepth(leafCount int) int {
	depth := 0
	for leafCount > 1 {
		leafCount = (leafCount + 1) / 2
		depth++
	}
	return depth
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

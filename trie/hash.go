package trie

import (
	"crypto/sha256"
	"golang.org/x/crypto/sha3"
	"common"
)

type MerkleHash []byte

func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// Keccak256Hash calculates and returns the Keccak256 hash of the input data,
// converting it to an internal Hash data structure.
func Keccak256Hash(data ...[]byte) (h common.Hash) {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(h[:0])
	return h
}

func Sha256(data ...[]byte) []byte {
	d := sha256.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
package trietest

import (
	mpt "github.com/Myriad-Dreamin/go-mpt/trie"
	"testing"
)

func TestBuildTrie(t *testing.T) {
	db, err := mpt.NewNodeBase("./testdb")
	if err != nil {
		t.Error(err)
		return
	}
	defer db.Close()
	// var tr *mpt.Trie
	_, err = mpt.NewTrie(mpt.HexToHash("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"), db)
	if err != nil {
		t.Error(err)
		return
	}
}
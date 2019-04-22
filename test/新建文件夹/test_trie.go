package main

import (
	"fmt"
	mpt "github.com/Myriad-Dreamin/go-mpt/trie"
)

func main() {
	db, err := mpt.NewNodeBase("./testdb")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	db.Put([]byte("key"), []byte("value"))
	fmt.Println(db.Get([]byte("key")))
	
}
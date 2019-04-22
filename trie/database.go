package trie

import (
	"github.com/syndtr/goleveldb/leveldb"
	"common"
)

type NodeBase struct {
	handler *leveldb.DB
}

func (db *NodeBase) Get(toGet []byte) ([]byte, error) {
	return db.handler.Get(toGet, nil)
}

func (db *NodeBase) Put(key []byte, value []byte) error {
	return db.handler.Put(key, value, nil)
}

func (db *NodeBase) Close() error {
	return db.handler.Close()
}

func NewNodeBase(path string) (*NodeBase, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &NodeBase{handler: db}, nil
}

func (db *NodeBase) node(hash common.Hash) node {
	enc, err := db.handler.Get(hash[:], nil)
	if err != nil || enc == nil {
		return nil
	}
	return mustDecodeNode(hash[:], enc)
}

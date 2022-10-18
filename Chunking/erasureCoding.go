package erasure

import "sync"

type Code struct {
	M int
	K int
	ShardLength int
	EncodeMatrix []byte
	galoisTables []byte
	decode *decodeNode
}

type decodeNode struct {
	children []*decodeNode
	mutex *sync.Mutex
	galoisTables []byte
	decodeNode []byte
}
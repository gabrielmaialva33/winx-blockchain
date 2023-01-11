package block

import "time"

type Block struct {
	Hash         [32]byte
	PreviousHash [32]byte
	Nonce        uint64
	Timestamp    int64
	Data         []byte // transactions
}

// NewBlock creates and returns Block
func NewBlock(nonce uint64, previousHash [32]byte, data []byte) *Block {
	return &Block{
		PreviousHash: previousHash,
		Nonce:        nonce,
		Timestamp:    time.Now().Unix(),
		Data:         data,
	}
}

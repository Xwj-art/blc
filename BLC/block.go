package BLC

import (
	"time"
)

type Block struct {
	TimeStamp     int64
	Hash          []byte
	PrevBlockHash []byte
	Height        int64
	Data          []byte
	Nonce         int64
}

func NewBlock(height int64, prevBlockHash, data []byte) *Block {
	block := Block{
		TimeStamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Data:          data,
	}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = int64(nonce)
	return &block
}

//
//func (b *Block) SetHash() {
//	timeStampBytes := IntToHex(b.TimeStamp)
//	heightBytes := IntToHex(b.Height)
//	blockBytes := bytes.Join([][]byte{
//		timeStampBytes,
//		heightBytes,
//		b.PrevBlockHash,
//		b.Data,
//	}, []byte{})
//	hash := sha256.Sum256(blockBytes)
//	b.Hash = hash[:]
//}

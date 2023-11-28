package BLC

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256-targetBit)
	return &ProofOfWork{
		Block:  block,
		target: target,
	}
}

func (pow *ProofOfWork) Run() ([]byte, int) {
	var nonce = 0
	var HashInt big.Int
	var hash [32]byte
	for {
		dataBytes := pow.prepareData(int64(nonce))
		hash = sha256.Sum256(dataBytes)
		HashInt.SetBytes(hash[:])
		if pow.target.Cmp(&HashInt) == 1 {
			break
		}
		nonce++
	}
	return hash[:], nonce
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	var data []byte
	timeStampBytes := IntToHex(pow.Block.TimeStamp)
	heightBytes := IntToHex(pow.Block.Height)
	data = bytes.Join([][]byte{
		timeStampBytes,
		heightBytes,
		pow.Block.PrevBlockHash,
		pow.Block.Data,
		IntToHex(nonce),
		IntToHex(targetBit),
	}, []byte{})
	return data
}

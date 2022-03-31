package blockchain

import (
	"sync"

	"github.com/k151202/gocoin/db"
	"github.com/k151202/gocoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height int `json:"height"`
	
}

var b *blockchain
var once sync.Once

// restore blocks from checkpoint
func (b *blockchain) restore(data []byte){
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

// adds a block
func (b *blockchain) AddBlock(data string){
	block := createBlock(data, b.NewestHash, b.Height+1)
	// update last hash and height
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// get all blocks
func (b *blockchain) Blocks() []*Block{
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

// initialize blockchain or restore checkpoint
func Blockchain()*blockchain {
	if b == nil{
		once.Do(func(){
			// make an empty blockchain
			b = &blockchain{"", 0}	
			// search for checkpoint on the db
			checkpoint := db.Checkpoint()
			// restore b from bytes
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				b.restore(checkpoint)
			}
		})
	}
	return b
}



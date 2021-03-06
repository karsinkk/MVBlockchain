package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Genesis block. First block in the chain.
var Genesis Block = Block{data: "Karsin", timeStamp: time.Now()}
var Latest *Block = &Genesis

// Block Structure
type Block struct {
	id          int
	prev        *Block
	data        string
	nonce       int
	prevHash    *string
	currentHash string
	timeStamp   time.Time
}

// Insert a new block in the chain.
func Insert(data string) int {

	newBlock := Block{id: Latest.id + 1, prev: Latest, data: data,
		prevHash: &Latest.currentHash, timeStamp: time.Now()}
	Latest = &newBlock
	return Latest.id
}

// POW algorithm -  Returns the nonce and computed hash.
func (block *Block) ComputeHash() (int, string) {
	data := fmt.Sprintf("%s%s", block.data, *block.prevHash)
	var z []byte
	for i := 0; ; i++ {
		s := fmt.Sprintf("%s%d", data, i)
		h := sha256.New()
		h.Write([]byte(s))
		z = h.Sum(nil)
		if (z[0] == 00) && (z[1] == 00) {
			t := fmt.Sprintf("%x", z)
			block.currentHash = t
			block.nonce = i
			return i, t
		}
	}
}

func main() {
	fmt.Println(time.Now())
	fmt.Println(Latest)
	fmt.Println(Insert("Karsin Kamakotti"))
	fmt.Println(*Latest.prevHash)
	fmt.Println(Insert("Karsin Kotti"))
	fmt.Println(*Latest.prevHash)
}

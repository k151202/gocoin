package main

import (
	"fmt"
	"net/http"
)

const port string = ":4000"

func main(){
	// chain := blockchain.GetBlockchain()
	// chain.AddBlock("Second Block")
	// chain.AddBlock("Thrid Block")
	// chain.AddBlock("Fourth Block")
	// for _, block := range chain.AllBlocks() {
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %s\n", block.Hash)
	// 	fmt.Printf("Prev Hash: %s\n", block.PrevHash)
	// }
	fmt.Printf("Listening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
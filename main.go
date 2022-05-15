package main

import (
	"NeuroResearchStorage/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.InitBlockchain()
	fmt.Println(" *** Welcome to blockchain demo! ***\nType help for list of available commands")
	cli(&bc)
}

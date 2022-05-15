package blockchain

import (
	"fmt"
	"log"
)

type Blockchain []*Block

var capacity = 10

// InitBlockchain Создание блокчейна на основе генезис-блока, который ключает в себя одну транзакцию с README.md
func InitBlockchain() Blockchain {
	initTx, err := CreateTransaction("README.md", "README.md", 0)
	if err != nil {
		log.Fatalln("Failed to initialize blockchain:", err.Error())
	}
	genesis, _ := CreateBlock([32]byte{}, []*Transaction{initTx})
	bc := make([]*Block, 1, capacity) // capacity improves performance
	bc[0] = genesis
	return bc
}

// ConnectBlock Присоединить блок в блокчейн; проверка на валидность не проводится,
// чтобы можно было вставить некорректные блоки и показать нарушение
// целостности блокчейна и роль связей каждого блока с предыдущим
func (bc *Blockchain) ConnectBlock(b *Block) {
	// optimizes appending
	if len(*bc) < cap(*bc) {
		*bc = append(*bc, b)
		return
	}
	// extend capacity
	capacity *= 10
	newbc := make([]*Block, len(*bc), capacity)
	// rewrite blockchain
	for i, v := range *bc {
		newbc[i] = v
	}
	newbc[len(*bc)] = b
}

// ValidateBlock Проверка блока, совпадает ли его хеш-значение с ID предыдущего блока
func (bc *Blockchain) ValidateBlock(b *Block) bool {
	return b.HashPrev == bc.GetLastBlock().ID
}

// ValidateChain Проверка всего блокчейна; как только найдено нарушение, дальше можно не смотреть
func (bc *Blockchain) ValidateChain() (invalidBlockNumber int) {
	for i := range (*bc)[1:] {
		if (*bc)[i+1].HashPrev != (*bc)[i].ID {
			return i
		}
	}
	return -1
}

// GetLastBlock Получить последний блок в цепочке
func (bc *Blockchain) GetLastBlock() *Block {
	if len(*bc) == 0 {
		log.Fatalln("Unable to get last block: blockchain is not initialized")
	}
	return (*bc)[len(*bc)-1]
}

// GetTotalTxCount Всего транзакций в блоках цепочки
func (bc *Blockchain) GetTotalTxCount() (count int) {
	for _, b := range *bc {
		count += len(b.TxList)
	}
	return
}

// Преобразовать в строку; verbose здесь касается только блоков
func (bc *Blockchain) toString(verbose bool) string {
	result := "{"
	for i, b := range *bc {
		result += b.toString(verbose)
		if i < len(*bc)-1 {
			result += ","
		}
	}
	return result + "\n}"
}

func (bc *Blockchain) Print(verbose bool) {
	fmt.Println(bc.toString(verbose))
}

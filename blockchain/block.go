package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type Block struct {
	ID       [32]byte       // Хеш-значение полей блока
	HashPrev [32]byte       // Хеш-значение предыдущего блока
	TxList   []*Transaction // Список транзакций
}

// CreateBlock Создать блок; его хеш-значения получается из конкатенации хеш-значений входящих в него транзакций
func CreateBlock(hashPrev [32]byte, txList []*Transaction) (b *Block, err error) {
	if len(txList) == 0 {
		return nil, errors.New("Failed to create block: no transactions to insert")
	}
	b = &Block{HashPrev: hashPrev, TxList: txList}
	txHashes := make([]byte, 32*len(txList)) // 32 is length of SHA-2 hash value
	// concatenated hash values of each tx
	for i, tx := range txList {
		for j, v := range tx.hash {
			txHashes[i*32+j] = v
		}
	}
	b.ID = sha256.Sum256(txHashes)
	return
}

// toString Конвертация в строку; verbose указывает, нужен ли подробный вывод
// Пришлось дублировать код из transaction.go для более читаемого вывода, проставив отступы
func (b *Block) toString(verbose bool) string {
	result := fmt.Sprintf(`
{
	ID: %x
	HashPrev: %x`, b.ID, b.HashPrev)
	if verbose {
		result += fmt.Sprint(`
	Transactions:`)
		for i, tx := range b.TxList {
			// copied to beautify output
			result += fmt.Sprintf(`
	{
		hash: %x,
		title: %q,
		filename: %q,
		nonce: %v
	}`, tx.hash, tx.title, tx.filename, tx.nonce)
			if i < len(b.TxList)-1 {
				result += ","
			}
		}
	}
	return result + "\n}"
}

func (b *Block) Print(verbose bool) {
	fmt.Println(b.toString(verbose))
}

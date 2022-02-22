package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Transaction struct {
	hash     []uint8 // transactionID
	title    string
	document *os.File // transaction data
	nonce    int
}

// func hashToString(h []uint8) string {
// 	return fmt.Sprintf("%x", h)
// }

func checkForError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func createTransaction(t string, doc *os.File, n int) (tx Transaction) {
	// t = strings.Join(strings.Fields(t), " ")
	// if len(t) < 3 {
	// 	log.Fatalln("Unable to create transaction: title too short")
	// }
	// not sure if that is necessary
	tx.title = t
	tx.document = doc
	tx.nonce = n
	tmpHash := sha256.New()
	fileContent, e := os.ReadFile(doc.Name())
	checkForError(e)
	nonceByte := []byte(strconv.Itoa(n))
	tmpHash.Write(append(fileContent, nonceByte...))
	tx.hash = tmpHash.Sum(nil)
	// if _, err := io.Copy(tmpHash, doc); err != nil {
	// 	log.Fatal(err)
	// }
	return
}

func (tx *Transaction) toString() (res string) {
	return fmt.Sprintf(`
{
	Hash: %x
	Title: %q
	File name: %q
	Nonce: %v
}`, tx.hash, tx.title, tx.document.Name(), tx.nonce)
}

func (tx Transaction) print() {
	fmt.Println(tx.toString())
}

func main() {
	f, e := os.Open("bitcoin_ru.pdf")
	checkForError(e)
	defer f.Close()
	t := createTransaction("Это исследование Сатоши Накамото", f, 25)
	t.print()
}

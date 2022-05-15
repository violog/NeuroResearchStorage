package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Transaction struct {
	hash     [32]byte
	title    string
	filename string
	content  []byte
	nonce    int
}

func CreateTransaction(t, fname string, n int) (tx *Transaction, err error) {
	t = strings.Join(strings.Fields(t), " ") // strip whitespace
	// storing as constants to make changes in a single place
	const (
		attachmentSizeLimit = 1073741824
		minTitleLength      = 8
		maxTitleLength      = 128
		creationFailure     = "Failed to create transaction: "
	)
	if len(t) < minTitleLength {
		return nil, errors.New(fmt.Sprintf(creationFailure+"title too short: min length is %v, got %v\n", minTitleLength, len(t)))
	} else if len(t) > maxTitleLength {
		return nil, errors.New(fmt.Sprintf(creationFailure+"title too long: max length is %v, got %v\n", maxTitleLength, len(t)))
	}
	invalid := regexp.MustCompile(`(?i)[^a-zа-яёіїґ\d .,:%'"\?()-]`)
	if invalid.MatchString(t) {
		return nil, errors.New(creationFailure + "given name (" + t + ") doesn't match rules (only latin and cyryllic, numbers and some special chars)")
	}
	tx = &Transaction{title: t}
	// prevent loading common viruses into blockchain
	exe := regexp.MustCompile("sh|exe|jar")
	fext := filepath.Ext(fname)
	if exe.MatchString(fext) {
		return nil, errors.New(creationFailure + "file extension " + fext + " is not allowed")
	}
	tx.filename = fname
	tx.nonce = n
	info, err := os.Stat(fname)
	if err != nil {
		return nil, err
	}
	if info.Size() > attachmentSizeLimit {
		return nil, errors.New(fmt.Sprintf(creationFailure+"file size limit exceeded: max is %v, got %v", attachmentSizeLimit, info.Size()))
	}
	// no sense in protection of low file size: file with required size can be easily created and used for attacks
	tx.content, err = os.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	tx.Hash()
	return
}

func (tx *Transaction) toString() (res string) {
	return fmt.Sprintf(`
{
	hash: %x,
	title: %q,
	filename: %q,
	nonce: %v
}`, tx.hash, tx.title, tx.filename, tx.nonce)
}

func (tx *Transaction) Hash() [32]byte {
	// Hash byte seq of data in order: title, filename, nonce, content
	byteNonce := []byte(strconv.Itoa(tx.nonce))
	txDataHashed := make([]byte, len(tx.title)+len(tx.filename)+len(tx.content)+len(byteNonce))
	for i, v := range []byte(tx.title) {
		txDataHashed[i] = v
	}
	for i, v := range []byte(tx.filename) {
		txDataHashed[i+len(tx.title)] = v
	}
	for i, v := range byteNonce {
		txDataHashed[i+len(tx.title)+len(tx.filename)] = v
	}
	for i, v := range tx.content {
		txDataHashed[i+len(tx.title)+len(tx.filename)+len(byteNonce)] = v
	}
	tx.hash = sha256.Sum256(txDataHashed)
	return tx.hash
}

func (tx Transaction) Print() {
	fmt.Println(tx.toString())
}

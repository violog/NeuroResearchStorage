package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// Every tx must contain an attached file: how to do that?
// Remember file's name, read bytes, then write it when destination is reached
type Transaction struct {
	hash     [32]byte
	title    string
	filename string
	content  []byte
	nonce    int
}

// For errors, let it be exit for now; I'll change the behaviour when I get more info about the project
func checkForError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func throwRequiredFoundError(msg, part string, req, found int) {
	log.Fatalln(fmt.Sprintf("%v (required %v %v, but found %v)\n", msg, part, req, found))
}

func createTransaction(t, fname string, n int) (tx Transaction) {
	// idea of error handling: return err with tx
	t = strings.Join(strings.Fields(t), " ") // strip whitespace
	// storing as constants to make changes in a single place
	const (
		attachmentSizeLimit = 1073741824
		minTitleLength      = 8
		maxTitleLength      = 128
		creationFailure     = "Failed to create transaction: "
		mt                  = "more than"
		lt                  = "less than"
	)
	if len(t) < minTitleLength {
		throwRequiredFoundError(creationFailure+"title too short", mt, minTitleLength+1, len(t))
	} else if len(t) > maxTitleLength {
		throwRequiredFoundError(creationFailure+"title too long", lt, maxTitleLength+1, len(t))
	}
	invalid, _ := regexp.Compile(`(?i)[^a-zа-яёіїґ\d .,:%'"\?()-]`)
	if invalid.MatchString(t) {
		log.Fatalln(creationFailure + "given name (" + t + ") doesn't match rules (only latin and cyryllic, numbers and some special chars)")
	}
	tx.title = t
	// prevent loading common viruses into blockchain
	exe, e := regexp.Compile("sh|exe|jar")
	fext := filepath.Ext(fname)
	if exe.MatchString(fext) {
		log.Fatalln(creationFailure + "file extension " + fext + " is not allowed")
	}
	tx.filename = fname
	tx.nonce = n
	info, e := os.Stat(fname)
	if info.Size() > attachmentSizeLimit {
		throwRequiredFoundError(creationFailure+"file size limit exceeded", lt, attachmentSizeLimit, int(info.Size()))
	}
	// no sense in protection of low file size: file with required size can be easily created and used for attacks
	tx.content, e = os.ReadFile(fname)
	checkForError(e)
	// Hash byte seq of data in order: title, filename, nonce, content
	txDataHashed := []byte(tx.title)
	txDataHashed = append(txDataHashed, []byte(tx.filename)...)
	txDataHashed = append(txDataHashed, []byte(strconv.Itoa(n))...)
	tx.hash = sha256.Sum256(append(txDataHashed, tx.content...))
	return
}

func (tx *Transaction) toString() (res string) {
	return fmt.Sprintf(`
{
	hash: %x
	title: %q
	filename: %q
	nonce: %v
}`, tx.hash, tx.title, tx.filename, tx.nonce)
}

func (tx Transaction) print() {
	fmt.Println(tx.toString())
}

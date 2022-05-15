// Функции-утилиты, используемые в CLI
package main

import (
	"NeuroResearchStorage/blockchain"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// getTxFromUserInput Создание транзакции из консоли пользователя
func getTxFromUserInput() (*blockchain.Transaction, error) {
	var t, fname string
	var n int
	fmt.Print("Enter the title (e.g. My research of Something): ")
	t, err := readLine(os.Stdin)
	if err != nil {
		return nil, err
	}
	fmt.Print("Enter input file's name (e.g. bitcoin_ru.pdf): ")
	fname, err = readLine(os.Stdin)
	if err != nil {
		return nil, err
	}
	fmt.Print("Enter nonce (e.g. 0): ")
	_, err = fmt.Scan(&n)
	if err != nil {
		return nil, err
	}
	return blockchain.CreateTransaction(t, fname, n)
}

// readAnswer Считать ответ: да/нет
func readAnswer(question string) bool {
	var ans string
	for {
		fmt.Printf("%v [y/n]: ", question)
		fmt.Scan(&ans)
		if ans == "Y" || ans == "y" {
			return true
		} else if ans == "N" || ans == "n" {
			return false
		} else {
			fmt.Println("Invalid answer: " + ans + ". Enter y for \"yes\", n for \"no\"\n")
		}
	}
}

// readLine Считать строку до символа \n вместо слова
func readLine(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input[:len(input)-1], nil // remove newline char
}

func printHelp() {
	fmt.Println(`
Blockchain demo manual

Description
All necessary information and additional help can be found in README.md.
Any command doesn't need arguments: just type it, press Enter and follow
instructions.

Commands:
b, block - insert pending transactions into a block and mark it pending
c, connect - connect pending blocks to the blockchain
h, help - print this help
i, info - print blockchain and mempool info
m, mode - switch modes (validation, verbose)
p, print - print blockchain
q, quit - leave app
t, tx	- create transaction and mark it pending
v, validate - validate blockchain
w, wipe - wipe pending transactions or blocks
`)
}

// strTo32Byte Конвертация строки в массив длиной 32 байта - нужно для проверки ввода хеша с консоли
func strTo32Byte(s string) [32]byte {
	var result [32]byte
	hexToDec := func(hex byte) byte {
		res, _ := strconv.ParseUint(string(hex), 16, 8)
		return byte(res)
	}
	for i := range result {
		result[i] = hexToDec(s[2*i])*16 + hexToDec(s[2*i+1])
	}
	return result
}

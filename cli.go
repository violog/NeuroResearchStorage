// Функционал интерфейса командной строки
package main

import (
	"NeuroResearchStorage/blockchain"
	"errors"
	"fmt"
	"log"
	"regexp"
)

// cli Интерфейс командной строки, принимающий на вход инициализированный блокчейн
func cli(bc *blockchain.Blockchain) {
	// Инициализация переменных
	const mempoolCapacity = 16
	var cmd string
	// Параметры валидации перед присоединением блока и подробного вывода
	var validateBeforeConnecting = true
	var verbose = false
	// Неподтверждённые транзакции и блоки
	pendingTxs := make([]*blockchain.Transaction, 0, mempoolCapacity)
	pendingBlocks := make([]*blockchain.Block, 0, mempoolCapacity)

	// Замыкания в данном контексте удобнее функций утилит
	// Добавить транзакцию в мемпул
	addPendingTx := func(tx *blockchain.Transaction) error {
		if tx == nil {
			return errors.New("Attempt to add empty transaction")
		}
		for _, v := range pendingTxs {
			if tx.Hash() == v.Hash() {
				return errors.New("Duplicate transactions: same transaction already exists in mempool")
			}
		}
		pendingTxs = append(pendingTxs, tx)
		return nil
	}

	// Добавить блок в условный мемпул, где находятся неподтверждённые блоки
	addPendingBlock := func(hashPrev [32]byte) (*blockchain.Block, error) {
		block, err := blockchain.CreateBlock(hashPrev, pendingTxs)
		if err != nil {
			return nil, err
		}
		pendingBlocks = append(pendingBlocks, block)
		pendingTxs = make([]*blockchain.Transaction, 0, mempoolCapacity)
		return block, nil
	}

	// Присоединить все блоки из мемпула
	connectPendingBlocks := func() (connected int, err error) {
		if validateBeforeConnecting {
			for i, b := range pendingBlocks {
				if bc.ValidateBlock(b) {
					bc.ConnectBlock(b)
				} else {
					pendingBlocks = pendingBlocks[i:]
					return i, errors.New("Encountered invalid block, aborted connecting")
				}
			}
		} else {
			for _, b := range pendingBlocks {
				bc.ConnectBlock(b)
			}
		}
		connected = len(pendingBlocks)
		pendingBlocks = make([]*blockchain.Block, 0, mempoolCapacity)
		return
	}

	for {
		fmt.Print("demo$> ")
		_, err := fmt.Scan(&cmd)
		if err != nil {
			log.Fatalln("could not read command with error:", err)
		}
		if len(cmd) > 8 {
			fmt.Println("Such long commands are unavailable, try help")
			continue
		}
		switch cmd {

		case "b", "block":
			var block *blockchain.Block
			if readAnswer("Enter hash of previous block manually?") {
				const invalidHashError = "Invalid hash value: use SHA-256 in hexadecimal format"
				var hash string
				fmt.Print("hashPrev: ")
				fmt.Scan(&hash)
				if !regexp.MustCompile(`[[:xdigit:]]{64}`).MatchString(hash) {
					fmt.Println(invalidHashError)
					continue
				}
				block, err = addPendingBlock(strTo32Byte(hash))
				if err != nil {
					fmt.Println(err)
					continue
				}
			} else {
				if len(pendingBlocks) > 0 {
					fmt.Println("There are other pending blocks. If you use last block's hash from blockchain, your block will be invalid")
					if readAnswer("Use last pending block's hash instead?") {
						block, err = addPendingBlock(pendingBlocks[len(pendingBlocks)-1].ID)
						if err != nil {
							fmt.Println(err)
							continue
						}
					} else {
						block, err = addPendingBlock(bc.GetLastBlock().ID)
						if err != nil {
							fmt.Println(err)
							continue
						}
					}
				} else {
					block, err = addPendingBlock(bc.GetLastBlock().ID)
					if err != nil {
						fmt.Println(err)
						continue
					}
				}
			}
			fmt.Print("Block successfully created. ")
			if readAnswer("Print?") {
				block.Print(verbose)
			}

		case "c", "connect":
			if len(pendingBlocks) == 0 {
				fmt.Println("Nothing to connect")
				continue
			}
			c, err := connectPendingBlocks()
			fmt.Printf("Successfully connected %v blocks\n", c)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Hint: you may disable validation with mode command")
			}

		case "h", "help":
			printHelp()

		case "i", "info":
			fmt.Printf(
				"**Blockchain**\nCurrent block height: %v\nValid chain: %v\nTotal tx count: %v\n\n**Mempool**\nPending transactions: %v\nPending blocks: %v\n",
				len(*bc),
				bc.ValidateChain() == -1,
				bc.GetTotalTxCount(),
				len(pendingTxs),
				len(pendingBlocks),
			)

		case "m", "mode":
			fmt.Printf("Validate before connecting: %v\nVerbose block printing: %v\n", validateBeforeConnecting, verbose)
			if readAnswer("Toggle validation?") {
				validateBeforeConnecting = !validateBeforeConnecting
			}
			if readAnswer("Toggle verbose?") {
				verbose = !verbose
			}

		case "p", "print":
			bc.Print(verbose)

		case "q", "quit":
			fmt.Println("Bye.")
			return

		case "t", "tx":
			tx, err := getTxFromUserInput()
			if err != nil {
				fmt.Println("Unable to create transaction: ", err.Error())
				continue
			}
			err = addPendingTx(tx)
			if err != nil {
				fmt.Println("Unable to add transaction to mempool: ", err.Error())
				continue
			}
			fmt.Print("Transaction successfully created. ")
			if readAnswer("Print?") {
				tx.Print()
			}

		case "v", "validate":
			n := bc.ValidateChain()
			if n != -1 {
				fmt.Printf("Chain is invalid on block height %v. ", n+1)
				if readAnswer("Print those blocks?") {
					fmt.Println("Previous block:")
					(*bc)[n].Print(false)
					fmt.Println("Next block with invalid hashPrev:")
					(*bc)[n+1].Print(false)
				}
				continue
			}
			fmt.Println("Chain is valid.")
		case "w", "wipe":
			if readAnswer("Wipe pending transactions?") {
				pendingTxs = make([]*blockchain.Transaction, 0, mempoolCapacity)
			}
			if readAnswer("Wipe pending blocks?") {
				pendingBlocks = make([]*blockchain.Block, 0, mempoolCapacity)
			}

		default:
			fmt.Printf("Command not found: %v. Try help to get available commands\n", cmd)
		}
	}
}

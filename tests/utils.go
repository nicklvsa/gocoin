package tests

import (
	"fmt"
	"gocoin/crypto/block"
	"gocoin/crypto/transaction"
	"gocoin/crypto/user"
	"time"
)

func makeRandomTransactions(amount int) ([]*transaction.Transaction, error) {
	var transacts []*transaction.Transaction

	for i := 0; i < amount; i++ {
		transact, err := transaction.New(
			makeUser("recipient", i),
			makeUser("sender", i+1),
			uint64((i + 1) * 2),
		)
		if err != nil {
			return nil, err
		}

		transacts = append(transacts, transact)
	}

	return transacts, nil
}

func makeUser(role string, idx int) *user.User {
	var email string

	switch role {
	case "sender":
		email = fmt.Sprintf("sender@sender%d.com", idx)
	case "recipient":
		email = fmt.Sprintf("recipient@recipient%d.com", idx)
	case "miner":
		email = fmt.Sprintf("miner@miner%d.com", idx)
	default:
		email = fmt.Sprintf("user@user%d.com", idx)
	}

	return &user.User{
		Email: email,
		Key:   fmt.Sprintf("some-key-%d", idx),
	}
}

func makeBlock(idx int) (*block.Block, error) {
	tr0, _ := transaction.New(
		makeUser("recipient", idx),
		makeUser("sender", idx),
		25,
	)

	tr1, _ := transaction.New(
		makeUser("recipient", idx),
		makeUser("sender", idx),
		30,
	)

	transactions := []*transaction.Transaction{
		tr0, tr1,
	}

	b, err := block.New(transactions, time.Now(), uint64(idx))
	if err != nil {
		return nil, err
	}

	return b, nil
}
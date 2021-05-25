package tests

import (
	"gocoin/crypto/chain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMining(t *testing.T) {
	blockchain, err := chain.New()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.NotNil(t, blockchain)

	miner := makeUser("miner", 0)

	transacts, err := makeRandomTransactions(50)
	if err != nil {
		t.Errorf(err.Error())
	}

	blockchain.PendingTransactions = transacts

	if err := blockchain.MinePending(miner); err != nil {
		t.Errorf(err.Error())
	}
}

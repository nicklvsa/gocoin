package tests

import (
	"encoding/json"
	"gocoin/crypto/chain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	blockchain, err := chain.New()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.NotNil(t, blockchain)

	block0, _ := makeBlock(0)
	block1, _ := makeBlock(1)
	block2, _ := makeBlock(2)

	assert.NotNil(t, block0)
	assert.NotNil(t, block1)
	assert.NotNil(t, block2)

	blockchain.NewBlock(block0)
	blockchain.NewBlock(block1)
	blockchain.NewBlock(block2)

	data, err := json.MarshalIndent(blockchain, " ", "\t")
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Logf("BLOCKCHAIN: %+v\n", string(data))
}

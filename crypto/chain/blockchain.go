package chain

import (
	"fmt"
	"gocoin/crypto/block"
	"gocoin/crypto/transaction"
	"gocoin/crypto/user"
	"time"
)

func New() (*Blockchain, error) {
	c := Blockchain{
		PendingTransactions: []*transaction.Transaction{},
		Chain: []*block.Block{},
		Difficulty: 1000,
		MinerReward: 50.0,
		BlockSize: 50,
	}

	return &c, nil
}

func (c *Blockchain) MinePending(miner *user.User) error {
	pendingLen := len(c.PendingTransactions)

	if pendingLen <= 1 {
		return fmt.Errorf("unable to mine without transactions")
	}

	for i := 0; i < pendingLen; i += c.BlockSize {
		end := i + c.BlockSize
		if i >= pendingLen {
			end = pendingLen
		}

		foundTransactions := c.PendingTransactions[i:end]

		newBlock, err := block.New(foundTransactions, time.Now(), uint64(len(c.Chain)))
		if err != nil {
			return err
		}

		hash := c.LatestBlock().Hash
		newBlock.PreviousHash = &hash
		newBlock.Mine(c.Difficulty)

		c.Chain = append(c.Chain, newBlock)
	}

	return nil
}

func (c *Blockchain) NewBlock(block *block.Block) {
	if c.Chain != nil && len(c.Chain) > 0 {
		block.PreviousHash = &c.LatestBlock().Hash
	}

	c.Chain = append(c.Chain, block)
}

func (c *Blockchain) LatestBlock() *block.Block {
	return c.Chain[len(c.Chain)-1]
}
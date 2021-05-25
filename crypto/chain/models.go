package chain

import (
	"gocoin/crypto/block"
	"gocoin/crypto/transaction"
)

type Blockchain struct {
	Chain []*block.Block `json:"chain"`
	PendingTransactions []*transaction.Transaction `json:"pending_transactions"`
	Difficulty int `json:"difficulty"`
	MinerReward float64 `json:"miner_reward"`
	BlockSize int `json:"block_size"`
}

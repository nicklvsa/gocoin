package chain

import (
	"gocoin/crypto/block"
	"gocoin/crypto/transaction"
	"gocoin/memauth"
	"net/url"
)

type Blockchain struct {
	Auth                *memauth.MemAuth
	Chain               []*block.Block             `json:"chain"`
	Nodes               []*Node                    `json:"nodes"`
	PendingTransactions []*transaction.Transaction `json:"pending_transactions"`
	Difficulty          int                        `json:"difficulty"`
	MinerReward         float64                    `json:"miner_reward"`
	BlockSize           int                        `json:"block_size"`
}

type NodeResponse struct {
	Chain []*block.Block `json:"chain"`
}

type Node struct {
	Address url.URL `json:"address"`
}

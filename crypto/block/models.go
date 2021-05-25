package block

import (
	"gocoin/crypto/transaction"
	"time"
)

type Block struct {
	Idx          uint64                     `json:"index"`
	Hash         string                     `json:"hash"`
	Nonse        int                        `json:"nonse"`
	PreviousHash *string                    `json:"previous_hash"`
	Transactions []*transaction.Transaction `json:"transactions"`
	CreatedAt    time.Time                  `json:"created_at"`
}

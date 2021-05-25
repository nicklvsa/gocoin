package transaction

import (
	"gocoin/crypto/user"
	"time"
)

type Transaction struct {
	Recipient *user.User `json:"recipient"`
	Sender    *user.User `json:"sender"`
	Amount    uint64     `json:"amount"`
	Hash      string     `json:"hash"`
	Signed    bool       `json:"signed"`
	CreatedAt time.Time  `json:"created_at"`
}

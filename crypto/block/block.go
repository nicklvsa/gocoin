package block

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"gocoin/crypto/transaction"
	"gocoin/shared"
	"strconv"
	"time"
)

func New(transactions []*transaction.Transaction, createdAt time.Time, idx uint64) (*Block, error) {
	b := Block{
		PreviousHash: shared.GetPointerToString(""),
		Transactions: transactions,
		CreatedAt: createdAt,
		Idx: idx,
	}

	hash, err := b.GenerateHash()
	if err != nil {
		return nil, err
	}

	b.Hash = hash
	return &b, nil
}

func (b *Block) Mine(diff int) (bool, error) {
	strDiff := strconv.Itoa(diff)
	
	for b.Hash[0:diff] != strDiff {
		b.Nonse++

		hash, err := b.GenerateHash()
		if err != nil {
			return false, err
		}

		b.Hash = hash
	}

	fmt.Printf("New block mined! Proof of work nonse %d", b.Nonse)
	return true, nil
}

func (b *Block) GenerateHash() (string, error) {
	if b.Transactions == nil || len(b.Transactions) <= 0 {
		return "", fmt.Errorf("no transactions occurred on block %d yet", b.Idx)
	}

	var hashes string
	for _, transact := range b.Transactions {
		hashes += transact.Hash
	}

	blockID := fmt.Sprintf("%s%d%s", hashes, b.Idx, b.CreatedAt.String())
	if b.PreviousHash != nil {
		blockID = fmt.Sprintf("%s%s", blockID, *b.PreviousHash)
	}

	hash := hmac.New(sha512.New, []byte("block123"))

	_, err := hash.Write([]byte(blockID))
	if err != nil {
		return "", fmt.Errorf("unable to construct hmac512 hashed block id")
	}

	return string(hash.Sum(nil)), nil
}
package block

import (
	"crypto/sha512"
	"fmt"
	"gocoin/crypto/transaction"
	"strconv"
	"time"
)

func New(transactions []*transaction.Transaction, createdAt time.Time, idx uint64) (*Block, error) {
	b := Block{
		PreviousHash: nil,
		Transactions: transactions,
		CreatedAt:    createdAt,
		Idx:          idx,
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

		fmt.Printf("Looking for new hash %s\n", hash)

		b.Hash = hash
	}

	fmt.Printf("New block mined! Proof of work nonse %d", b.Nonse)
	return true, nil
}

func (b *Block) HasValidTransactions() bool {
	for _, transact := range b.Transactions {
		if !transact.IsValid() {
			return false
		}
	}

	return true
}

func (b *Block) GenerateHash() (string, error) {
	if b.Transactions == nil || len(b.Transactions) <= 0 {
		return "", fmt.Errorf("no transactions occurred on block %d yet", b.Idx)
	}

	var hashes string
	for _, transact := range b.Transactions {
		hashes += transact.Hash
	}

	blockID := fmt.Sprintf("%s%d%s%d", hashes, b.Idx, b.CreatedAt.String(), b.Nonse)
	if b.PreviousHash != nil {
		blockID = fmt.Sprintf("%s%s", blockID, *b.PreviousHash)
	}

	hash := sha512.New()

	_, err := hash.Write([]byte(blockID))
	if err != nil {
		return "", fmt.Errorf("unable to construct hmac512 hashed block id")
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

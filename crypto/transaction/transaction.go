package transaction

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"gocoin/crypto/user"
	"time"
)

func New(recipient, sender *user.User, amount uint64) (*Transaction, error) {
	t := Transaction{
		CreatedAt: time.Now(),
		Recipient: recipient,
		Sender: sender,
		Amount: amount,
	}
	
	hash, err := t.GenerateHash()
	if err != nil {
		return nil, err
	}

	t.Hash = hash
	return &t, nil
}

func (t *Transaction) GenerateHash() (string, error) {
	transactID := fmt.Sprintf("%s%s%d%s", t.Sender.Key, t.Recipient.Key, t.Amount, t.CreatedAt.String())
	hash := hmac.New(sha512.New, []byte("transaction123"))

	_, err := hash.Write([]byte(transactID))
	if err != nil {
		return "", fmt.Errorf("unable to construct hmac512 hashed tranaction id")
	}

	return string(hash.Sum(nil)), nil
}
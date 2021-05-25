package transaction

import (
	"crypto/sha512"
	"fmt"
	"gocoin/crypto/user"
	"time"
)

func New(recipient, sender *user.User, amount uint64) (*Transaction, error) {
	t := Transaction{
		CreatedAt: time.Now(),
		Recipient: recipient,
		Sender:    sender,
		Amount:    amount,
		Signed:    false,
	}

	hash, err := t.GenerateHash()
	if err != nil {
		return nil, err
	}

	t.Hash = hash
	return &t, nil
}

func (t *Transaction) Sign(key, senderKey string) error {
	current, err := t.GenerateHash()
	if err != nil {
		return err
	}

	if t.Hash != current {
		return fmt.Errorf("hash may have been modified! cancelling...")
	}

	if key != senderKey {
		return fmt.Errorf("transaction signature was not valid")
	}

	t.Signed = true
	return nil
}

func (t *Transaction) IsValid() bool {
	current, err := t.GenerateHash()
	if err != nil {
		return false
	}

	if t.Hash != current {
		return false
	}

	if t.Recipient.IsEqual(t.Sender) {
		return false
	}

	if !t.Signed {
		return false
	}

	return true
}

func (t *Transaction) GenerateHash() (string, error) {
	transactID := fmt.Sprintf("%s%s%d%s", t.Sender.UserID.String(), t.Recipient.UserID.String(), t.Amount, t.CreatedAt.String())
	hash := sha512.New()

	_, err := hash.Write([]byte(transactID))
	if err != nil {
		return "", fmt.Errorf("unable to construct sha512 hashed tranaction id")
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

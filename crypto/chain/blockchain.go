package chain

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"gocoin/crypto/block"
	"gocoin/crypto/transaction"
	"gocoin/crypto/user"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gofrs/uuid"
)

func New() (*Blockchain, error) {
	c := Blockchain{
		PendingTransactions: []*transaction.Transaction{},
		Chain:               []*block.Block{},
		Difficulty:          15,
		MinerReward:         50.0,
		BlockSize:           50,
	}

	if err := c.initOriginBlock(); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Blockchain) initOriginBlock() error {
	senderUUID, _ := uuid.NewV4()
	recipientUUID, _ := uuid.NewV4()

	originUserRecipient := user.User{
		Email:  "recipient@recipient.origin-block",
		UserID: recipientUUID,
	}

	originUserSender := user.User{
		Email:  "sender@sender.origin-block",
		UserID: senderUUID,
	}

	transact, err := transaction.New(&originUserRecipient, &originUserSender, 1)
	if err != nil {
		return err
	}

	transactions := []*transaction.Transaction{
		transact,
	}

	b, err := block.New(transactions, time.Now(), 0)
	if err != nil {
		return err
	}

	b.PreviousHash = nil
	c.Chain = append(c.Chain, b)

	return nil
}

func (c *Blockchain) GenerateKeys() (string, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	privKey := x509.MarshalPKCS1PrivateKey(key)
	if err := os.WriteFile("../keys/priv.pem", privKey, 0644); err != nil {
		return "", err
	}

	pubKey, err := x509.MarshalPKIXPublicKey(key.PublicKey)
	if err != nil {
		return "", err
	}
	if err := os.WriteFile("../keys/pub.pem", pubKey, 0644); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", pubKey), nil
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

		latest := c.LatestBlock()
		if latest != nil {
			newBlock.PreviousHash = &latest.Hash
		}

		newBlock.Mine(c.Difficulty)
		fmt.Printf("Mining new block %+v", *newBlock)

		c.Chain = append(c.Chain, newBlock)
	}

	return nil
}

func (c *Blockchain) NewTransaction(sender, recipient *user.User, amount uint64, key, senderKey string) error {
	transact, err := transaction.New(recipient, sender, amount)
	if err != nil {
		return err
	}

	if err := transact.Sign(key, senderKey); err != nil {
		return err
	}

	if !transact.IsValid() {
		return fmt.Errorf("invalid transaction signature")
	}

	c.PendingTransactions = append(c.PendingTransactions, transact)
	return nil
}

func (c *Blockchain) NewNode(address string) error {
	addr, err := url.Parse(address)
	if err != nil {
		return err
	}

	node := Node{
		Address: *addr,
	}

	c.Nodes = append(c.Nodes, &node)
	return nil
}

func (c *Blockchain) ResolveSourceOfTruth() {
	currentMax := len(c.Chain)
	resp := make(chan NodeResponse)

	fetchNode := func(node *Node, response chan<- NodeResponse) {
		client := http.Client{}

		body := map[string]interface{}{
			"last_block": c.LatestBlock(),
		}

		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return
		}

		req, err := http.NewRequest("POST", node.Address.String(), bytes.NewReader(bodyBytes))
		if err != nil {
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			return
		}

		defer resp.Body.Close()

		raw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}

		var nodeResp NodeResponse
		if err := json.Unmarshal(raw, &nodeResp); err != nil {
			return
		}

		response <- nodeResp
	}

	for _, node := range c.Nodes {
		go fetchNode(node, resp)
	}

	go func(chain *Blockchain) {
		for {
			select {
			case data := <-resp:
				if data.Chain != nil {
					if len(data.Chain) > currentMax && c.IsValid() {
						currentMax = len(data.Chain)
						chain.Chain = data.Chain
					}
				}
			}
		}
	}(c)
}

func (c *Blockchain) IsValid() bool {
	for i := 1; i < len(c.Chain); i++ {
		block0 := c.Chain[i-1]
		block1 := c.Chain[i]

		current, err := block1.GenerateHash()
		if err != nil {
			return false
		}

		if !block1.HasValidTransactions() {
			return false
		}

		if block1.Hash != current {
			return false
		}

		if block1.PreviousHash != nil {
			if *block1.PreviousHash != block0.Hash {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func (c *Blockchain) NewBlock(block *block.Block) {
	if c.Chain != nil && len(c.Chain) > 0 {
		block.PreviousHash = &c.LatestBlock().Hash
	}

	c.Chain = append(c.Chain, block)
}

func (c *Blockchain) LatestBlock() *block.Block {
	chainLen := len(c.Chain)

	if chainLen-1 < chainLen {
		return nil
	}

	return c.Chain[len(c.Chain)-1]
}

package blockchain

import (
	"gocoin/shared"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	key := params.ByName("key")
	senderKey := params.ByName("sender_key")
	senderID := params.ByName("sender_id")
	recipientID := params.ByName("recipient_id")
	amount := params.ByName("amount")

	if shared.IsStringPtrNilOrEmpty(&senderID, &recipientID, &amount) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("a sender, recipient, and amount must be provided"))
		return
	}

	amountNum, err := strconv.Atoi(amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid amount passed when creating transaction"))
		return
	}

	sender, err := shared.GlobalChain.Auth.GetUserByID(senderID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to find a user sent to CreateTransaction"))
		return
	}

	recipient, err := shared.GlobalChain.Auth.GetUserByID(recipientID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to find a user sent to CreateTransaction"))
		return
	}

	if err := shared.GlobalChain.NewTransaction(sender, recipient, uint64(amountNum), key, senderKey); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to create new transaction"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully added transaction"))
	return
}

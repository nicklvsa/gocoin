package sync

import (
	"gocoin/shared"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ResolveTruths(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if shared.GlobalChain == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("global block chain has not been initialized"))
		return
	}

	if shared.GlobalChain.ResolveSourceOfTruth() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("successfully resolved conflicting node source"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully resolved no conflicting node sources"))
	return
}

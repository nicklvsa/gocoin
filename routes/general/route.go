package general

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Healthcheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

package router

import (
	"net/http"

	"github.com/user/api/handler"
)

func SetUpRoutes(mux *http.ServeMux) {
	mux.Handle("/hash", handler.PostOnly(handler.ArtificalWait(handler.GetHash())))
}

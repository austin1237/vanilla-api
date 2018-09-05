package router

import (
	"net/http"

	"github.com/user/api/hasher"
	"github.com/user/api/middleware"
	"github.com/user/api/stats"
)

func AddRoutes(mux *http.ServeMux) {
	mux.Handle("/hash", stats.StartTime(middleware.PostOnly(middleware.ArtificalWait(hasher.GenerateHash(stats.CalcDuration())))))
	mux.Handle("/stats", middleware.GetOnly(stats.GetStats()))
}

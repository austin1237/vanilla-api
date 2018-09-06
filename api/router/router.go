package router

import (
	"net/http"

	"github.com/user/api/handler"
	"github.com/user/api/middleware"
	"github.com/user/api/server"
	"github.com/user/api/stats"
)

func CreateRouter(sStats *stats.ServerStats, serv server.Api) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/hash", middleware.StartTime(middleware.PostOnly(middleware.ArtificalWait(handler.Hash(sStats)))))
	mux.Handle("/stats", middleware.GetOnly(handler.Stats(sStats)))
	mux.Handle("/shutdown", handler.ShutDown(serv))
	return mux
}

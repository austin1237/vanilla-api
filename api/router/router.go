package router

import (
	"net/http"

	"github.com/user/api/handler"
	"github.com/user/api/middleware"
	"github.com/user/api/server"
	"github.com/user/api/stats"
)

// CreateRouter will return a new http mux with routes attached
func CreateRouter(sStats *stats.ServerStats, serv server.Api) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/hash", middleware.ReqID(middleware.ReqInfo(middleware.StartTime(middleware.PostOnly(middleware.ArtificalWait(handler.Hash(sStats)))))))
	mux.Handle("/stats", middleware.ReqID(middleware.ReqInfo(middleware.GetOnly(handler.Stats(sStats)))))
	mux.Handle("/shutdown", middleware.ReqID(middleware.ReqInfo(handler.ShutDown(serv))))
	return mux
}

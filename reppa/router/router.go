package router

import (
	"fmt"
	"github.com/beldpro-ci/reppa/reppa/common"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var log = common.GetLogger()

func logRequest(r *http.Request) {
	log.Info("incoming request",
		zap.String("user-agent", r.UserAgent()),
		zap.String("host-header", r.Host),
		zap.String("uri", r.RequestURI),
		zap.String("method", r.Method),
		zap.String("time", time.Now().Format("02/Jan/2006:15:04:05 -0700")))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logRequest(r)
	fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)

	return router
}

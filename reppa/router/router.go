package router

import (
	"fmt"
	"github.com/beldpro-ci/reppa/reppa/common"
	"github.com/beldpro-ci/reppa/reppa/core/git"
	"github.com/davecgh/go-spew/spew"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var log = common.GetLogger()

type RouterConfig struct {
	GitConfig *git.GitConfig
}

func New(cfg *RouterConfig) *httprouter.Router {
	router := httprouter.New()
	router.GET("/ping", ping)
	router.GET("/repositories/:name", newRepoHandler(cfg.GitConfig))

	log.Debug("Router initialized",
		zap.String("cfg", spew.Sprintf("%#v", cfg)))

	return router
}

func logRequest(r *http.Request) {
	log.Info("incoming request",
		zap.String("user-agent", r.UserAgent()),
		zap.String("host-header", r.Host),
		zap.String("uri", r.RequestURI),
		zap.String("method", r.Method),
		zap.String("time", time.Now().
			Format("02/Jan/2006:15:04:05 -0700")))
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logRequest(r)
	fmt.Fprint(w, "PONG")
}

func newRepoHandler(cfg *git.GitConfig) func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	gm, err := git.New(cfg)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logRequest(r)
		var repo = ps.ByName("name")

		if err := gm.InitBareRepository(repo); err != nil {
			log.Error("Errored initializing bare repo",
				zap.Error(err))
		}

		fmt.Fprintf(w, "repo=%s!\n", repo)
	}
}

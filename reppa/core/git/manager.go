package git

import (
	"github.com/beldpro-ci/reppa/reppa/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os/exec"
	"path"
)

var log = common.GetLogger()

// GitConfig presents the basic configuration to be used by GitManager.
type GitConfig struct {
	// RepositoriesRoot corresponds to the absolute path to where git
	// repositories live.
	RepositoriesRoot string

	// GitBinary (optional) is the location of the git binary to use when
	// git executing commands. If not specified, `git` is looked up in
	// PATH. An error occurs if nothing found.
	GitBinary string
}

type GitManager struct {
	logger *zap.Logger
	cfg    *GitConfig
	git    *Git
}

// New instantiates a new GitManager. It requires that the configuration has
// at least `RepositoriesRoot` set.
func New(cfg *GitConfig) (gm *GitManager, err error) {
	if cfg.RepositoriesRoot == "" {
		return nil, errors.New(
			"No git root specified")
	}

	if cfg.GitBinary == "" {
		gitLocation, err := exec.LookPath("git")
		if err != nil {
			return nil, errors.Wrapf(err,
				"Couldn't search for `git` in $PATH")
		}

		cfg.GitBinary = gitLocation
	}

	gm = &GitManager{
		cfg: cfg,
		logger: log.With(
			zap.String("cfg", spew.Sprintf("%#v", cfg))),
		git: &Git{Binary: cfg.GitBinary},
	}

	gm.logger.Debug("GitManager initialized")

	return gm, nil
}

const BARE_REPO = true

func (gm *GitManager) InitBareRepository(name string) error {
	var repositoryLocation = path.Join(gm.cfg.RepositoriesRoot, name)

	gm.logger.Debug("Initializing bare repo",
		zap.String("name", name),
		zap.String("location", repositoryLocation))

	if err := gm.git.InitBare(repositoryLocation); err != nil {
		return errors.Wrapf(err,
			"Errored initializing bare repo [%s]",
			repositoryLocation)
	}

	return nil
}

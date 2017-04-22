package git

import (
	"github.com/pkg/errors"
	"os"
	"os/exec"
)

type Git struct {
	// Binary is where `git` is present.
	Binary string
}

// InitBare initializes a given `RepositoryLocation` as a bare git repo.
// If the location doesn't exists, creates it.
// It expects `repositoryLocation` to be an absolute path.
func (g *Git) InitBare(repositoryLocation string) error {
	_, err := os.Stat(repositoryLocation)
	if err != nil {
		if err != os.ErrNotExist {
			return errors.Wrapf(err,
				"Couldn't inspect repo location for bare init [%s]",
				repositoryLocation)
		}

		mkdirErr := os.MkdirAll(repositoryLocation, 0755)
		if mkdirErr != nil {
			return errors.Wrapf(mkdirErr,
				"Couldn't create directory while initializing bare repo [%s]",
				repositoryLocation)
		}
	}

	err = g.execCmd(g.Binary, "-C", repositoryLocation, "init", "--bare")
	if err != nil {
		return errors.Wrapf(err,
			"Execution of `init bare` command failed.")
	}

	return nil
}

func (g *Git) execCmd(args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrapf(err,
			"Command execution failed [%v]. output=[%s]",
			cmd, out)
	}

	return nil
}

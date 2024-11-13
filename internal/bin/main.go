package bin

import (
	"embed"
	"os"
	"path/filepath"

	e "github.com/Chaitanyabsprip/dot/internal/core/embed"
	"github.com/rwxrob/bonzai/comp"
	"github.com/rwxrob/bonzai"
)

//go:embed bin
var embedFs embed.FS

var Cmd = &bonzai.Cmd{
	Name:  `bin`,
	Short: `bin is a utility to manage github-cli configuration`,
	Comp:  comp.Cmds,
	Cmds:  []*bonzai.Cmd{setupCmd},
}

var setupCmd = &bonzai.Cmd{
	Name:  `setup`,
	Opts:  `slim|quik|full`,
	Short: `Setup bin directory`,
	Comp:  comp.Opts,
	Do: func(x *bonzai.Cmd, args ...string) error {
		binDir := filepath.Join(os.Getenv("HOME"), ".local")
		return e.SetupAll(embedFs, "bin", binDir, nil)
	},
}

package tmux

import (
	"embed"
	e "github.com/Chaitanyabsprip/dot/internal/core/embed"

	"github.com/rwxrob/bonzai/comp"
	"github.com/rwxrob/bonzai"

	"github.com/Chaitanyabsprip/dot/internal/core/oscfg"
)

// TODO(me):
// - dependencies
// - live reload

// TODO(me):
// - Create options to control what level of setup is done
// - The options will be one of slim/quik/full
// - Slim: install only the files that are needed and skip the rest
// - Quik: install only the configuration files and skip the
//   optional dependencies
// - Full: install everything, all config files and all dependencies

//go:embed all:tmux
var embedFs embed.FS

var Cmd = &bonzai.Cmd{
	Name:  `tmx`,
	Usage: `tmx <command>`,
	Short: `tmx is a utility to manage tmux configuration and related scripts`,
	Comp:  comp.Cmds,
	Cmds:  []*bonzai.Cmd{setupCmd, runCmd},
}

var setupCmd = &bonzai.Cmd{
	Name:  `setup`,
	Usage: `tmux setup`,
	Opts:  `slim|full`,
	Short: `Setup tmux copies configuration files to config directory`,
	Long:  ``,
	Call: func(x *bonzai.Cmd, _ ...string) error {
		return e.SetupAll(embedFs, "tmux", oscfg.ConfigDir())
	},
}

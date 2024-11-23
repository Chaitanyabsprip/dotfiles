package fish

import (
	"embed"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/comp"

	e "github.com/Chaitanyabsprip/dotfiles/internal/core/embed"

	"github.com/Chaitanyabsprip/dotfiles/internal/core/oscfg"
)

//go:embed fish
var embedFs embed.FS

var Cmd = &bonzai.Cmd{
	Name:  `fish`,
	Short: `fish is a utility to manage fish configuration`,
	Comp:  comp.Cmds,
	Cmds:  []*bonzai.Cmd{setupCmd},
}

var setupCmd = &bonzai.Cmd{
	Name:  `setup`,
	Opts:  `slim|quik|full`,
	Short: `setup fish`,
	Comp:  comp.Opts,
	Do: func(x *bonzai.Cmd, args ...string) error {
		return e.SetupAll(embedFs, "fish", oscfg.ConfigDir(), nil)
	},
}

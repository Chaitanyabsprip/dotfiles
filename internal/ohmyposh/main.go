package ohmyposh

import (
	"embed"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/comp"

	e "github.com/Chaitanyabsprip/dotfiles/internal/core/embed"

	"github.com/Chaitanyabsprip/dotfiles/internal/core/oscfg"
)

//go:embed oh-my-posh.rc.toml
var embedFs embed.FS

var Cmd = &bonzai.Cmd{
	Name:  `ohmyposh`,
	Short: `manage ohmyposh configuration`,
	Comp:  comp.Cmds,
	Cmds:  []*bonzai.Cmd{setupCmd},
}

var setupCmd = &bonzai.Cmd{
	Name:  `setup`,
	Opts:  `slim|quik|full`,
	Short: `setup ohmyposh`,
	Comp:  comp.Opts,
	Do: func(x *bonzai.Cmd, args ...string) error {
		return e.SetupAll(embedFs, "ohmyposh", oscfg.ConfigDir(), nil)
	},
}
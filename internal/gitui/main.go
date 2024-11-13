package gitui

import (
	"embed"

	e "github.com/Chaitanyabsprip/dot/internal/core/embed"
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/comp"

	"github.com/Chaitanyabsprip/dot/internal/core/oscfg"
)

//go:embed gitui
var embedFs embed.FS

var Cmd = &bonzai.Cmd{
	Name:  `gitui`,
	Short: `gitui is a utility to manage gitui configuration`,
	Comp:  comp.Cmds,
	Cmds:  []*bonzai.Cmd{setupCmd},
}

var setupCmd = &bonzai.Cmd{
	Name:  `setup`,
	Opts:  `slim|quik|full`,
	Short: `Setup gitui`,
	Comp:  comp.Opts,
	Do: func(x *bonzai.Cmd, args ...string) error {
		return e.SetupAll(embedFs, "gitui", oscfg.ConfigDir(), nil)
	},
}

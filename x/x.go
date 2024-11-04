package x

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/comp"

	"github.com/Chaitanyabsprip/dot/x/colors"
	"github.com/Chaitanyabsprip/dot/x/distro"
	"github.com/Chaitanyabsprip/dot/x/have"
	"github.com/Chaitanyabsprip/dot/x/install"
	"github.com/Chaitanyabsprip/dot/x/workdirs"
)

var Cmd = &bonzai.Cmd{
	Name:  `x`,
	Usage: `x <command>`,
	Short: `x is a utility to manage scripts`,
	Comp:  comp.Cmds,
	Cmds: []*bonzai.Cmd{
		colors.Cmd,
		creashCmd,
		distro.Cmd,
		have.Cmd,
		install.Cmd,
		workdirs.Cmd,
	},
}

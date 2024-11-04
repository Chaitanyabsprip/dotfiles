package dot

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/comp"

	"github.com/Chaitanyabsprip/dot/internal/alacritty"
	"github.com/Chaitanyabsprip/dot/internal/bat"
	"github.com/Chaitanyabsprip/dot/internal/bin"
	"github.com/Chaitanyabsprip/dot/internal/gh"
	"github.com/Chaitanyabsprip/dot/internal/git"
	"github.com/Chaitanyabsprip/dot/internal/gitui"
	"github.com/Chaitanyabsprip/dot/internal/kitty"
	"github.com/Chaitanyabsprip/dot/internal/tmux"
	"github.com/Chaitanyabsprip/dot/x"
)

var Cmd = &bonzai.Cmd{
	Name:  `dot`,
	Alias: `d`,
	Usage: `dotfiles`,
	Short: `Manage dotfiles`,
	Comp:  comp.Cmds,
	Cmds: []*bonzai.Cmd{
		alacritty.Cmd,
		bat.Cmd,
		bin.Cmd,
		gh.Cmd,
		git.Cmd,
		gitui.Cmd,
		kitty.Cmd,
		tmux.Cmd,
		x.Cmd,
	},
}

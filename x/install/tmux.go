package install

import (
	"fmt"
	"os"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/run"

	"github.com/Chaitanyabsprip/dotfiles/x/distro"
	"github.com/Chaitanyabsprip/dotfiles/x/have"
)

var TmuxCmd = &bonzai.Cmd{
	Name:  `tmux`,
	Alias: `i`,
	Do:    func(_ *bonzai.Cmd, _ ...string) error { return Tmux() },
}

func Tmux() error {
	if ok, _ := have.Executable(`tmux`); ok {
		fmt.Println(`tmux is already installed`)
		return nil
	}
	switch distro.Name() {
	case `Arch Linux`:
		return WithRoot(`pacman`, `-S`, `tmux`)
	case `Ubuntu`, `Debian GNU/Linux`:
		return WithRoot(`apt-get`, `install`, `-y`, `tmux`)
	case `Fedora Linux`:
		return run.Exec(`dnf`, `install`, `tmux`, `-y`)
	case `Darwin`:
		return run.Exec(`brew`, `install`, `tmux`)
	default:
		fmt.Fprintln(
			os.Stderr,
			`Unsupported operating system. Please install tmux manually.`,
		)
	}
	return nil
}
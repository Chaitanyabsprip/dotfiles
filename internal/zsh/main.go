package zsh

import (
	"embed"
	"os"
	"path/filepath"
	"runtime"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/bonzai/cmds/help"
	"github.com/rwxrob/bonzai/comp"
	"github.com/rwxrob/bonzai/run"

	e "github.com/Chaitanyabsprip/dotfiles/internal/core/embed"
	"github.com/Chaitanyabsprip/dotfiles/internal/ohmyposh"
	"github.com/Chaitanyabsprip/dotfiles/internal/shell"
	"github.com/Chaitanyabsprip/dotfiles/x/install"

	"github.com/Chaitanyabsprip/dotfiles/internal/core/oscfg"
)

//go:embed all:zsh
var embedFs embed.FS

var Cmd = &bonzai.Cmd{
	Name:  `zsh`,
	Short: `zsh is a utility to manage zsh configuration`,
	Comp:  comp.Cmds,
	Cmds: []*bonzai.Cmd{
		initCmd,
		setupCmd,
		install.ZshCmd.WithName(`install`),
	},
}

var initCmd = &bonzai.Cmd{
	Name:  `init`,
	Opts:  `slim|quik|full`,
	Short: `setup zsh to a specific level of configuration`,
	Long: `
The "setup" command installs and configures Zsh with various options 
that determine the extent of the setup. Zsh is installed for all forms
options.

The primary setup modes are:

- slim: Bare minimum level of personal configuration. Enough to get
  setup instantly without any extra installs. Installs only static
  configuration files that are required.
- quik: Enough configuration to get a familiar environment while doing 
  minimum number of third party installs. Install plugins and drops into
  zsh shell.
- full: Provides a complete Zsh setup, installing the "zap" plugin 
  manager in addition to the "zshenv" configuration file and the Zsh 
  shell. This is suitable for users who want a fully featured Zsh 
  environment with plugin management capabilities.

When no option is specified, the "setup" command defaults to "slim"
mode.

# ENVIRONMENT:
  - LAUNCH: Set to a truthy value to launch Zsh after setup.
`,
	MaxArgs: 1,
	Comp:    comp.Opts,
	Cmds:    []*bonzai.Cmd{help.Cmd},
	Do: func(x *bonzai.Cmd, args ...string) (err error) {
		zshenvPath := filepath.Join(os.Getenv(`HOME`), `.zshenv`)
		if len(args) == 0 {
			args = append(args, `slim`)
		}
		overrides := map[string]string{
			`zsh/.zshenv`:           zshenvPath,
			`zsh/conf.d/brew.sh`:    ``,
			`zsh/conf.d/cdpath.off`: ``,
			`zsh/conf.d/clone.sh`:   ``,
			`zsh/conf.d/eza.sh`:     ``,
			`zsh/conf.d/fzf.sh`:     ``,
			`zsh/conf.d/jira.sh`:    ``,
			`zsh/conf.d/nvm.sh`:     ``,
			`zsh/conf.d/opts.sh`:    ``,
			`zsh/conf.d/pipx.sh`:    ``,
			`zsh/conf.d/prompt.sh`:  ``,
			`zsh/conf.d/pyenv.sh`:   ``,
		}
		mode := args[0]
		if mode == `slim` || mode == `quik` || mode == `full` {
			err := install.ZshCmd.Run()
			if err != nil {
				return err
			}
		}
		if mode == `quik` || mode == `full` {
			delete(overrides, `zsh/conf.d/clone.sh`)
			delete(overrides, `zsh/conf.d/prompt.sh`)
		}
		if mode == `full` {
			if err := shell.Cmd.Run(`setup`); err != nil {
				return err
			}
			if err := ohmyposh.Cmd.Run(`setup`); err != nil {
				return err
			}
			if err := install.OhMyPosh(); err != nil {
				return err
			}
			defer func() { err = install.Zap() }()
			delete(overrides, `zsh/conf.d/brew.sh`)
			delete(overrides, `zsh/conf.d/cdpath.disable`)
			delete(overrides, `zsh/conf.d/completions.sh`)
			delete(overrides, `zsh/conf.d/eza.sh`)
			delete(overrides, `zsh/conf.d/fzf.sh`)
			delete(overrides, `zsh/conf.d/jira.sh`)
			delete(overrides, `zsh/conf.d/nvm.sh`)
			delete(overrides, `zsh/conf.d/pipx.sh`)
			delete(overrides, `zsh/conf.d/pyenv.sh`)
			if runtime.GOOS == `darwin` {
				delete(overrides, `zsh/conf.d/brew.sh`)
			}
		}
		err = e.SetupAll(
			embedFs,
			`zsh`,
			oscfg.ConfigDir(),
			overrides,
		)
		if err != nil {
			return err
		}
		if os.Getenv(`LAUNCH`) != `` {
			return run.Exec(`zsh`, `-l`)
		}
		return nil
	},
}

var setupCmd = &bonzai.Cmd{
	Name:  `setup`,
	Alias: `conf`,
	Do: func(_ *bonzai.Cmd, _ ...string) error {
		zshenvPath := filepath.Join(os.Getenv(`HOME`), `.zshenv`)
		overrides := map[string]string{
			`zsh/.zshenv`: zshenvPath,
		}
		err := e.SetupAll(
			embedFs,
			`zsh`,
			oscfg.ConfigDir(),
			overrides,
		)
		return err
	},
}
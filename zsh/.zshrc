#!/bin/zsh

HISTORY_IGNORE="(la|ls|cd|pwd|exit|history|cd -|cd ..|q|c|vim|nvim)"
SAVEHIST=200000
HISTFILE=~/.config/zsh/.zsh-history
HISTFILESIZE=1000000000
HISTSIZE=1000000000

have starship && eval "$(starship init zsh)"

# Cannot use autocd option along with CDPATH
# setopt autocd                 # Automatically cd into typed directory.
setopt hist_expire_dups_first # delete duplicates first when HISTFILE size exceeds HISTSIZE
setopt hist_find_no_dups      # ignore dups when going through history
setopt hist_ignore_dups       # ignore duplicated commands history list
setopt hist_ignore_all_dups   # Delete old recorded entry if new entry is a duplicate
setopt hist_ignore_space      # ignore commands that start with space
setopt hist_verify            # show command with history expansion to user before running it
setopt inc_append_history     # add commands to HISTFILE in order of execution
stty stop undef               # Disable ctrl-s to freeze terminal.

set -o vi
bindkey -v

# Load seperated config files
plugin_file=$HOME/.config/zsh/plugs.zsh
config_dir=$HOME/.config/zsh/conf.d

[ -r $plugin_file ] && source $plugin_file

for conf in "$config_dir/"*.sh; do
	[ -r "$conf" ] && zsh-defer source "$conf"
done
unset conf

# key-bindings
autoload -Uz edit-command-line
zle -N edit-command-line

# initialisations
have brew && zsh-defer eval "$(brew shellenv | grep -v 'export PATH')"
have rbenv && zsh-defer eval "$(rbenv init - zsh)"

# function cd() {
# 	builtin cd "$@"
# 	if [ -d '.venv' ] && [ -f '.venv/bin/activate' ]; then
# 		source .venv/bin/activate
# 	fi
# }

# zsh-defer unfunction _have

# [[ $TMUX ]] || tmux new -As home

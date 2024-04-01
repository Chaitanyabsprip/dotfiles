#!/bin/sh

_have() { type "$1" >/dev/null 2>&1; }

# Only to be used as an alias to disable commands
donothing() { echo "$*" >/dev/null 2>&1; }

# _have nvim && alias nvim='donothing'
# alias clear='donothing'
# alias exit='donothing'

alias mkdir='mkdir -pv'
alias view='vi -R'

_have note && alias bm='note -b'
_have note && alias did='note -c'
_have note && alias todo='note -t'

_have /usr/bin/vim && alias vi=/usr/bin/vim
_have /usr/local/bin/nvim && alias vim=/usr/local/bin/nvim vimdiff='vim -d'
_have bat && alias cat='bat'

md() { [ -z "$1" ] && exit 1 || mkdir "$1" && cd "$1" || return; }

_have fd && _have fzf && edit() {
	fd . -Ht f -d 6 -E .git -E .DS_Store |
		fzf --height=20% --border --margin=0 --padding=0 \
			--bind='enter:become(nvim {})'
}

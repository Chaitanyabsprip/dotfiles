#!/bin/zsh

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

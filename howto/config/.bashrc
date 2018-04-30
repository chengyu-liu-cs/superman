############ Alias
alias ll="ls -la --color=auto"
alias ls="ls --color=auto"
alias c='clear'
alias ..='cd ..'
alias ...='cd ../..'

alias grep='grep --color=auto'
alias egrep='egrep --color=auto'
alias fgrep='fgrep --color=auto'

alias now='date +"%T"'
alias nowtime=now
alias nowdate='date +"%d-%m-%Y"'

alias ports='netstat -tulanp'

alias df="df -Tha --total"
alias du="du -ach | sort -h"
alias ps="ps auxf"
alias psg="ps aux | grep -v grep | grep -i -e VSZ -e"
alias histg="history | grep"

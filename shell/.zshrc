export PATH="/usr/local/opt/bzip2/bin:$PATH"
export LDFLAGS="-L/usr/local/opt/bzip2/lib"
export CPPFLAGS="-I/usr/local/opt/bzip2/include"
eval "$(pyenv init -)"
export PATH="/usr/local/opt/openjdk@11/bin:$PATH"
export PATH="/usr/local/sbin:$PATH"
export JAVA_HOME='/usr/local/Cellar/openjdk@11/11.0.12/libexec/openjdk.jdk/Contents/Home'


alias dkb='docker build'
alias dki='docker images'
alias dkp='docker ps -a'
alias dkr='docker run'
alias dksp='docker system prune -af'


alias gta='git add .'
alias gtc='git commit -m'
alias gts='git status'
alias gtk='git checkout'
alias gtb='git branch'
alias gtp='git pull'
alias gtpo='git pull origin'
alias gtl='git log'
alias gtrh='git reset --soft head'
alias gtdh='git diff head'

alias ll='ls -al'
export PATH="/usr/local/sbin:$PATH"


# secrets
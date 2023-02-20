# export PATH="/usr/local/opt/bzip2/bin:$PATH"
# export LDFLAGS="-L/usr/local/opt/bzip2/lib"
# export CPPFLAGS="-I/usr/local/opt/bzip2/include"
# eval "$(pyenv init -)"
# export PATH="/usr/local/opt/openjdk@11/bin:$PATH"
# export PATH="/usr/local/sbin:$PATH"
# export PATH="/Users/song/.local/bin:$PATH"
# export JAVA_HOME='/usr/local/Cellar/openjdk@11/11.0.12/libexec/openjdk.jdk/Contents/Home'

# chrome
# export chrome='open -na Google\ Chrome --args --user-data-dir=/tmp/temporary-chrome-profile-dir --disable-web-security --disable-site-isolation-trials'

# node
# npm config set registry https://registry.npm.taobao.org/
export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.aliyun.com/homebrew/homebrew-bottles
export PATH="/usr/local/opt/node@18/bin:$PATH"
export LDFLAGS="-L/usr/local/opt/node@18/lib"
export CPPFLAGS="-I/usr/local/opt/node@18/include"

#python
# alias python3.8='python3'
alias python='python3.6'
alias python3='python3.6'
alias pip='pip3'

#golang
export GOPROXY='https://goproxy.cn/,direct'
export GOPATH="$HOME/go"
export PATH="$GOPATH/bin:$PATH"
export GO111MODULE=on
# go env -w GO111MODULE=on

# docker
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
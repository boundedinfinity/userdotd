if command -v go >/dev/null 2>&1 ; then
    export GOPATH=$HOME/go
    export PATH="$PATH:$GOPATH/bin"
fi

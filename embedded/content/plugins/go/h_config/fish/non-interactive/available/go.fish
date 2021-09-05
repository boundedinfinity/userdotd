if type -q go
    set -gx GOPATH $HOME/go
    set -gx PATH $PATH $GOPATH/bin
end

for i in $HOME/.config/bash/completions/*.bash; do
    if test -r $i
        source $i
    fi
done

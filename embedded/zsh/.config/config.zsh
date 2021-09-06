for i in $HOME/.config/zsh/non-interactive/enabled/*.zsh
    if test -r $i
        source $i
    fi
done

if [[ -o interactive ]]; then
    for i in $HOME/.config/zsh/interactive/enabled/*.zsh; do
        if test -r $i
            source $i
        fi
    done
fi

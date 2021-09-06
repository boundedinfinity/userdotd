for i in $HOME/.config/bash/non-interactive/enabled/*.bash; do
    if test -r $i
        source $i
    fi
done

case "$-" in
*i*)
    for i in $HOME/.config/bash/interactive/enabled/*.bash; do
        if test -r $i
            source $i
        fi
    done
;;
*)
;;
esac

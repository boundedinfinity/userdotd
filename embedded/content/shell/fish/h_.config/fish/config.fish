for i in $HOME/.config/fish/non-interactive/enabled/*.fish
    if test -r $i
        source $i
    end
end

if status --is-interactive
    for i in $HOME/.config/fish/interactive/enabled/*.fish
        if test -r $i
            source $i
        end
    end
end

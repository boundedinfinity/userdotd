for i in $HOME/.config/userdotd/fish/login.d/enabled/*.fish
    if test -r $i
        source $i
    end
end

# function userdotd_exit --on-process-exit %self
#     for i in $HOME/.config/userdotd/fish/logout.d/enabled/*.fish
#         if test -r $i
#             source $i
#         end
#     end
# end

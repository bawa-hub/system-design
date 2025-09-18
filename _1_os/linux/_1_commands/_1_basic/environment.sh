# shell stores two basic types of data in the environment:
# environment variables and shell variables

# Shell variables are bits of data placed there by bash, and
# environment variables are everything else

# shell stores some programmatic data, namely, aliases and shell functions.

# list of environment variables
printenv
printenv | less # if large amount of data
printenv USER # list value of variable
set | less # display both the shell and environment variables
echo $HOME # show value of variable

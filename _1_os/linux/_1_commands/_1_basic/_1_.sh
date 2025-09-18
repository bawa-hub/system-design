# print working directory
pwd 

# change directory
cd /var  
cd ~  # home directory
cd ~bawa #home dir of user bawa

# list directories
ls 
ls /path/to/any/folder
ls -l # more info
ls -a #show hidden files
ls -la #show more info and hidden files also
ls ~ /usr #show user home directory and /usr directory simultaneouly

# create file
touch <filename>

# description of file content
file <filename>

# show file data (for less content)
cat <filename>

# show large file data
less /path/to/file
man less # know more about less

# previously executed commands
history
ctrl+R write_part_of_command_ ctrl+R # for searching commands 

# clear terminal screen
clear

# copy 
cp <source_file_path> <destination_file_path>

# move 
mv <source_file_path> <destination_file_path>

# make directory
mkdir

# remove file
rm <file_path>

# remove directory 
rm -f <dir> # forces remove
rm -r <dir> # recursively remove
rmdir <dir>

# find file in system
find /path/to/dir -name <filename>
find /path/to/dir -type d -name <folder>  # d = directory

# help for commands
<command> --help

# detail manual of command
man <command>

# get help about anything
apropos <anything>

# what a command does
whatis <command>

# know the path of command binaries
which <command>

# make alias of command
# alias any_name='any_command'
alias foobar='la -la' # alias will remove after reboot
# add in ~/.bashrc to add alias permanently
unalias <alias_name> # remove alias

# close terminal
exit
logout

# current username
whoami

# logged in users details
who 
w # for detailed details

# all about user
id

# hostname
hostname

# about linux
uname <options>

# session stuff
ss

# environment variables
env

# show all shells present in system
cat /etc/shells 

# current amount of free space on our disk drives
df

# amount of free memory
free

# prompt string
echo $PS1
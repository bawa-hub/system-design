# stdout (standard out)
echo hello world > hello.txt
# > (redirection operator)
# send ouput to a file not on screen
# > will overwrite the output
echo hello there >> hello.txt
# >> will append the output


# stdin (standard in)
# we have stdin from devices like keyboard, 
# files, o/p from other process and terminal 
cat < hello.txt > hi.txt
# takes input from hello.txt and redirect to hi.txt
# < stdin redirection


# pipe and tee
ls -la /etc | less
# | get the stdout of a command and make that the stdin to another process
ls | tee output.txt
# tee write output to two different streams 


# env (environment)
env # show list of environment variables
echo $PATH # path variables where system searches for binaries


# stderr (Standard error)
# cut
# paste


# head
head /var/log/syslog # show first 10 lines in this file
head -n 100 /var/log/syslog # show 100 lines


# tail
tail /var/log/syslog # show last 10 lines in this file
tail -n 100 /var/log/syslog # show 100 lines
tail -f /var/log/syslog # can see changing things in the file


# expand and unexpand
# join and split
# sort
# tr (translate)
# uniq (Unique)
# wc and ni


# grep
# search files for characters that matches a certain pattern
grep hello hello.txt
grep -i somepattern somefile # grep pattern 
env | grep -i User
# can use regex
ls /somedir | grep '.txt$'  # return all files ending with .txt in somedir

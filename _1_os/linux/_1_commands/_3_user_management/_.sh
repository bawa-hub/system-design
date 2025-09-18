# current user 
whoami

# Create low priviledge user --
sudo  adduser <name> # did many configuration for us

useradd -h
sudo useradd <name> # lazy, didn't do any configuration
sudo passwd <username> # add password
sudo usermod <username> --shell /bin/bash # change default shelll
sudo usermod -l <old_username> <new_username> # change username


# list all users
cat /etc/passwd

# password of all users
sudo cat /etc/shadow

# modify user
usermod -h # all commands options help

# switch user
switch - <username>  # switch to username
switch - # switch to root

# edit sudoers file
sudo visudo

# delete user
sudo userdel <username>

# add group 
sudo groupadd <groupname>

# list of groups
cat /etc/group

# groups of user
groups

# add user to sudoers (define who can use sudo) group --
sudo usermod -aG sudo <username>

# add user to created group
sudo usermod -aG <groupname> <username> 

# remove user from group
sudo gpasswd -d <username> <groupname>

# delete group
sudo groupdel <groupname>
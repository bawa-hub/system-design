# processes - instance of a running program

# list of running processes:
ps
# PID: Process ID
# TTY: Controlling terminal associated with the process (we'll go in detail about this later)
# STAT: Process status code
# TIME: Total CPU usage time
# CMD: Name of executable/command
ps -aux


# daemons - processes that we don't start (background process)
# daemon process has d letter at end. eg: sshd

# systemd (master daemon/boss of the daemons) --
# it has two main jobs:
# service manager, initialization system (init)

# boot -> kernel -> systemd -> mounting file system, starting all services

# Note: systemd is not the only init system out there
# In past there are also others like: sysV init, upstart ...

# all process branching
pstree

# service manager (control all of our daemons)
systemctl

sudo systemctl start/stop/restart/reload-or-restart/status <service_name>

# daemons starts by itself on boot or not start at boot time:
# sudo systemctl disable <name>
sudo systemctl disable ntp # ntp - network time protocol
# sudo systemctl enable <name>
# check daemon is active or not
sudo systemctl is-active ntp
sudo systemctl is-enable ntp

# list all active daemons/services
sudo systemctl list-units   # systemd thinks all daemons as units
sudo systemctl list-units -t service # only service type will show
# search for particular service (eg nginx)
sudo systemctl list-units | grep nginx 
sudo systemctl list-unit-files | grep nginx # if nginx is disabled


# process run by user
ps -u <username>

# kill process
pgrep firefox # ps+grep = pgrep , get firefox pid
kill <pid>


# process details
top
htop # show detailed details
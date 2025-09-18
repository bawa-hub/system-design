# login to database 
sudo -i -u postgres
psql

# change password
\password <username>

# Connect to PostgreSQL database
psql -d <db_name> -U  <user_name> -W
psql -h <host_address> -d <db_name> -U  <user_name> -W # database that resides on another host
psql -U <user_name> -h <host_addr> "dbname=db sslmode=require" # to use SSL mode for the connection

# Switch connection to a new database
\c <db_name> <user_name> # omit username if want to use for same user

# List available databases
\l

# List available tables
\dt

# Describe a table
\d table_name

# List available schema
\dn

# List available functions
\df

# List available views
\dv

# List users and their roles
\du

# Execute the previous command
\g

# Command history
\s

# save the command history to a file
\s <filename>

# Execute psql commands from a file
\i <filename>

# Get help on psql commands
\? # To know all available psql commands
\? <command_name> # for specific command help
\? ALTER TABLE

# Turn on and off query execution time
\timing

# Edit command in your own editor
\e # fire an editor to write commands
\ef # for functions

# Switch output options
\a # switches from aligned to non-aligned column output.
\H # formats the output to HTML format.

# Quit psql
\q
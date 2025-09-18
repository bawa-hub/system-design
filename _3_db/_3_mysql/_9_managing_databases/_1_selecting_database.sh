# login to mysql
mysql -u root -p

# list of databases
show databases;

# use database
use database_name;

# show current database
select database();

# login with db selected
mysql -u root -D db_name -p
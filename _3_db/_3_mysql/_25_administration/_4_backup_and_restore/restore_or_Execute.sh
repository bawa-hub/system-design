mysql -u root -p
use <db_name>;


# dump/backup of database
# https://www.sqlshack.com/how-to-backup-and-restore-mysql-databases-using-the-mysqldump-command/
 mysqldump -u root -p cedarwood_prod12jan23 > cedardump.sql

# restore or excecute any file
mysql -u root -p sakila < C:\MySQLBackup\sakila_20200424.sql
# or
source <sql_file>
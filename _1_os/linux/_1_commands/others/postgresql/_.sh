sudo apt update
sudo apt install postgresql

# login to postgresql as user "postgres"
sudo -i -u postgres

# restore data from tar file
pg_restore --dbname=<db_name> --verbose <filename.tar>

# enter into shell
psql
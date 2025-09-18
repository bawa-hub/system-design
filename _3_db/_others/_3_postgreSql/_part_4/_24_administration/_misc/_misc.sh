# docker postgres
docker run -e POSTGRES_PASSWORD=postgres --name pg postgres # pull image
docker exec -it pg psql -U postgres # run

# install on mac
download: https://www.geeksforgeeks.org/install-postgresql-on-mac/
save below line in ~/.zshrc
export PATH=/Library/PostgreSQL/15/bin:$PATH
psql -U postgres


# postgres psql commands
https://www.geeksforgeeks.org/postgresql-psql-commands/

# restore database
pg_restore -U postgres -d dvdrental C:\users\sample_datbase\dvdrental.tar

# Listing tables in PostgreSQL using psql command
\dt
\dt+ # for more info

# Showing tables using pg_catalog schema
SELECT *
FROM pg_catalog.pg_tables
WHERE schemaname != 'pg_catalog' AND 
    schemaname != 'information_schema';
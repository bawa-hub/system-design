# PostgreSQL DESCRIBE TABLE using psql
\d <table_name>

# DESCRIBE TABLE using information_schema
SELECT 
   table_name, 
   column_name, 
   data_type 
FROM 
   information_schema.columns
WHERE 
   table_name = 'city';
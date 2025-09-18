-- get info about columns:
DESCRIBE table_name;
DESC table_name;

-- list of columns in a table
SHOW COLUMNS FROM table_name;

-- from any database
SHOW COLUMNS FROM database_name.table_name;

-- get more info about column
SHOW FULL COLUMNS FROM table_name;

-- filter the columns of the table
SHOW COLUMNS FROM table_name LIKE pattern;
SHOW COLUMNS FROM table_name WHERE expression;

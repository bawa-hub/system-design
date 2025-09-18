SELECT col1, col2, col3 FROM table_name;
SELECT * FROM table_name;

-- MySQL doesn’t require the FROM clause
SELECT select_list;
SELECT 1+1;
SELECT NOW();

-- column alias
SELECT expression AS column_alias;
-- AS keyword is optonal
SELECT expression column_alias;
-- If the column alias contains spaces
SELECT CONCAT('Jane', ' ', 'Dow') AS 'Full name';
-- list all users in a MySQL database server
SELECT user FROM mysql.user;
-- OR
USE MySQL;
SELECT user FROM user;

-- get more information on the user table
DESC user;

-- get the information on the  current user
SELECT user();
-- OR
SELECT current_user();

-- list all users that are currently logged in the MySQL database server
SELECT 
    user, 
    host, 
    db, 
    command 
FROM 
    information_schema.processlist;




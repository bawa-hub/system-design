-- syntax
-- CREATE USER [IF NOT EXISTS] account_name
-- IDENTIFIED BY 'password';

-- account_name convention
--  username@hostname 
-- username without hostname equivalent to
-- username@% 
-- username and hostname contains special chars
-- 'username'@'hostname'



-- connect to mysql server
mysql -u root -p

-- show users
select user from mysql.user

-- create user
CREATE USER bawa@localhost IDENTIFIED BY 'Bawa.vikram';

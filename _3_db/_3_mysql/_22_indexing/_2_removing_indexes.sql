-- To remove an existing index from a table
DROP INDEX index_name ON table_name
[algorithm_option | lock_option];

-- algorith_option
ALGORITHM [=] {DEFAULT|INPLACE|COPY}

-- For the index removal, the following algorithms are supported:
-- COPY: The table is copied to the new table row by row, the DROP INDEX is then performed on the copy of the original table. The concurrent data manipulation statements such as INSERT and UPDATE are not permitted.
-- INPLACE: The table is rebuilt in place instead of copied to the new one. MySQL issues an exclusive metadata lock on the table during the preparation and execution phases of the index removal operation. This algorithm allows for concurrent data manipulation statements

-- ALGORITHM clause is optional. If you skip it, MySQL uses INPLACE. In case the INPLACE is not supported, MySQL uses COPY.
-- DEFAULT has the same effect as omitting the ALGORITHM clause.


-- lock_option
-- lock_option controls the level of concurrent reads and writes on the table while the index is being removed.
LOCK [=] {DEFAULT|NONE|SHARED|EXCLUSIVE}

-- DEFAULT: this allows you to have the maximum level of concurrency for a given algorithm. First, it allows concurrent reads and writes if supported. If not, allow concurrent reads if supported. If not, enforce exclusive access.
-- NONE: if the NONE is supported, you can have concurrent reads and writes. Otherwise, MySQL issues an error.
-- SHARED: if the SHARED is supported, you can have concurrent reads, but not writes. MySQL issues an error if the concurrent reads are not supported.
-- EXCLUSIVE: this enforces exclusive access.


-- e.g

CREATE TABLE leads(
    lead_id INT AUTO_INCREMENT,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    information_source VARCHAR(255),
    INDEX name(first_name,last_name),
    UNIQUE email(email),
    PRIMARY KEY(lead_id)
);

-- removes the name index from the leads table:
DROP INDEX name ON leads;
-- drops the email index from the leads table with a specific algorithm and lock:
DROP INDEX email ON leads
ALGORITHM = INPLACE 
LOCK = DEFAULT;

-- To drop the primary key whose index name is PRIMARY
DROP INDEX `PRIMARY` ON table_name;




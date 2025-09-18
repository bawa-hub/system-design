-- create indexes for a table at the time of creation
CREATE TABLE t(
   c1 INT PRIMARY KEY,
   c2 INT NOT NULL,
   c3 INT NOT NULL,
   c4 VARCHAR(10),
   INDEX (c2, c3)
);
-- To add an index for a column or a set of columns
CREATE INDEX index_name ON table_name (column_list) -- By default, MySQL creates the B-Tree index if you don’t specify the index type.
-- permissible index type based on the storage engine of the table
Storage Engine - Allowed Index Types InnoDB - BTREE MyISAM - BTREE MEMORY / HEAP - HASH,
BTREE -- show indexes of a table
SHOW INDEXES
FROM table;
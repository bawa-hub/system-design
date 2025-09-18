-- MySQL uses indexes to quickly find rows with specific column values. 

-- Without an index, MySQL must scan the whole table to locate the relevant rows. 

-- The larger table, the slower it searches.
-- An index is a data structure such as B-Tree that improves the speed of data retrieval on a table 
-- at the cost of additional writes and storage to maintain it.

-- query optimizer may use indexes to quickly locate data without having to scan every row in a table for a given query.

-- When you create a table with a primary key or unique key, 
-- MySQL automatically creates a special index named PRIMARY. This index is called the clustered index.
-- Other indexes other than the PRIMARY index are called secondary indexes or non-clustered indexes.


-- https://www.analyticsvidhya.com/blog/2021/06/understand-the-concept-of-indexing-in-depth/
-- https://www.freecodecamp.org/news/database-indexing-at-a-glance-bb50809d48bd/
-- https://www.geeksforgeeks.org/indexing-in-databases-set-1/
-- https://levelup.gitconnected.com/an-in-depth-look-at-database-indexing-for-performanceand-what-are-the-trade-offs-for-using-them-e2debd4b5c1d
-- https://www.educba.com/mysql-index-types/
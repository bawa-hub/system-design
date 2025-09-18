-- Supported joins in Mysql:
-- inner join, left join, right join, cross join, self join

-- create two tables called members and committees:
CREATE TABLE members (
    member_id INT AUTO_INCREMENT,
    name VARCHAR(100),
    PRIMARY KEY (member_id)
);

CREATE TABLE committees (
    committee_id INT AUTO_INCREMENT,
    name VARCHAR(100),
    PRIMARY KEY (committee_id)
);

INSERT INTO members(name)
VALUES('John'),('Jane'),('Mary'),('David'),('Amelia');

INSERT INTO committees(name)
VALUES('John'),('Mary'),('Amelia'),('Joe');

-- Inner Join
-- SELECT column_list
-- FROM table_1
-- INNER JOIN table_2 ON join_condition;

-- The inner join clause compares each row from the first table with every row from the second table.
-- the inner join clause includes only matching rows from both tables.
SELECT column_list
FROM table_1
INNER JOIN table_2 USING (column_name); -- if column_name is same in both tables

-- find members who are also the committee members:
SELECT 
    m.member_id, 
    m.name AS member, 
    c.committee_id, 
    c.name AS committee
FROM
    members m
INNER JOIN committees c ON c.name = m.name; -- INNER JOIN committees c USING(name);



-- Left Join
-- the left join selects all data from the left table whether there are matching rows exist in the right table or not.
-- In case there are no matching rows from the right table found, the left join uses NULLs for columns of the row from the right table in the result set.
-- SELECT column_list 
-- FROM table_1 
-- LEFT JOIN table_2 ON join_condition;

-- uses a left join clause to join the members with the committees table:
SELECT 
    m.member_id, 
    m.name AS member, 
    c.committee_id, 
    c.name AS committee
FROM
    members m
LEFT JOIN committees c USING(name);

-- find members who are not the committee members,
SELECT 
    m.member_id, 
    m.name AS member, 
    c.committee_id, 
    c.name AS committee
FROM
    members m
LEFT JOIN committees c USING(name)
WHERE c.committee_id IS NULL;

-- Right Join
-- SELECT column_list 
-- FROM table_1 
-- RIGHT JOIN table_2 ON join_condition;


-- find rows in the right table that does not have corresponding rows in the left table
SELECT column_list 
FROM table_1 
RIGHT JOIN table_2 USING (column_name)
WHERE column_table_1 IS NULL;

--  use the right join to join the members and committees tables:
SELECT 
    m.member_id, 
    m.name AS member, 
    c.committee_id, 
    c.name AS committee
FROM
    members m
RIGHT JOIN committees c on c.name = m.name;

-- find the committee members who are not in the members table,
SELECT 
    m.member_id, 
    m.name AS member, 
    c.committee_id, 
    c.name AS committee
FROM
    members m
RIGHT JOIN committees c USING(name)
WHERE m.member_id IS NULL;


-- Cross Join
-- cross join makes a Cartesian product of rows from the joined tables

-- Suppose the first table has n rows and the second table has m rows. The cross join that joins the tables will return nxm rows.
-- SELECT select_list
-- FROM table_1
-- CROSS JOIN table_2;

--  use the cross join clause to join the members with the committees tables:
SELECT 
    m.member_id, 
    m.name AS member, 
    c.committee_id, 
    c.name AS committee
FROM
    members m
CROSS JOIN committees c;
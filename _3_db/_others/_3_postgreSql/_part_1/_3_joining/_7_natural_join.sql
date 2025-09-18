-- A natural join is a join that creates an implicit join based on the same column names in the joined tables.

-- syntax
SELECT select_list
FROM T1
NATURAL [INNER, LEFT, RIGHT] JOIN T2;

-- A natural join can be an inner join, left join, or right join. 
-- If you do not specify a join explicitly e.g., INNER JOIN, LEFT JOIN, RIGHT JOIN, 
-- PostgreSQL will use the INNER JOIN by default.


-- sample data for demonstration
DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
	category_id serial PRIMARY KEY,
	category_name VARCHAR (255) NOT NULL
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
	product_id serial PRIMARY KEY,
	product_name VARCHAR (255) NOT NULL,
	category_id INT NOT NULL,
	FOREIGN KEY (category_id) REFERENCES categories (category_id)
);

INSERT INTO categories (category_name)
VALUES
	('Smart Phone'),
	('Laptop'),
	('Tablet');

INSERT INTO products (product_name, category_id)
VALUES
	('iPhone', 1),
	('Samsung Galaxy', 1),
	('HP Elite', 2),
	('Lenovo Thinkpad', 2),
	('iPad', 3),
	('Kindle Fire', 3);


-- join the products table with the categories table:
SELECT * FROM products
NATURAL JOIN categories;
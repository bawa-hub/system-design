-- to check if a value is NULL or not.

-- sample data for demonstration
CREATE TABLE contacts(
    id SERIAL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(15),
    PRIMARY KEY (id)
);

INSERT INTO contacts(first_name, last_name, email, phone)
VALUES ('John','Doe','john.doe@example.com',NULL),
    ('Lily','Bush','lily.bush@example.com','(408-234-2764)');


-- find the contact who does not have a phone number
SELECT
    id,
    first_name,
    last_name,
    email,
    phone
FROM
    contacts
WHERE
    phone = NULL;
-- The statement returns no row. 
-- This is because the expression phone = NULL in the WHERE clause always returns false.    

-- Even though there is a NULL in the phone column, the expression NULL = NULL returns false. 
-- This is because NULL is not equal to any value even itself

-- use the following statement instead:
SELECT
    id,
    first_name,
    last_name,
    email,
    phone
FROM
    contacts
WHERE
    phone IS NULL;


-- find the contact who does have a phone number,
SELECT
    id,
    first_name,
    last_name,
    email,
    phone
FROM
    contacts
WHERE
    phone IS NOT NULL;
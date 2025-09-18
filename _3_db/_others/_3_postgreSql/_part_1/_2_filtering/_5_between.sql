-- match a value against a range of values.

-- syntax
value BETWEEN low AND high;
-- equivalent to 
value >= low and value <= high

-- to check if a value is out of a range
value NOT BETWEEN low AND high;
-- equivalent to
value < low OR value > high

-- select payments whose amount is between 8 and 9 (USD)
SELECT
	customer_id,
	payment_id,
	amount
FROM
	payment
WHERE
	amount BETWEEN 8 AND 9;

-- get payments whose amount is not in the range of 8 and 9
    SELECT
	customer_id,
	payment_id,
	amount
FROM
	payment
WHERE
	amount NOT BETWEEN 8 AND 9;


-- If you want to check a value against of date ranges, you should use the literal date in ISO 8601 format i.e., YYYY-MM-DD.



-- get the payment whose payment date is between 2007-02-07 and 2007-02-15
SELECT
	customer_id,
	payment_id,
	amount,
 payment_date
FROM
	payment
WHERE
	payment_date BETWEEN '2007-02-07' AND '2007-02-15';
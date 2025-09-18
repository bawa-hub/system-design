-- employees who work in the offices located in the USA.

SELECT 
    lastName, firstName
FROM
    employees
WHERE
    officeCode IN (SELECT 
            officeCode
        FROM
            offices
        WHERE
            country = 'USA');

--             +-----------+-----------+
-- | lastName  | firstName |
-- +-----------+-----------+
-- | Murphy    | Diane     |
-- | Patterson | Mary      |
-- | Firrelli  | Jeff      |
-- | Bow       | Anthony   |
-- | Jennings  | Leslie    |
-- | Thompson  | Leslie    |
-- | Firrelli  | Julie     |
-- | Patterson | Steve     |
-- | Tseng     | Foon Yue  |
-- | Vanauf    | George    |
-- +-----------+-----------+
-- 10 rows in set (0.052 sec)

-- eloquent version
-- Employees::whereIn('officeCode, $officeCodes)
-- MySQL subquery with IN and NOT IN operators

--  find the customers who have not placed any orders as follows:

SELECT 
    customerName
FROM
    customers
WHERE
    customerNumber NOT IN (SELECT DISTINCT
            customerNumber
        FROM
            orders);


--             +--------------------------------+
-- | customerName                   |
-- +--------------------------------+
-- | Havel & Zbyszek Co             |
-- | American Souvenirs Inc         |
-- | Porto Imports Co.              |
-- | Asian Shopping Network, Co     |
-- | Natürlich Autos                |
-- | ANG Resellers                  |
-- | Messner Shopping Network       |
-- | Franken Gifts, Co              |
-- | BG&E Collectables              |
-- | Schuyler Imports               |
-- | Der Hund Imports               |
-- | Cramer Spezialitäten, Ltd      |
-- | Asian Treasures, Inc.          |
-- | SAR Distributors, Co           |
-- | Kommission Auto                |
-- | Lisboa Souveniers, Inc         |
-- | Precious Collectables          |
-- | Stuttgart Collectable Exchange |
-- | Feuer Online Stores, Inc       |
-- | Warburg Exchange               |
-- | Anton Designs, Ltd.            |
-- | Mit Vergnügen & Co.            |
-- | Kremlin Collectables, Co.      |
-- | Raanan Stores, Inc             |
-- +--------------------------------+
-- 24 rows in set (0.001 sec)
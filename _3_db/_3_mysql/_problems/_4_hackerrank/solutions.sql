-- 1. Revision the select query 1
select * from CITY Where COUNTRYCODE = 'USA' AND POPULATION > 100000;

-- 2. Revision the select query 2
SELECT NAME FROM CITY WHERE COUNTRYCODE = 'USA' AND POPULATION > 120000;

-- 3. Select All
SELECT * FROM CITY;

-- 4. Select By ID
SELECT * FROM CITY WHERE ID = 1661;

-- 5. Japanes city attributes
SELECT * FROM CITY WHERE COUNTRYCODE = 'JPN';

-- 6. Japanes city names
SELECT NAME FROM CITY WHERE COUNTRYCODE = 'JPN'

-- 7. Whether Obsevation Station 1
SELECT CITY, STATE FROM STATION;

-- 8. Whether Obsevation Station 3
SELECT DISTINCT CITY FROM STATION WHERE ID%2 = 0;

-- 9. Whether Obsevation Station 4
SSELECT COUNT(CITY) - COUNT(DISTINCT CITY) FROM STATION;

-- 10. Whether Obsevation Station 5
SELECT City, length(City) AS len FROM STATION
ORDER BY len asc, City
LIMIT 1;

SELECT City, length(City) AS len FROM STATION
ORDER BY len desc, City
LIMIT 1;

-- 11. Whether Obsevation Station 6
SELECT DISTINCT CITY FROM STATION WHERE LOWER(SUBSTRING(CITY ,1,1)) IN ('a','e','i','o','u');

-- 12. Whether Obsevation Station 7
SELECT DISTINCT CITY FROM STATION WHERE LOWER(SUBSTRING(CITY ,-1)) IN ('a','e','i','o','u');

-- 13. Whether Obsevation Station 8
SELECT DISTINCT CITY FROM STATION WHERE (LOWER(SUBSTRING(CITY ,1,1)) IN ('a','e','i','o','u')) AND (LOWER(SUBSTRING(CITY ,-1)) IN ('a','e','i','o','u'));

select Distinct City from Station where left(city,1) in ('A','E','I','O','U') and right(city,1) in ('a','e','i','o','u') ;

-- 14. Whether Obsevation Station 9
SELECT DISTINCT CITY FROM STATION WHERE LOWER(SUBSTRING(CITY ,1,1)) NOT IN ('a','e','i','o','u');

-- 15. Whether Obsevation Station 10
SELECT DISTINCT CITY FROM STATION WHERE LOWER(SUBSTRING(CITY ,-1)) NOT IN ('a','e','i','o','u');

-- 16. Whether Obsevation Station 11
SELECT DISTINCT CITY FROM STATION WHERE LOWER(SUBSTRING(CITY ,1,1)) NOT IN ('a','e','i','o','u') OR LOWER(SUBSTRING(CITY ,-1)) NOT IN ('a','e','i','o','u');

-- 17. Whether Obsevation Station 12
SELECT DISTINCT CITY FROM STATION WHERE LOWER(SUBSTRING(CITY ,1,1)) NOT IN ('a','e','i','o','u') AND LOWER(SUBSTRING(CITY ,-1)) NOT IN ('a','e','i','o','u');

-- 18. Higher than 75 marks
SELECT Name FROM STUDENTS WHERE Marks > 75 ORDER BY SUBSTRING(Name, -3, LENGTH(Name)) ASC, ID;

-- 19. Employee Names
SELECT NAME FROM Employee ORDER BY name ASC;

-- 20. Employee Salaries
SELECT name FROM Employee WHERE salary > 2000 AND months < 10 ORDER BY employee_id;

-- 21. Type of Traingle
SELECT 
    CASE 
    WHEN (A + B <= C) OR (B + C <= A) OR (A + C <= B) THEN 'Not A Triangle'
    WHEN (A=B) AND (B=C) THEN 'Equilateral' 
    WHEN (A<>B) AND (B<>C) AND (A<>C) THEN 'Scalene'
    WHEN ((A=B) AND (A<>C)) OR ((A=C) AND (A<>B)) OR ((B=C) AND (B<>A)) THEN 'Isosceles'
    END
from TRIANGLES;

-- 22. The PADS
select concat(name,'(', left(occupation,1),')' ) from OCCUPATIONS order by name, occupation; 

select concat('There are a total of ', count(*),' ' ,lower(occupation),'s.') from occupations 
group by occupation order by count(occupation),occupation asc ;

-- 23. Occupations
-- 24. Binary Tree Nodes
-- 25. New Companies
-- 26. Revising Aggregations -The Count Function
-- 27. Revising Aggregations -The Sum Function
-- 28. Revising Aggregations - Averages
-- 29. Average Population
-- 30. Japan Population
-- 31. Population Density Difference
-- 32. The Blunder
-- 33. The Earners
-- 34. Whether Obsevation Station 2
-- 35. Whether Obsevation Station 13
-- 35. Whether Obsevation Station 14
-- 36. Whether Obsevation Station 15
-- 37. Whether Obsevation Station 16
-- 38. Whether Obsevation Station 17
-- 39. Whether Obsevation Station 18
-- 40. Whether Obsevation Station 19
-- 42. Whether Obsevation Station 20
-- 43. Population Census
-- 44. African Cities
-- 45. Average Population of Each Continent
-- 46. The Report 
-- 47. Top Competitors
-- 48. Ollivander's Inventory
-- 49. Challenges
-- 50. Contest Leaderboard
-- 51. SQL Project Planning
-- 52. Placements
-- 53. Symmetric Pairs
-- 54. Interviews
-- 55. 15 Days of Learning SQl
-- 56. Draw The Traingle 1
-- 57. Draw THe Traingle 2
-- 58. Print Prime Numbers


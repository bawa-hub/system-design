-- BEFORE INSERT triggers are automatically fired before an insert event occurs on the table.

-- CREATE TRIGGER trigger_name
--     BEFORE INSERT
--     ON table_name FOR EACH ROW
-- trigger_body;


--  create a new table called WorkCenters:
DROP TABLE IF EXISTS WorkCenters;

CREATE TABLE WorkCenters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    capacity INT NOT NULL
);


--  create another table called WorkCenterStats that stores the summary of the capacity of the work centers:
DROP TABLE IF EXISTS WorkCenterStats;

CREATE TABLE WorkCenterStats(
    totalCapacity INT NOT NULL
);


-- Creating BEFORE INSERT trigger example
-- trigger updates the total capacity in the WorkCenterStats table before a new work center is inserted into the WorkCenter table:
DELIMITER $$

CREATE TRIGGER before_workcenters_insert
BEFORE INSERT
ON WorkCenters FOR EACH ROW
BEGIN
    DECLARE rowcount INT;
    
    SELECT COUNT(*) 
    INTO rowcount
    FROM WorkCenterStats;
    
    IF rowcount > 0 THEN
        UPDATE WorkCenterStats
        SET totalCapacity = totalCapacity + new.capacity;
    ELSE
        INSERT INTO WorkCenterStats(totalCapacity)
        VALUES(new.capacity);
    END IF; 

END $$

DELIMITER ;



-- Testing the MySQL BEFORE INSERT trigger
INSERT INTO WorkCenters(name, capacity)
VALUES('Mold Machine',100);


SELECT * FROM WorkCenterStats;    

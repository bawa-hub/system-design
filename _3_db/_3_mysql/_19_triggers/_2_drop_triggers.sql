-- DROP TRIGGER [IF EXISTS] [schema_name.]trigger_name;


-- create a table called billings
CREATE TABLE billings (
    billingNo INT AUTO_INCREMENT,
    customerNo INT,
    billingDate DATE,
    amount DEC(10 , 2 ),
    PRIMARY KEY (billingNo)
);

-- create a new trigger called BEFORE UPDATE that is associated with the billings table
DELIMITER $$
CREATE TRIGGER before_billing_update
    BEFORE UPDATE 
    ON billings FOR EACH ROW
BEGIN
    IF new.amount > old.amount * 10 THEN
        SIGNAL SQLSTATE '45000' 
            SET MESSAGE_TEXT = 'New amount cannot be 10 times greater than the current amount.';
    END IF;
END$$    
DELIMITER ;

-- drop the before_billing_update trigger:
DROP TRIGGER before_billing_update;

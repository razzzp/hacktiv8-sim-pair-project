
-- for creating initial data
INSERT INTO `products` (`ID`, `Name`, `Price`, `Stock`) VALUES
(1, 'Product1', 100.00, 4),
(2, 'Product2', 2.50, 100);

INSERT INTO `staff` (`ID`, `Name`, `Email`, `Position`) VALUES
(1, 'Bob', 'bob@mail.com', 'Sales'),
(2, 'Alice', 'alice@mail.com', 'POS');

INSERT INTO sales (ProductID, Quantity, Sale_date)
VALUES 
(1, 4, NOW()),
(2, 10, NOW());
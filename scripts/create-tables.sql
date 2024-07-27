-- create products table

-- DROP TABLES
DROP TABLE IF EXISTS Sales;
DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Staff;
--
CREATE TABLE Products(
	ID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100),
    Price DECIMAL(10,2),
    Stock INT,
 	CONSTRAINT Chk_Price CHECK (Price >=0),
	CONSTRAINT Chk_Stock CHECK (Stock >= 0)
);

-- create staff table
CREATE TABLE Staff(
	ID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(100),
    Email VARCHAR(50) UNIQUE NOT NULL,
    Position VARCHAR(50)
);

-- create sales table
CREATE TABLE Sales(
	ID INT PRIMARY KEY AUTO_INCREMENT,
    ProductID INT NOT NULL,
    Quantity INT NOT NULL,
    Sale_date DATE NOT NULL,
	CONSTRAINT fk_salesProduct FOREIGN KEY (ProductID) 	REFERENCES Products(ID)
);


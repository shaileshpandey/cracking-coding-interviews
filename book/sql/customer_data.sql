CREATE TABLE customers (
    customer_id bigserial NOT NULL PRIMARY KEY,
    customer_name character varying(30),
    contact_title character varying(30),
    address character varying(60),
    city character varying(15),
    region character varying(15),
    postal_code character varying(10),
    country character varying(15),
    phone character varying(24)
);

CREATE TABLE products (
    product_id serial NOT NULL,
    product_name character varying(40) NOT NULL,
    unit_price real
);

CREATE TABLE orders (
    order_id serial NOT NULL,
    customer_id integer,
    order_date date,
    shipped_date date
);

CREATE TABLE order_details (
    order_id serial NOT NULL,
    product_id integer NOT NULL,
    unit_price real NOT NULL,
    quantity smallint NOT NULL,
    discount real NOT NULL
);

INSERT INTO customers VALUES 
(1, 'Maria Anders', 'Sales Representative', 'Obere Str. 57', 'Berlin', NULL, '12209', 'Germany', '030-0074321'),
(2, 'Frédérique Citeaux', 'Marketing Manager', '24, place Kléber', 'Strasbourg', NULL, '67000', 'France', '88.60.15.31'),
(3, 'Elizabeth Lincoln', 'Accounting Manager', '23 Tsawassen Blvd.', 'Tsawassen', 'BC', 'T2F 8M4', 'Canada', '(604) 555-4729'),
(4, 'Yang Wang', 'Owner', 'Hauptstr. 29', 'Bern', NULL, '3012', 'Switzerland', '0452-076545'),
(5, 'Elizabeth Brown', 'Sales Representative', 'Berkeley Gardens 12  Brewery', 'London', NULL, 'WX1 6LT', 'UK', '(171) 555-2282'),
(6, 'Ann Devon', 'Sales Agent', '35 King George', 'London', NULL, 'WX3 6FW', 'UK', '(171) 555-0297'),
(7, 'Peter Franken', 'Marketing Manager', 'Berliner Platz 43', 'München', NULL, '80805', 'Germany', '089-0877310'),
(8, 'Paolo Accorti', 'Sales Representative', 'Via Monte Bianco 34', 'Torino', NULL, '10100', 'Italy', '011-4988260'),
(9, 'Yoshi Latimer', 'Sales Representative', 'City Center Plaza 516 Main St.', 'Elgin', 'OR', '97827', 'USA', '(503) 555-6874'),
(10, 'Jaime Yorres', 'Owner', '87 Polk St. Suite 5', 'San Francisco', 'CA', '94117', 'USA', '(415) 555-5938');

INSERT INTO products VALUES 
(11, 'Baseball', 5.5),
(12, 'Socks', 7.50),
(13, 'Water Bottle', 8.40),
(14, 'Running Shoes', 75.99),
(15, 'Hockey Stick', 49.95),
(16, 'Soccer Ball', 27.99),
(17, 'Batting Gloves', 15.99),
(18, 'Sweat Pants', 45.99),
(19, 'Camping Tent', 250),
(20, 'Canoe', 550),
(21, 'Mountain Bike', 850),
(22, 'Shorts', 35);

INSERT INTO orders VALUES 
(248, 1, '1996-07-04', '1996-08-01'),
(249, 2, '1996-07-05', '1996-08-16'),
(250, 3, '1996-07-08', '1996-08-05'),
(251, 4, '1996-07-08', '1996-08-05'),
(252, 5, '1996-07-09', '1996-08-06'),
(253, 6, '1996-07-10', '1996-07-24'),
(254, 7, '1996-07-11', '1996-08-08'),
(255, 8, '1996-07-12', '1996-08-09'),
(256, 9, '1996-07-15', '1996-08-12'),
(257, 5, '1996-07-16', '1996-08-13'),
(257, 1, '1996-08-21', '1996-08-23'),
(257, 2, '1996-09-22', '1996-09-27'),
(257, 5, '1996-10-16', '1996-10-20'),
(258, 1, '1996-10-18', '1996-10-22');

INSERT INTO order_details VALUES 
(248, 11, 5.5, 1, 0),
(248, 16, 27.99, 1, 0),
(248, 12, 7.5, 5, 0),
(249, 19, 212.5, 1, 0.15),
(249, 15, 49.95, 20, 0),
(250, 20, 495, 1, 0.10),
(250, 13, 7.14, 3, 0.15),
(250, 12, 6.375, 15, 0.15),
(251, 17, 15.1905, 6, 0.05),
(251, 22, 33.25, 2, 0.05),
(251, 16, 27.99, 20, 0),
(252, 13, 7.98, 4, 0.05),
(252, 18, 43.6905, 2, 0.05),
(252, 21, 680, 1, 0.20),
(253, 14, 75.99, 1, 0),
(253, 19, 237.5, 42, 0.05),
(253, 13, 8.40, 4, 0),
(254, 13, 8.40, 4, 0),
(254, 12, 6.375, 15, 0.15),
(258, 22, 35, 4, 0);
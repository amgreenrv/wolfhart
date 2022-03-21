USE wolfhart;
CREATE TABLE products (
    id INT PRIMARY KEY NOT NULL,
    prodname VARCHAR(45) NOT NULL,
    category VARCHAR(45) NOT NULL,
    price FLOAT NOT NULL,
    img VARCHAR(255) NOT NULL,
    proddesc VARCHAR(255) NOT NULL,
    featured TINYINT(1) NOT NULL
);

INSERT INTO products (
    id,prodname,category,price,img,proddesc,featured
)

VALUES
(111,'H1 Mini-ITX','case',199.99,'img/product/case-h1.png','Compact Mini-ITX Case w/ AIO PSU & Riser',0),
(112,'H210 Mid-Tower','case',149.99,'img/product/case-h210.png','Mid-Tower Case with Tempered Glass',1),
(113,'H510 Special Edition','case',249.99,'img/product/case-h510SE.png','Limited Edition Compact Mid-Tower',0),
(114,'N7 B550','power',129.99,'img/product/motherboard.png','AMD Motherboard with Wi-Fi and Controller',1),
(115,'C850 Gold','power',199.99,'img/product/psu.png','850 Watt Power Supply Unit',1),
(116,'Aer RGB 2 140mm','accessories',79.99,'img/product/aer-140.png','Aer RGB 2 140mm fans with RGB & Fan Controller',0),
(117,'Kraken X63 AIO','accessories',299.99,'img/product/aio-x63.png','360mm Liquid Cooler with LCD Display',1),
(118,'Kraken Z73 AIO','accessories',189.99,'img/product/aio-z73.png','AIO Liquid Cooler with Infinity Mirror Display',0);

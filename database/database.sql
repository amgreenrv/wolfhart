USE wolfhart;
CREATE TABLE products (
    id INT PRIMARY KEY NOT NULL,
    prodname VARCHAR(45) NOT NULL,
    category VARCHAR(45) NOT NULL,
    price FLOAT NOT NULL,
    img VARCHAR(255) NOT NULL,
    proddesc VARCHAR(255) NOT NULL,
    featured VARCHAR(45) NOT NULL
);

INSERT INTO products (
    id,prodname,category,price,img,proddesc,featured
)

VALUES
(111, 'H1 Mini-ITX', 'case', 199.99, 'img/product/case-h1.png', 'Compact Mini-ITX Case w/ AIO PSU & Riser', 'false'),
(112, 'H210 Mid-Tower', 'case', 149.99, 'img/product/case-h210.png', 'Mid-Tower Case with Tempered Glass', 'true'),
(113, 'H510 Special Edition', 'case', 249.99, 'img/product/case-h510SE.png', 'Limited Edition Compact Mid-Tower', 'false'),
(114, 'N7 B550', 'components', 129.99, 'img/product/motherboard.png', 'AMD Motherboard with Wi-Fi and Controller', 'true'),
(115, 'C850 Gold', 'power', 199.99, 'img/product/psu.png', '850 Watt Power Supply Unit', 'true'),
(116, 'Aer RGB 2 140mm', 'accessories', 79.99, 'img/product/aer-140.png', 'Aer RGB 2 140mm fans with RGB & Fan Controller', 'false'),
(117, 'Kraken X63 AIO', 'accessories', 299.99, 'img/product/aio-x63.png', '360mm Liquid Cooler with LCD Display', 'true'),
(118, 'Kraken Z73 AIO', 'accessories', 189.99, 'img/product/aio-z73.png', 'AIO Liquid Cooler with Infinity Mirror Display', 'false');

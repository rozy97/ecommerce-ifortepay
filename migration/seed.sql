INSERT INTO users(email, password, created_at, updated_at, deleted_at) VALUES
('userdummy@mail.com', MD5(CONCAT('1234', 's3cr3tp4ssw0rd')), NOW(), NOW(), NULL);

INSERT INTO products(sku, name, price, quantity, created_at, updated_at, deleted_at) VALUES
('120P90', 'Google Home', 49.99, 1000, NOW(), NOW(), NULL),
('43N23P', 'Macbook Pro', 5399.99, 1000, NOW(), NOW(), NULL),
('A304SD', 'Alexa Speaker', 109.50, 1000, NOW(), NOW(), NULL),
('234234', 'Raspberry Pi B', 30.00, 1000, NOW(), NOW(), NULL);
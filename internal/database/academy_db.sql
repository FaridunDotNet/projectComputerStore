create database academy_db;

-- Создание таблицы для администраторов
CREATE TABLE IF NOT EXISTS admins (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );

-- Создание таблицы для категорий товара
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS customers (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    date_of_birth DATE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );
-- Создание таблицы для клиентов

-- Создание таблицы для заказов
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) NOT NULL,
    customer_id INTEGER REFERENCES customers(id) NOT NULL,
    title VARCHAR(100) NOT NULL,
    start_date DATE DEFAULT NULL,
    finish_date DATE DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );

-- Создание таблицы для деталей заказа
CREATE TABLE IF NOT EXISTS order_details (
                                             id SERIAL PRIMARY KEY,
                                             order_id INTEGER NOT NULL REFERENCES orders(id),
    product_id INTEGER NOT NULL REFERENCES products(id),
    quantity_ordered INTEGER NOT NULL,
    unit_price FLOAT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );

-- Создание таблицы для продуктов
CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(255) NOT NULL,
    description TEXT,
    price FLOAT NOT NULL,
    quantity_available INTEGER NOT NULL,
    category_id INTEGER NOT NULL REFERENCES categories(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );


-- Создание таблицы для отзывов
CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES products(id) NOT NULL,
    customer_id INTEGER REFERENCES customers(id) NOT NULL,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 10),
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
    );

--Список всех заказов с информацией о клиентах и продуктах
SELECT
    orders.id AS order_id,
    customers.first_name,
    customers.last_name,
    customers.email,
    products.name AS product_name,
    products.price AS product_price,
    orders.start_date,
    orders.finish_date
FROM
    orders
        JOIN
    customers ON orders.customer_id = customers.id
        JOIN
    products ON orders.product_id = products.id;

--Список всех деталей заказов с информацией о заказах и продуктах
SELECT
    order_details.id AS order_detail_id,
    orders.id AS order_id,
    customers.first_name,
    customers.last_name,
    products.name AS product_name,
    order_details.quantity_ordered,
    order_details.unit_price,
    (order_details.quantity_ordered * order_details.unit_price) AS total_price,
    order_details.created_at
FROM
    order_details
        JOIN
    orders ON order_details.order_id = orders.id
        JOIN
    customers ON orders.customer_id = customers.id
        JOIN
    products ON order_details.product_id = products.id;

--Отчет о заказах с суммарной стоимостью по каждому клиенту
SELECT
    customers.id AS customer_id,
    customers.first_name,
    customers.last_name,
    customers.email,
    SUM(order_details.quantity_ordered * order_details.unit_price) AS total_spent
FROM
    customers
        JOIN
    orders ON customers.id = orders.customer_id
        JOIN
    order_details ON orders.id = order_details.order_id
GROUP BY
    customers.id, customers.first_name, customers.last_name, customers.email
ORDER BY
    total_spent DESC;

-- Список всех продуктов с их категориями
SELECT
    products.id AS product_id,
    products.name AS product_name,
    products.description,
    products.price,
    categories.name AS category_name,
    products.quantity_available
FROM
    products
        JOIN
    categories ON products.category_id = categories.id;


-- Список всех отзывов с информацией о продуктах и клиентах
SELECT
    reviews.id AS review_id,
    products.name AS product_name,
    customers.first_name,
    customers.last_name,
    reviews.rating,
    reviews.comment,
    reviews.created_at
FROM
    reviews
        JOIN
    products ON reviews.product_id = products.id
        JOIN
    customers ON reviews.customer_id = customers.id
ORDER BY
    reviews.created_at DESC;


--Отчёт о среднем рейтинге продукции
SELECT
products.id AS product_id,
    products.name AS product_name,
    AVG(reviews.rating) AS average_rating
FROM
    products
JOIN
    reviews ON products.id = reviews.product_id
GROUP BY
    products.id, products.name
ORDER BY
    average_rating DESC;



CREATE TABLE "user" (
    id INTEGER GENERATED ALWAYS AS IDENTITY CONSTRAINT user_table_pk PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(512) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE product (
    nm_id INTEGER CONSTRAINT product_table_pk PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    brand_id INT NOT NULL,
    site_brand_id INT NOT NULL,
    supplier_id INT NOT NULL,
    sale INT NOT NULL,
    price INT NOT NULL,
    sale_price INT NOT NULL,
    rating FLOAT NOT NULL,
    feedbacks INT NOT NULL,
    colors VARCHAR(255),
    quantity INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE price_history (
    id INTEGER GENERATED ALWAYS AS IDENTITY CONSTRAINT price_history_table_pk PRIMARY KEY,
    nm_id INT NOT NULL,
    dt TIMESTAMP NOT NULL,
    price INT NOT NULL,
    FOREIGN KEY (nm_id) REFERENCES product(nm_id) ON DELETE CASCADE
);

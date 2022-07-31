CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    email VARCHAR(60) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    password VARCHAR(80) NOT NULL,
    types INTEGER NULL DEFAULT NULL,
    CONSTRAINT users FOREIGN KEY (types) REFERENCES user_type (id),
    email_verified_at TIMESTAMP NULL DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL
    );
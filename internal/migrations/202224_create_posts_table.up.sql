CREATE TABLE IF NOT EXISTS posts(
   id serial PRIMARY KEY,
    title VARCHAR(30) NOT NULL,
    slug VARCHAR(60) NOT NULL,
    details VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    blueTick Boolean DEFAULT FALSE,
    user_id INTEGER NOT NULL,
    CONSTRAINT posts FOREIGN KEY (user_id) REFERENCES users (id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL
    );
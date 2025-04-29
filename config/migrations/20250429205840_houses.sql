-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users(
    id uuid NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) UNIQUE NOT NULL,
    created_at NUMERIC,
    updated_at NUMERIC,
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS listings(
    id uuid NOT NULL PRIMARY KEY,
    user_id uuid NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    address VARCHAR(255) NOT NULL,
    rent DECIMAL NOT NULL,
    number_of_bedrooms DECIMAL NOT NULL,
    number_of_bathrooms DECIMAL NOT NULL,
    created_at NUMERIC,
    updated_at NUMERIC,
    is_deleted BOOLEAN DEFAULT FALSE,
    CONSTRAINT listings_user_id_foreign_key FOREIGN KEY (user_id) REFERENCES users(id)
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS listings;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd

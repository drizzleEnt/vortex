-- +goose Up
CREATE TABLE clients(
    id serial primary key,
    client_name VARCHAR(256) NOT NULL,
    version INT NOT NULL DEFAULT 0,
    image VARCHAR(256),
    cpu VARCHAR(256),
    memory VARCHAR(256),
    priority FLOAT NOT NULL DEFAULT 0,
    need_restart BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    spawned_at TIMESTAMP,
    updated_at TIMESTAMP
);
CREATE TABLE AlgorithmStatus(
    id serial primary key,
    client_id INT REFERENCES clients (id) ON DELETE CASCADE NOT NULL,
    TWAP BOOLEAN NOT NULL DEFAULT FALSE,
    HFT BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE clients;
DROP TABLE AlgorithmStatus;

-- +goose Up
CREATE TABLE AlgorithmStatus(
    id serial primary key,
    client_id INT NOT NULL,
    VWAP BOOLEAN NOT NULL DEFAULT FALSE,
    TWAP BOOLEAN NOT NULL DEFAULT FALSE,
    HFT BOOLEAN NOT NULL DEFAULT FALSE,
);

-- +goose Down
DROP TABLE AlgorithmStatus;

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS currencies
(
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  code VARCHAR(3) NOT NULL UNIQUE,
  symbol VARCHAR(5) NOT NULL
);

CREATE TABLE IF NOT EXISTS countries
(
  id SMALLSERIAL PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  code VARCHAR(3) NOT NULL UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS currencies;
-- +goose StatementEnd

-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles
(
  id SMALLSERIAL PRIMARY KEY,
  role VARCHAR(32) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users
(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role_id SMALLINT NOT NULL,
  email_verified BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (role_id) REFERENCES roles (id)
);
CREATE INDEX IF NOT EXISTS idx_user_role_id ON users (role_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd

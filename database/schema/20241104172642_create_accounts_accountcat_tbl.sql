-- +goose Up
-- +goose StatementBegin
CREATE TYPE account_category_options AS ENUM ('cash', 'bank', 'creditCard', 'debitCard', 'investment', 'loan', 'other');

CREATE TABLE IF NOT EXISTS account_categories
(
  id SMALLSERIAL PRIMARY KEY,
  category account_category_options NOT NULL UNIQUE,
  description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS accounts
(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(255) NOT NULL,
  balance NUMERIC(20, 2) NOT NULL DEFAULT 0,
  description VARCHAR(255),
  account_number VARCHAR(32),
  category_id SMALLINT NOT NULL,
  institution_id SMALLINT,
  user_id UUID NOT NULL,
  currency_id SMALLINT NOT NULL,
  FOREIGN KEY (category_id) REFERENCES account_categories (id),
  FOREIGN KEY (institution_id) REFERENCES institutions (id) ON DELETE SET NULL,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  FOREIGN KEY (currency_id) REFERENCES currencies (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS account_categories;
DROP TYPE IF EXISTS account_category_options;
-- +goose StatementEnd

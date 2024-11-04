-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS institutions
(
  id SMALLSERIAL PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(255) NOT NULL,
  code VARCHAR(8) UNIQUE,
  country_id SMALLINT NOT NULL,
  logo_id UUID,
  user_id UUID NOT NULL,
  FOREIGN KEY (country_id) REFERENCES countries (id),
  FOREIGN KEY (logo_id) REFERENCES attachments (id) ON DELETE SET NULL,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_institution_country_id ON institutions (country_id);
CREATE INDEX IF NOT EXISTS idx_institution_logo_id ON institutions (logo_id);
CREATE INDEX IF NOT EXISTS idx_institution_user_id ON institutions (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS institutions;
-- +goose StatementEnd

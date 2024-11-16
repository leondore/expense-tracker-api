-- name: ListAccounts :many
SELECT
  a.id,
  a.name,
  a.balance,
  a.description,
  a.account_number,
  ac.id AS category_id,
  ac.category,
  i.name AS institution,
  c.code AS currency
FROM accounts AS a
JOIN account_categories AS ac ON a.category_id = ac.id
JOIN institutions AS i ON a.institution_id = i.id
JOIN currencies AS c ON a.currency_id = c.id
WHERE a.user_id = $1
ORDER BY ac.id, a.name;
--

-- name: ReadAccount :one
SELECT
  a.id,
  a.name,
  a.balance,
  a.description,
  a.account_number,
  ac.id AS category_id,
  ac.category,
  i.name AS institution,
  c.code AS currency
FROM accounts AS a
JOIN account_categories AS ac ON a.category_id = ac.id
JOIN institutions AS i ON a.institution_id = i.id
JOIN currencies AS c ON a.currency_id = c.id
WHERE a.id = $1
AND a.user_id = $2;

-- name: CreateAccount :one
INSERT INTO accounts (name, description, account_number, category_id, institution_id, user_id, currency_id)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, name;
--

-- name: UpdateAccount :one
UPDATE accounts
SET name = $2,
description = $3,
account_number = $4,
category_id = $5,
institution_id = $6,
currency_id = $7,
updated_at = NOW()
WHERE id = $1
AND user_id = $8
RETURNING id, name;
--

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
AND user_id = $2;

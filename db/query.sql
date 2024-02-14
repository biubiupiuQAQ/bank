-- name: CreateAccount :execresult
INSERT INTO account (
    account_name,
    balance,
    currency
) VALUES (?,?,?);

-- name: GetAuthor :one
SELECT * FROM account
WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account
WHERE id = ?
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateAccount :execresult
UPDATE account
SET balance = ?
WHERE id = ?;

-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = ?;
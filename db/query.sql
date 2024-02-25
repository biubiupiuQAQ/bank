-- name: CreateAccount :execresult
INSERT INTO account (
    account_name,
    balance,
    currency
) VALUES (?,?,?);

-- name: GetAccount :one
SELECT * FROM account
WHERE id = ? LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM account
WHERE id =? LIMIT 1
FOR UPDATE;

-- name: ListAccounts :many
SELECT * FROM account
WHERE account_name = ?
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

-- name: CreateTransfer :execresult
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
    ?,?,?
);

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = ? 
LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
WHERE 
    from_account_id = ?
    OR
    to_account_id = ?
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: CreateUser :execresult
INSERT INTO users(
    account_id,
    amount
)   VALUES (
    ?,?
);

-- name: GetUser :one
SELECT * FROM users
WHERE id = ?
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE account_id = ?
ORDER BY id
LIMIT ?
OFFSET ?;

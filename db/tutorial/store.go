package tutorial

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// NewStore create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx 执行一个数据库事务
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// 开启一个数据库事务
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 快递事务tx，new一个store对象
	q := New(tx)
	// 执行事务q
	err = fn(q)
	if err != nil {
		// 执行回滚
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("事务 err: %v, 回滚 errL %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams 交易的输入参数
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_accout_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult 交易的结果
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   User     `json:"form entry"`
	ToEntry     User     `json:"to_entry"`
}

// TransferTx 执行账户间的金钱交易
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// 对行加锁
		From_account, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		transfer_result, err := q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		resinfo, err := transfer_result.LastInsertId()
		if err != nil {
			return err
		}
		result.Transfer, err = q.GetTransfer(ctx, resinfo)

		from_entry := CreateUserParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		}

		user_result, err := q.CreateUser(ctx, from_entry)
		if err != nil {
			return err
		}
		fromentry_id, err := user_result.LastInsertId()
		if err != nil {
			return err
		}
		result.FromEntry = User{
			ID:        fromentry_id,
			AccountID: from_entry.AccountID,
			Amount:    from_entry.Amount,
		}

		to_entry := CreateUserParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		}
		user_result1, err := q.CreateUser(ctx, to_entry)
		if err != nil {
			return err
		}
		toentry_id, err := user_result1.LastInsertId()
		if err != nil {
			return err
		}
		result.ToEntry = User{
			ID:        toentry_id,
			AccountID: to_entry.AccountID,
			Amount:    to_entry.Amount,
		}

		// 更新账户余额

		_, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: From_account.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromAccount, err = q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		To_account, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}
		_, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: To_account.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToAccount, err = q.GetAccount(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

package tutorial

import (
	"context"
	"log"
	"testing"

	"github.com/biubiupiuQAQ/bank/tree/master/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	ctx := context.Background()
	arg := CreateAccountParams{
		AccountName: util.RandomName(),
		Balance:     util.RandomMoney(),
		Currency:    util.RandomCurrency(),
	}
	account_ob, err := testQueries.CreateAccount(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, account_ob)
	account_id, err := account_ob.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	account, err := testQueries.GetAccount(ctx, account_id)
	if err != nil {
		log.Fatal(err)
	}
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	var id int64 = 3
	account, err := testQueries.GetAccount(context.Background(), id)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("id: ", account.ID, ",name: ", account.AccountName, ",balance: ", account.Balance, ",currency: ", account.Currency, ",time: ", account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	arg := UpdateAccountParams{
		ID:      4,
		Balance: util.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
		return
	}
	require.NotEmpty(t, account)
}

// func TestDeleteAccount(t *testing.T) {
// 	var id int64 = 6
// 	err := testQueries.DeleteAccount(context.Background(), id)
// 	require.NoError(t, err)
// }

func TestListAccount(t *testing.T) {
	arg := ListAccountsParams{
		AccountName: "alen",
		Limit:       5,
		Offset:      0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	n := len(accounts)
	for i := 0; i < n; i++ {
		log.Println("ID: ", accounts[i].ID, ",name: ", accounts[i].AccountName, ",balance: ", accounts[i].Balance, ",currency: ", accounts[i].Currency, ",time: ", accounts[i].CreatedAt)
	}

}

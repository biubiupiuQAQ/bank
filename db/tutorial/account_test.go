package tutorial

import (
	"context"
	"testing"

	"BANK/db/util"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	ctx := context.Background()
	arg := CreateAccountParams{
		AccountName: util.RandomName(),
		Balance:     util.RandomMoney(),
		Currency:    util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
}

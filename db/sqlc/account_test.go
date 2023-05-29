package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/foyez/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func deleteAccountUtil(t *testing.T, id int64) {
	err := testQueries.DeleteAccount(context.Background(), id)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestCreateAccount(t *testing.T) {
	account := createRandomAccount(t)

	t.Cleanup(func() {
		deleteAccountUtil(t, account.ID)
	})
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	t.Cleanup(func() {
		deleteAccountUtil(t, account1.ID)
	})
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account2.Balance, arg.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	t.Cleanup(func() {
		deleteAccountUtil(t, account1.ID)
	})
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	deleteAccountUtil(t, account1.ID)
}

func TestListAccounts(t *testing.T) {
	var createdAccounts []Account
	for i := 0; i < 10; i++ {
		account := createRandomAccount(t)
		createdAccounts = append(createdAccounts, account)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

	t.Cleanup(func() {
		for _, account := range createdAccounts {
			deleteAccountUtil(t, account.ID)
		}
	})

}

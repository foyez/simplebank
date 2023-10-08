package db

import (
	"context"
	"testing"
	"time"

	"github.com/foyez/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testStore.CreateAccount(context.Background(), arg)
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
	err := testStore.DeleteAccount(context.Background(), id)
	require.NoError(t, err)

	account2, err := testStore.GetAccount(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
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
	account2, err := testStore.GetAccount(context.Background(), account1.ID)
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

	account2, err := testStore.UpdateAccount(context.Background(), arg)
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
	var lastAccount Account

	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
		createdAccounts = append(createdAccounts, lastAccount)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	// require.Len(t, accounts, 5)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, account.Owner, lastAccount.Owner)
	}

	t.Cleanup(func() {
		for _, account := range createdAccounts {
			deleteAccountUtil(t, account.ID)
		}
	})

}

package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		want := Bitcoin(20)

		fmt.Printf("address of bitcoin in test is %p \n", &wallet.balance)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(Bitcoin(2))

		want := Bitcoin(8)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalnce := Bitcoin(20)
		wallet := Wallet{startingBalnce}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalnce)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted and error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

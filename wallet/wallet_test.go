package wallet_test

import (
	"testing"

	"github.com/MarkMoelter/learn-golang/wallet"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet_ := wallet.Wallet{}
		wallet_.Deposit(wallet.Bitcoin(10))

		assertBalance(t, wallet_, wallet.Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T){
		startingBalance := wallet.Bitcoin(20)
		wallet_ := wallet.Wallet{}
		wallet_.Deposit(startingBalance)
		err := wallet_.Withdraw(wallet.Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet_, wallet.Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := wallet.Bitcoin(20)
		wallet_ := wallet.Wallet{}
		wallet_.Deposit(startingBalance)
		err := wallet_.Withdraw(wallet.Bitcoin(100))

		assertError(t, err, wallet.ErrInsufficientFunds)
		assertBalance(t, wallet_, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet wallet.Wallet, want wallet.Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t. Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

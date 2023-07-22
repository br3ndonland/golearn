package pointers

import (
	"fmt"
	"testing"
)

func assertCorrectBalance(t testing.TB, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertCorrectBalanceString(t testing.TB, got, want Bitcoin) {
	t.Helper()
	gotString := got.String()
	wantString := fmt.Sprintf("%d BTC", got)
	if gotString != wantString {
		t.Errorf("String format incorrect: got %s want %s", gotString, wantString)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertCorrectBalance(t, got, want)
		assertCorrectBalanceString(t, got, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertCorrectBalance(t, got, want)
		assertCorrectBalanceString(t, got, want)
	})
}

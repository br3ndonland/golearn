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
	gotString := got.String()
	wantString := fmt.Sprintf("%d BTC", got)
	if gotString != wantString {
		t.Errorf("String format incorrect: got %s want %s", gotString, wantString)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("error expected")
	}
	gotErr := got.Error()
	wantErr := want.Error()
	if gotErr != wantErr {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		got := wallet.Balance()
		want := Bitcoin(10)
		assertCorrectBalance(t, got, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertCorrectBalance(t, got, want)
	})
	t.Run("withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		withdrawalAmount := Bitcoin(100)
		wallet := Wallet{balance: startingBalance}
		got := wallet.Withdraw(withdrawalAmount)
		want := wallet.ErrInsufficientFunds(withdrawalAmount)
		assertError(t, got, want)
	})
}

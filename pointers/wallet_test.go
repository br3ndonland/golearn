package pointers

import "testing"

func assertCorrectBalance(t testing.TB, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := Bitcoin(10)
	assertCorrectBalance(t, got, want)
}

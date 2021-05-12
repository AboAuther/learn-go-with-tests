package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.balance

		if got != want {
			t.Errorf("got:%s,want:%s", got, want)
		}
	}
	assertErrors := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdrew with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		_ = wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdrew insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startBalance)
		assertErrors(t, err)
	})
}

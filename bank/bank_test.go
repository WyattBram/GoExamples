package bank

import "testing"

func TestFunctionalities(t *testing.T) {

	assertBalance := func(t *testing.T, want Bitcoin, w Wallet) {
		got := w.Balance()
		if got != want {
			t.Errorf("Got: %s, Want: %s", got, want)
		}
	}

	assertError := func(t *testing.T, err, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("Wanted an error, didnt get one")
		} else if want != err {
			t.Errorf("Got: %q, Want %q", err, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("Did not want an error, got one")
		}
	}
	t.Run("Test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(5))

		want := Bitcoin(5)

		assertBalance(t, want, wallet)

	})

	t.Run("Test Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(5))

		want := Bitcoin(1)
		err := wallet.Withdraw(Bitcoin(4))

		assertNoError(t, err)
		assertBalance(t, want, wallet)

	})

	t.Run("Test Overdraft", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		err := wallet.Withdraw(Bitcoin(20))

		want := Bitcoin(10)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, want, wallet)
	})
}

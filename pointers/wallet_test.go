package pointers

import "testing"

func TestWallet(t *testing.T)  {

  assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()
    if got != want {
      t.Errorf("got %s, want %s", got, want)
    }
  }

  assertError := func (t testing.TB,got error, want string) {
    t.Helper()
    if got == nil {
      t.Fatal("didn't get an error but wanted one")
    }

    if got.Error() != want {
      t.Errorf("got %q want %q", got, want)
    }
    
  }

  t.Run("deposit", func(t *testing.T) {
    wallet := Wallet{}

    wallet.Deposit(Bitcoin(10))
    assertBalance(t, wallet, Bitcoin(10))

  })
  t.Run("withdraw", func(t *testing.T) {
    wallet := Wallet{bitcoin: Bitcoin(10)}

    err := wallet.Withdraw(Bitcoin(5))

    assertError(t, err, "Not enough balance")
    assertBalance(t, wallet, Bitcoin(5))

  })
}
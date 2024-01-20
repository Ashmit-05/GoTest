package pointers

import "testing"

func TestWallet(t *testing.T)  {
  t.Run("deposit", func(t *testing.T) {
    wallet := Wallet{}
    wallet.Deposit(Bitcoin(10))

    assertBalance(t, wallet, Bitcoin(10))
  })

  t.Run("withdraw with funds", func(t *testing.T) {
    wallet := Wallet{bitcoin: Bitcoin(10)}
    err := wallet.Withdraw(Bitcoin(5))

    assertNoError(t,err)
    assertBalance(t, wallet, Bitcoin(5))
  })

  t.Run("withdraws insufficient funds", func(t *testing.T) {
    wallet := Wallet{bitcoin: Bitcoin(10)}
    err := wallet.Withdraw(Bitcoin(50))

    assertError(t,err,ErrInsufficientFunds)
    assertBalance(t, wallet, Bitcoin(10))
  })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()
    if got != want {
      t.Errorf("got %s, want %s", got, want)
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
  t.Helper()

  if got != nil {
    t.Fatal("got an error but didn't want one")
  }
}

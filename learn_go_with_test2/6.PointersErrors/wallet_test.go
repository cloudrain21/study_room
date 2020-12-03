package main

import (
    "fmt"
    "testing"
)

func TestWallet(t *testing.T) {

    commonAssert := func(t *testing.T, got, want interface{}) {
        t.Helper()
        if got != want {
            switch got.(type) {
                case Bitcoin:
                    t.Errorf("got (%d) want (%d)", got, want)
                case string:
                    t.Errorf("got (%s) want (%s)", got, want)
            }
        }
    }

    assertRefactor := func(t *testing.T, wallet Wallet, want Bitcoin) {
        t.Helper()
        got := wallet.Balance()

        if got != want {
            t.Errorf("got (%d) want (%d)", got, want)
        }
    }

    assertError := func(t *testing.T, err error, want error) {
        t.Helper()

        if err == nil {
            t.Fatal("noooooo. error must not be nil")
        }

        if err != want {
            t.Errorf("got (%s) want (%s)", err, want)
        }
    }

    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(100))

        assertRefactor(t, wallet, Bitcoin(100))
    })

    t.Run("string", func(t *testing.T) {
        wallet := Wallet{Bitcoin(100)}
        wallet.Deposit(Bitcoin(100))

        got := wallet.Balance().String()
        want := "200 BTC"

        commonAssert(t, got, want)
    })

    t.Run("withdraw", func(t *testing.T) {
        wallet := Wallet{Bitcoin(200)}
        _ = wallet.Withdraw(100)

        got := wallet.Balance().String()
        want := "100 BTC"

        commonAssert(t, got, want)

        got1 := wallet.Balance()
        want1 := Bitcoin(100)

        commonAssert(t, got1, want1)
    })

    t.Run("withdraw over", func(t *testing.T) {
        wallet := Wallet{Bitcoin(100)}
        err := wallet.Withdraw(200)

        assertRefactor(t, wallet, Bitcoin(100))
        assertError(t, err, ErrInsufficientBalance)
    })
}

func ExampleWallet() {
    w := Wallet{Bitcoin(100)}
    w.Deposit(Bitcoin(100))
    fmt.Println(w.Balance())

    // Output:
    // 200 BTC
}

func BenchmarkWallet(b *testing.B) {
    wallet := Wallet{Bitcoin(0)}
    for i:=0; i<b.N; i++ {
        wallet.Deposit(1)
        _ = wallet.Withdraw(1)
    }
    fmt.Println(wallet.Balance())
}

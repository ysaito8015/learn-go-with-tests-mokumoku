# Pointers & errors

## Write the test first

```go
package main

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```

## Try to run the test

```shell
# learn-go-with-tests/pointers_errors/01_wallet [learn-go-with-tests/pointers_errors/01_wallet.test]
./wallet_test.go:7:12: undefined: Wallet
FAIL    learn-go-with-tests/pointers_errors/01_wallet [build failed]

```

## Write the minimal amout of code for the test to run and check the failing test output

```go
package main

type Wallet struct {
}

func (w Wallet) Balance() int {
	return 0
}

```

## Write enough code to make it pass

```go
```


## ????

```go
package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("memory address of balance in Balance() is %v \n", &w.balance)
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("memory address of balance in Deposit() is %v \n", &w.balance)
	w.balance += amount
}
```


## Refactor

```go
package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
```

```go
package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
```


## Write the test first

```go
package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdoraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
```


## Try to run the test


## Write the minimal amount of code for the test to run and check the failingtest output


## Write enough code to make it pass


## Refactor
- `Withdraw` 関数を口座の残高以上実行したらどうなるか
- マイナスになってしまう
- 関数呼び出し元に, `err` を返すことでチェックとハンドリングができる


```go
package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdoraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
}
```


## Write the test first


```go
package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraow insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		if err == nil {
			t.Error("wanted an error but didin't get one")
		}
	})
}
```


## Try and run the test


## Write the minimal amout of code for the test to run and check the failing test output


```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.balance -= amount
	return nil
}
```


## Write enough code to make it pass
- `errors.New()` 関数は, エラーメッセージとともに, 新しい `error` 型を作成する


```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}
```


## Refactor
- テストヘルパー関数を作る
- エラーがユーザーに返されるとして, 有効なエラーメッセージを返すように変える


```go
	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didin't get one")
		}
	}

```


## Write the test first




## Try to run the test


## Write ecnough code to make it pass


## Refactor


```go
package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraow insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got.Error() != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
```


### Unchecked errors
- `errcheck` パッケージを利用して, コンパイラが指摘しないエラーを見つける

```shell
$ go install github.com/kisielk/errcheck

$ errcheck .
wallet_test.go:17:18:   wallet.Withdraw(Bitcoin(10))
```


```go
package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraow insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didin't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
```


## Wrapping up
### Pointers
- Go は 関数やメソッドに, 値を渡すときにそのコピーを作成する
    - 値を変更する関数を書くときは, そのポインタが必要
- とても大きいデータ構造の参照や, たったひとつのインスタンスのみが必要なケース(sqlx など)などのときに値コピーでなくポインタが必要


### nil
- ポインタは nil になれる
- ポインタが渡されたときに, nil でないかどうかを確認する


### Errors
- 関数やメソッドを呼び出したときに, 失敗を示す方法
- "Don't just checck errors, handle them gracefully"
		- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully


### Create new times from existing ones
- 特定用途に使う値を追加するときに有用

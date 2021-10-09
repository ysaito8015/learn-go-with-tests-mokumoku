# Hello, World

## How it works

## How to test
- ドメインコードを切り分ける
    - `fmt.Println("Hello, World")` -> `func Hello() string {}` と `fmt.Println(Hello())` へ
- `hello_test.go` を新たに作成する
    - `got`, `want` などの変数を活用して, 再利用できるようにする

## Go modules?
- `go test`
    - Go 1.16 以降の場合は, エラーメッセージが現れる
- `go mod init hello` を実行


## Back to Testing
### Writing tests
- ファイル名は `foo_test.go`
- テスト関数は `TestBar()` と始める
- テスト関数の引数は `t *testing.T` ひとつだけ
    - "hook" into the testing framework
- `testing` パッケージをインポートする
#### if
#### Declaring variables
#### t.Errorf
- `fmt.Printf` みたいなもの
- `%q`: 文字列などが `""` で囲まれて出力される


### Go doc
- Go 1.14 以降では `godoc` コマンドはインストールが必要
- `go get golang.org/x/tools/cmd/godoc`


### Hello, YOU
- ここから, 本体コードよりテストを先に書いていく
- 挨拶の受け取り手を特定していく


### A note on source control
- ここまでで, 一度コミットしておく
- main にはプッシュせず, リファクタリングに備える


### Constants
- 定数を導入する


## Hello, world... again
- まずテストを書く
- `t.Run` 関数
    - テストをシナリオごとにグループ化する
    - サブテスト
- `func (t *T) Run(name string, f func(t *T)) bool`
    - 引数 `f` で渡された関数を実行する
    - 別の goroutine で実行される
        - `t.Parallel()` の実行が必要
- リファクタリングは, プロダクションコードのためのもの "だけ" ではない
- テストにたいして何がコードに求められているか要求事項を明確にすることが重要
- ヘルパー関数を導入
    - `t.Helper()` の宣言がヘルパーメソッドの定義内で必要
        - コメントアウトすると, エラーが発生した行数ではなく, ヘルパーメソッド内で定義した行数が返されてしまう
    - `t *testing.TB` 引数: ヘルパー関数がテストとベンチマーク両方で利用できる

```
package testing // import "testing"

func (b *B) Run(name string, f func(b *B)) bool
    Run benchmarks f as a subbenchmark with the given name. It reports whether
    there were any failures.

    A subbenchmark is like any other benchmark. A benchmark that calls Run at
    least once will not be measured itself and will be called once with N=1.

func (m *M) Run() (code int)
    Run runs the tests. It returns an exit code to pass to os.Exit.

func (t *T) Run(name string, f func(t *T)) bool
    Run runs f as a subtest of t called name. It runs f in a separate goroutine
    and blocks until f returns or calls t.Parallel to become a parallel test.
    Run reports whether f succeeded (or at least did not fail before calling
    t.Parallel).

    Run may be called simultaneously from multiple goroutines, but all such
    calls must return before the outer test function for t returns.


```


### Backto source control
- ひとつ前のコミットを `amend` する


### Discipline
- もう一度このサイクルを実行する
    1. テストを書く
    2. コンパイラを通す
    3. テストを実行する
        - 失敗することを確認する
        - エラーメッセージが意味のあるものかを確認する
    4. テストをパスするためにプロダクトコードを書く
    5. リファクタリングする
- つまらないものにみえるかも知れないが, このフィードバックループを堅持することが重要
- 関連するテストを確かなものにすることだけでなく, 安全なテストとともにリファクタリングすることで, グッドソフトウェアをデザインすることを確かなものにする
- テストが失敗することを確認することは, 重要. なぜなら, それは, エラーメッセージがどのようにみられるかを確認できるから.
- テストが失敗し, 何が問題かというはっきりとしたアイディアをもたらさない時, 開発者としては, コードベースで仕事することが困難になり得る
- テストを実行することを単純にするためにツールを設定すること, そして,  テストが速いことを確かなものにすることによって, コードを書いているときにフロー状態に入ることができる


## Keep going! More requirements
- さらなる要求
    - 第二引数
    - 挨拶の言語を特定する
        - デフォルトは英語

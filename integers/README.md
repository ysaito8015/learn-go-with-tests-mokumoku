# Intergers
- `Add` 関数を書くことで理解する

## Write the test first
- `Errorf()` 関数のフォーマット指定子 `%d` を使う
    - `string` ではなく `int` を出力するということを明確にするため

## Try and run the test


## Write the minimal amount of code for the test to run and check the failing test output
- まずコンパイラを満足させるコードを書く
- `(x int y int)` と書くより `(x, y int)` とかく
- 名前付き戻り値を使っていない
    - 名前付き戻り値を使うのは, 文脈から結果の意味が明確でないときに使う


## Write enough code to make it pass
- `return 4` してテストをパスコードを書く
- 更にテストにケースを追加できるが, いたちごっこになる
- Property Based Testing
    - 後で出てくる


## ReFactor
- 関数にコメントを追加する
- それらのコメントは Go Doc に現れる


## Examples
- 作った関数の example を追加することができる
- `Foo_test.g` の中に書くことで, Go Doc に表示される
- `ExampleAdd()` 関数内のコメント, `// Output: 6` を削除すると, Example 関数は実行されない
- `godoc` コマンドは, ローカルで Go のドキュメントを表示するサーバを起動する
    - デフォルトは `loclhost:6060`


## Wrapping up
- より良いドキュメントを書いた

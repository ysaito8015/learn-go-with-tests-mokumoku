# Arrays and slices
- 配列は, 同じ型の多数の要素を, 変更可能かつ特定の並びで格納する


## Write the test first
- 配列の表現
  - `[要素数]型{初期化リテラル}`
  - `[...]型{初期化リテラル}`


## Try to run the test


## Write the minimal amount of code for the test to run and check the failing test output

## Write enough code to make it pass

## Refactor
- `range` キーワードを使う

## Write the test first

## Try and run the test

## Write the minimal amount of code for the test to run and check the failing test output


## Write enough code to make it pass

## Refactor
- `sum_test.go` をリファクタリング
- 多くのテストケースを用意することがゴールではない
- むしろ, コードへの信用を可能な限り高めることである
- テストが多すぎると, 実際の問題になり得る
- また, メンテナンスによりオーバーヘッドをただ加えるに過ぎない
- **すべてのテストはコストを持つ**
- 今回のテストでは, `Sum` 関数に２つのテストケースが用意されていおり, 重複している

### coverage tool
- 組み込みのテストツールにはカバレッジツールが用意されている
- カバレッジ 100% を達成することがゴールではないにしろ
- coverage tool はテストによってカバーされていないコードの領域を特定すること助けることができる


### New function
- `SumAll`
    - スライスを複数引数に取ることができる
    - スライスごとに要素を足し合わせ, スライスとして返す


## Write the test first

## Try and run the test

## Write the minimal amount of code for the test to run and check the failing test output
- 可変長引数, `...型`, variable length arguments, variadic arguments

### invalid operation
- スライスの比較は nil とのみ許されている
- `./sum_test.go:35:9: invalid operation: got != want (slice can only be compared to nil)`
- `reflect.DeepEqual` を２つの値を比較するために利用する
    - "type safe" ではない
    - スライスと文字列を比較するような無意味な評価もできてしまう

```
package reflect // import "reflect"

func DeepEqual(x, y interface{}) bool
    DeepEqual reports whether x and y are “deeply equal,” defined as follows.
    Two values of identical type are deeply equal if one of the following cases
    applies. Values of distinct types are never deeply equal.

    Array values are deeply equal when their corresponding elements are deeply
    equal.

    Struct values are deeply equal if their corresponding fields, both exported
    and unexported, are deeply equal.

    Func values are deeply equal if both are nil; otherwise they are not deeply
    equal.

    Interface values are deeply equal if they hold deeply equal concrete values.

    Map values are deeply equal when all of the following are true: they are
    both nil or both non-nil, they have the same length, and either they are the
    same map object or their corresponding keys (matched using Go equality) map
    to deeply equal values.

    Pointer values are deeply equal if they are equal using Go's == operator or
    if they point to deeply equal values.

    Slice values are deeply equal when all of the following are true: they are
    both nil or both non-nil, they have the same length, and either they point
    to the same initial entry of the same underlying array (that is, &x[0] ==
    &y[0]) or their corresponding elements (up to length) are deeply equal. Note
    that a non-nil empty slice and a nil slice (for example, []byte{} and
    []byte(nil)) are not deeply equal.

    Other values - numbers, bools, strings, and channels - are deeply equal if
    they are equal using Go's == operator.

    In general DeepEqual is a recursive relaxation of Go's == operator. However,
    this idea is impossible to implement without some inconsistency.
    Specifically, it is possible for a value to be unequal to itself, either
    because it is of func type (uncomparable in general) or because it is a
    floating-point NaN value (not equal to itself in floating-point comparison),
    or because it is an array, struct, or interface containing such a value. On
    the other hand, pointer values are always equal to themselves, even if they
    point at or contain such problematic values, because they compare equal
    using Go's == operator, and that is a sufficient condition to be deeply
    equal, regardless of content. DeepEqual has been defined so that the same
    short-cut applies to slices and maps: if x and y are the same slice or the
    same map, they are deeply equal regardless of content.

    As DeepEqual traverses the data values it may find a cycle. The second and
    subsequent times that DeepEqual compares two pointer values that have been
    compared before, it treats the values as equal rather than examining the
    values to which they point. This ensures that DeepEqual terminates.

```

## Write enough code to make it pass


## Refactor
- スライスの cap が 2 なのに `mySlice[10] = 1` のような代入はランタイムエラーを起こす
- `append` 関数を使う


### Next requirement
- `SumAll` を `SumAllTails` に変える
    - 各スライスの末尾の合計を計算する
    - 各スライスの最初の要素 (head) 以外の合計の集合が tail の集合


## Write the test first

## Try and run the test

## Write the minimal amount of code for the test to run and check the failing test output

## Write enough code to make it pass
- `n[1:]` インデックスで指定する
    - 1 以上, からスライスの終わりまで

## Refactor
- 空のスライスをわたしたときに, 何が起きるか
- 末尾が空のスライスの場合は？

## Write the test first

## Try and run the test

```
panic: runtime error: slice bounds out of range [1:0] [recovered]
        panic: runtime error: slice bounds out of range [1:0]
```

## Write enough code to make it pass

## Refactor


## Wrapping up
- 配列
- スライス
    - いろいろな方法で宣言できる
    - 固定の cap を持っているが, `append` 関数で要素を追加できる
    - スライスのスライスの仕方
- `len`  関数
- test coverage tool
- `reflect.DeepEqual` 関数
- 配列から作ったスライスと, 配列から `copy` 関数で作った配列の違い
    - https://play.golang.org/p/bTrRmYfNYCp
- len が巨大なスライスのコピーの作成方法の例示
    - https://play.golang.org/p/GGK9IBeOy4M
- The Go Blog: Go Slices: usage and internals, 5 Jan. 2011
    - https://go.dev/blog/slices-intro
        - スライスとは, 配列のセグメントのディスクリプタ (要約)
        - スライスは, 配列へのポインタ, 長さ, 容量の要素を持つ構造体
        - スライスからスライスを作る時, スライスの構造体のデータをコピーしない
        - スライスは, cap を超えて大きくなれない
        - cap を超えて要素が追加される時, スライスのコピーが作られて, (元の cap +1) *2 の cap のスライスが作られる

# Structs, methods & interfaces

## Write the test first

```go
package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("fot %.2f want %.2f", got, want)
	}
}
```

## Try the test first

```shell
$ go test -v --cover
# shapes [shapes.test]
./shapes_test.go:6:9: undefined: Perimeter
FAIL    shapes [build failed]

```

## Write the minimal amount of code for the test to run and check the failing test output

```go
package main

func Perimeter(h, w float64) float64 {
	return 0
}
```

## Write enough code to make it pass

```go
package main

func Perimeter(h, w float64) float64 {
	p := 2 * (h + w)
	return p
}
```


```go
package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10, 0)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```


```go
package main

func Perimeter(h, w float64) float64 {
	return 2 * (h + w)
}

func Area(h, w float64) float64 {
	return h * w
}
```


## Refactor
- use struct for specific shape


```go
package main

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```


```go
package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{Width: 12.0, Height: 6.0}
	got := rectangle.Area()
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
```

## Write the test first
- `Area` function for circle
- `fmt` フォーマット指定子
    - `%g`: 浮動小数点用, `%e` を指数部が大きいときに使う, `%f` はそれ以外の時


```go
package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("calculate area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("calculate area of circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
```

## Try to run the test

```shell
$ go test -v --cover
# shapes [shapes.test]
./shapes_test.go:28:13: undefined: Circle
FAIL    shapes [build failed]
```

## Write the minimal smount of code for the test to run and check the failing test output

```go
package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
```

## Write the minimal amount of code for the test to run and check the failing test output

## Write enough code to make it pass

## Refactor
- `interface`


```go
package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %f", got, want)
		}
	}

	t.Run("calculate area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("calculate area of circle", func(t *testing.T) {
		circle := Circle{Radius: 10}
		checkArea(t, circle, 314.1592653589793)
	})
}
```

```go
package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
```


### Decoupling
- 切り離し
- テストケースに定義した `checkArea` ヘルパー関数では, shape が, `Rectangle`, `Circle`, `Triangle` のいずれでもかまわない
- interface を宣言することは, ヘルパー関数を特定の型から切り離し, 動作に必要なメソッドのみを持つ.


## Further refactoring
- "Table Driven Tests"
    - https://github.com/golang/go/wiki/TableDrivenTests

## Write the test first

## Try to run the test

## Write the minimal amount of code for the test to run and check the failing

## Write enough code to make it pass

## Refactor
- Test-Driven Development by Example
		- https://www.oreilly.com/library/view/test-driven-development/0321146530/
> The test speaks to us more clearly, as if it were an assertion of truth, **not a sequence of operations**
- テストは,  **一連の操作ではなく** 真実の主張であるかのように, より明確に我々に語りかける

## Make sure your test output is helpful
- テーブル定義した構造体の一部のテストケースのみを実行する

```shell
$ go test -v --cover -run TestArea/Rectangle
```


## Wrapping up
- 関連データをまとめて, コードの意図を明確にする独自のデータ型を作成するための構造体を宣言する
- 違う型で使用できる関数を定義できるようにインタフェースを宣言する
- データ型に機能的に追加できるように, また, インタフェースを実装できるようにメソッドを追加する
- テストの主張をより明確にし, テストスーツをより拡張しやすく, メンテナンスしやすくするために, テーブルドリブンテストを行う

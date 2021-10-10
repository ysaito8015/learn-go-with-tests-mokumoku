# Iteration
- `for` 文を使って繰り返しを表現する
- 一文字を５回繰り返す関数を作る


## Try and run the test

## Write the minimal amount of code for the test to run and check the failing test output
- 

## Write enough code to make it pass

```
func Repeat(character string) string {
	str := make([]string, 5)
	for i := 0; i < 5; i++ {
		str = append(str, character)
	}
	return fmt.Sprint(strings.Join(str, ""))
}
```


```
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return fmt.Sprint(repeated)
}
```

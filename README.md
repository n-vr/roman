# Roman

This package can convert roman numerals to decimal.

## Install

```shell
go get github.com/n-vr/roman
```

## Example usage

This a a simple go program that shows how to use this package.

```go
import (
    "fmt"

    "github.com/n-vr/roman"
)

func main() {
    value, err := roman.RomanToDecimal("MMXXIV")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
}
```
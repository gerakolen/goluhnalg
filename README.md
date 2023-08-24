# Introduction

Goluhnalg - its a package for validate, calculate and generate [luhn number](https://en.wikipedia.org/wiki/Luhn_algorithm).

# Prerequisites

- You need to have Go installed on your computer. The version used in this project is 1.20.

# Usage

## Validation

```go
err := goluhn.Validate("1111222233334444")
if err != nil {
  return err
}
```

## Generation

```go
n, err := goluhn.Generate(9)
if err != nil {
  return err
}

fmt.Printf("Luhn number: %s\n", n)
```

## Calculation

```go
checkDigit, LuhnNum, err := goluhn.Calculate("7992739871")
if err != nil {
  return err
}

fmt.Printf("Check digit: %s\n", checkDigit)
fmt.Printf("Luhn number: %s\n", LuhnNum)
```
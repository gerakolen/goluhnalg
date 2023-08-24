package goluhnalg

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const ()

func Validate(number string) error {
	sum, err := calculateSum(number, len(number)%2)
	if err != nil {
		return err
	}

	if sum%10 != 0 {
		return errors.New("invalud number")
	}

	return nil
}

func Calculate(number string) (string, string, error) {
	parity := (len(number) + 1) % 2
	sum, err := calculateSum(number, parity)
	if err != nil {
		return "", "", nil
	}

	luhn := sum % 10
	if luhn != 0 {
		luhn = 10 - luhn
	}

	return strconv.FormatInt(luhn, 10), fmt.Sprintf("%s%d", number, luhn), nil
}

func Generate(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	var str strings.Builder
	for i := 0; i < length-1; i++ {
		randomNumber := strconv.Itoa(rand.Intn(9))
		str.WriteString(randomNumber)
	}

	_, res, _ := Calculate(str.String())
	return res
}

func calculateSum(number string, parity int) (int64, error) {
	var (
		sum          int64
		parsedNumber = regexp.MustCompile(`[[:space:]|\p{Z}]`).ReplaceAllString(number, "")
	)

	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(parsedNumber[i]))
		if digit < 0 || digit > 10 {
			return 0, errors.New("invalid digit")
		}
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += int64(digit)
	}
	return sum, nil
}

package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// ErrInvalidString ошибка некорректной строки.
var ErrInvalidString = errors.New("invalid string")
var builder = strings.Builder{}

// Добавление в итоговую строку символа.
func add(symbol int32, multiplier int) {
	builder.WriteString(strings.Repeat(string(symbol), multiplier))
}

// Сброс итоговой строки.
func reset() {
	builder.Reset()
}

// Unpack - Распаковка строки.
func Unpack(text string) (string, error) {
	var symbol int32
	var multiplier int
	var operation bool

	reset()

	for index, runeValue := range text {
		if !unicode.IsDigit(runeValue) {
			if operation {
				add(symbol, 1)
			} else {
				operation = true
			}
			symbol = runeValue
			if index+1 == len(text) {
				add(symbol, 1)
			}
		} else {
			if !operation {
				return "", ErrInvalidString
			}

			multiplier, _ = strconv.Atoi(string(runeValue))
			add(symbol, multiplier)
			operation = false
		}
	}

	return builder.String(), nil
}

package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// Структура команды.
type command struct {
	symbol     int32
	multiplier int
}

// ErrInvalidString ошибка некорректной строки.
var ErrInvalidString = errors.New("invalid string")

// Обработка команды.
func process(command command) string {
	return strings.Repeat(string(command.symbol), command.multiplier)
}

// Добавление команды в коллекцию.
func addCommand(list []command, symbol int32, multiplier int) []command {
	return append(list, command{symbol, multiplier})
}

// Разбор строки на команды.
func parseToCommands(text string) ([]command, error) {
	var list []command
	var operation bool
	var symbol int32

	for _, runeValue := range text {
		isDigit := unicode.IsDigit(runeValue)

		// Если формат строки некорректный.
		if isDigit && !operation {
			return list, ErrInvalidString
		}

		if isDigit {
			multiplier, _ := strconv.Atoi(string(runeValue))
			list = addCommand(list, symbol, multiplier)
			operation = false
			continue
		}

		if operation {
			list = addCommand(list, symbol, 1)
		} else {
			operation = true
		}

		symbol = runeValue
	}

	// Если последняя операция не завершена.
	if operation {
		list = addCommand(list, symbol, 1)
	}

	return list, nil
}

// Сборка строки.
func assemble(list []command) (string, error) {
	result := strings.Builder{}

	for _, command := range list {
		result.WriteString(process(command))
	}

	return result.String(), nil
}

// Unpack - Распаковка строки.
func Unpack(text string) (string, error) {
	// Разбор строки на команды.
	commands, err := parseToCommands(text)
	// Если формат строки некорректный.
	if err != nil {
		return "", err
	}

	// Сборка новой строки.
	return assemble(commands)
}

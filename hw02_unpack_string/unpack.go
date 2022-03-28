package main

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
	var isShield bool
	var symbol int32

	for _, runeValue := range text {
		isDigit := unicode.IsDigit(runeValue)

		// Если формат строки некорректный.
		if isDigit && !isShield && !operation {
			return list, ErrInvalidString
		}

		// Если символ - неэкранированная цифра, в коллекцию добавляется
		// новая пара символ - множитель и текущая операция закрывается.
		if isDigit && !isShield {
			multiplier, _ := strconv.Atoi(string(runeValue))
			list = addCommand(list, symbol, multiplier)
			operation = false
			continue
		}

		// Если операция уже открыта, в коллекцию добавляется
		// новая пара символ - 1.
		if operation {
			list = addCommand(list, symbol, 1)
			operation = false
		}

		// Если символ - неэкранированный слеш,
		// значит символ и есть экран следующего символа.
		if !isShield && string(runeValue) == `\` {
			isShield = true
			continue
		} else {
			isShield = false
		}

		// Открытие новой операции
		// и запоминание текущего символа.
		operation = true
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

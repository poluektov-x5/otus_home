package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// TopSize - Размер топ-листа.
const TopSize = 10

var regular = regexp.MustCompile("[^A-Za-zА-Яа-я0-9-\\s]+")

// Top10 - получение списка топ-10.
func Top10(text string) []string {
	// Форматирование оригинального текста.
	text = format(text)

	// Преобразование текста в оригинальный список слов (с повторами).
	list := strings.Fields(text)

	// Фильтрация оригинального списка слов (с повторами).
	list = filterList(list)

	// Преобразование текста в карту с количеством повторений.
	rangeList := convertToRangeMap(list)

	// Преобразование карты с количеством повторений в список слов (без повторов).
	list = convertRangeMapToList(rangeList)

	// Сортировка списка по частоте и по алфавиту.
	sortByRangeByAbc(list, rangeList)

	return limit(list, TopSize)
}

// Форматирование оригинального текста.
func format(text string) string {
	// Удаление знаков препинания.
	text = regular.ReplaceAllString(text, "")

	// Преобразование к нижнему регистру.
	text = strings.ToLower(text)

	return text
}

// Фильтрация оригинального списка слов (с повторами).
func filterList(list []string) []string {
	filteredList := make([]string, 0, len(list))
	for _, item := range list {
		if item != "-" {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}

// Преобразование оригинального списка слов в карту с количеством повторений.
func convertToRangeMap(list []string) map[string]int {
	// Создание карты слов с количеством повторений.
	rangeList := make(map[string]int)
	for _, item := range list {
		rangeList[item]++
	}

	return rangeList
}

// Преобразование карты с количеством повторений в список слов (без повторов).
func convertRangeMapToList(rangeList map[string]int) []string {
	list := make([]string, 0, len(rangeList))
	for key := range rangeList {
		list = append(list, key)
	}

	return list
}

// Сортировка списка по частоте и по алфавиту.
func sortByRangeByAbc(list []string, rangeList map[string]int) []string {
	sort.Slice(list, func(i, j int) bool {
		// Сортировка по алфавиту.
		if rangeList[list[i]] == rangeList[list[j]] {
			return list[i] < list[j]
		}

		// Сортировка по частоте.
		return rangeList[list[i]] > rangeList[list[j]]
	})

	return list
}

// Ограничение списка слов.
func limit(list []string, limit int) []string {
	if len(list) > limit {
		return list[:limit]
	}

	return list
}

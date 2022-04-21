package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

// TopSize - Размер топ-листа.
const TopSize = 10

// Top10 - получение списка топ-10.
func Top10(text string) []string {
	// Преобразование текста в карту с количеством повторений.
	rangeList := convertToRangeMap(text)

	// Преобразование карты с количеством повторений в список слов (без повторов).
	list := convertToList(rangeList)

	// Сортировка списка по частоте и по алфавиту.
	sortByRangeByAbc(list, rangeList)

	return limit(list, TopSize)
}

// Преобразование текста в карту с количеством повторений.
func convertToRangeMap(text string) map[string]int {
	// Преобразование текста в оригинальный список слов (с повторами).
	list := strings.Fields(text)

	// Создание карты слов с количеством повторений.
	rangeList := make(map[string]int)
	for _, item := range list {
		rangeList[item]++
	}

	return rangeList
}

// Преобразование карты с количеством повторений в список слов (без повторов).
func convertToList(rangeList map[string]int) []string {
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

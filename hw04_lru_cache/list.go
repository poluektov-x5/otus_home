package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	first  *ListItem
	last   *ListItem
	length int
}

// Len - Получение количества элементов в списке.
func (l list) Len() int {
	return l.length
}

// Front - Получение первого элемента списка.
func (l list) Front() *ListItem {
	return l.first
}

// Back - Получение последнего элемента списка.
func (l list) Back() *ListItem {
	return l.last
}

// PushFront - Добавление элемента в начало списка.
func (l *list) PushFront(v interface{}) *ListItem {
	// Создание нового элемента в начале списка.
	item := &ListItem{Value: v, Next: l.Front()}

	// Установка в соседнем элементе ссылки на новый элемент.
	firstItem := l.Front()
	if firstItem != nil {
		firstItem.Prev = item
	}

	// Установка в списке ссылки на первый элемент.
	l.first = item
	if l.length == 0 {
		l.last = item
	}

	// Инкремент счетчика количества элементов.
	l.length++

	return item
}

// PushBack - Добавление элемента в конец списка.
func (l *list) PushBack(v interface{}) *ListItem {
	// Создание нового элемента в конце списка.
	item := &ListItem{Value: v, Prev: l.Back()}

	// Установка в соседнем элементе ссылки на новый элемент.
	lastItem := l.Back()
	if lastItem != nil {
		lastItem.Next = item
	}

	// Установка в списке ссылки на последний элемент.
	l.last = item
	if l.length == 0 {
		l.first = item
	}

	// Инкремент счетчика количества элементов.
	l.length++

	return item
}

// Remove - Удаление элемента из списка.
func (l *list) Remove(item *ListItem) {
	// Получение предыдущего и следующего элементов.
	prevItem := item.Prev
	nextItem := item.Next

	// Если элемент не первый и не последний.
	if prevItem != nil && nextItem != nil {
		prevItem.Next = nextItem
		nextItem.Prev = prevItem
	} else if prevItem != nil { // Если элемент последний.
		prevItem.Next = nil
		l.last = prevItem
	} else if nextItem != nil { // Если элемент первый.
		nextItem.Prev = nil
		l.first = nextItem
	} else { // Если единственный элемент.
		l.last = nil
		l.first = nil
	}

	// Дикремент счетчика количества элементов.
	l.length--
}

// MoveToFront - Перемещение элемента в начало списка.
func (l *list) MoveToFront(item *ListItem) {
	l.Remove(item)
	l.PushFront(item.Value)
}

func NewList() List {
	return new(list)
}

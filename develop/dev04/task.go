package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

/*
	Написать функцию поиска всех множеств анаграмм по словарю.


	Например:
	'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
	'листок', 'слиток' и 'столик' - другому.


	Требования:
	Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
	Выходные данные: ссылка на мапу множеств анаграмм
	Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
	слово из множества.
	Массив должен быть отсортирован по возрастанию.
	Множества из одного элемента не должны попасть в результат.
	Все слова должны быть приведены к нижнему регистру.
	В результате каждое слово должно встречаться только один раз.

*/

var set map[string]struct{}
var indexes map[string]string
var ans map[string][]string

func prepareData(s []string) []string {
	res := make([]string, 0, len(s))

	for i := 0; i < len(s); i++ {
		res = append(res, strings.ToLower(s[i]))
	}

	return res
}

func sortBySlice(s string) []rune {
	sr := []rune(s)
	sort.Slice(sr, func(i int, j int) bool {
		return sr[i] < sr[j]
	})
	return sr
}

func anagram(s []string) {
	set = make(map[string]struct{}, len(s))
	indexes = make(map[string]string, len(s))
	ans = make(map[string][]string, len(s))

	s = prepareData(s)

	var lexStr string = ""
	for i := 0; i < len(s); i++ {
		lexStr = string(sortBySlice(s[i]))

		if _, ok := indexes[lexStr]; !ok {
			indexes[lexStr] = s[i]
		}

		if _, ok := set[s[i]]; !ok {
			set[s[i]] = struct{}{}
			indx := indexes[lexStr]
			ans[indx] = append(ans[indx], s[i])

		}
	}

	for key, sls := range ans {
		if len(sls) == 1 {
			delete(ans, key)
			continue
		}

		slices.SortFunc(sls, func(a string, b string) int {
			return strings.Compare(a, b)
		})

	}
}

func main() {
	data := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "СтоЛИК", "ПЯТКА", "листок"}
	anagram(data)
	fmt.Println(ans)
}

package usecase

type StringFinderUsecase interface {
	FindMaxSubstring(s string) string
}

type StringFinder struct{}

func NewStringFinder() *StringFinder {
	return &StringFinder{}
}

func (u *StringFinder) FindMaxSubstring(s string) string {
	// var result string
	// for i := 0; i < len(s); i++ {
	// 	sub := ""
	// 	for j := i; j < len(s); j++ {
	// 		if index := strings.IndexByte(sub, s[j]); index == -1 {
	// 			sub += string(s[j])
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	if utf8.RuneCountInString(sub) > utf8.RuneCountInString(result) {
	// 		result = sub
	// 	}
	// }
	// return result
	var maxSubstring string
	var currentSubstring string

	// Используем хеш-таблицу для отслеживания повторяющихся символов
	charMap := make(map[rune]bool)

	// Проходим по всей строке, одновременно поддерживая текущую подстроку без повторяющихся символов
	for _, c := range s {
		if charMap[c] {
			// Если символ уже встречался ранее, обрезаем текущую подстроку до следующего индекса после последнего вхождения символа
			for i, char := range currentSubstring {
				if char == c {
					currentSubstring = currentSubstring[i+1:]
					break
				}
				delete(charMap, char)
			}
		}

		// Добавляем символ в текущую подстроку и хеш-таблицу
		currentSubstring += string(c)
		charMap[c] = true

		// Если текущая подстрока больше максимальной, обновляем максимальную
		if len(currentSubstring) > len(maxSubstring) {
			maxSubstring = currentSubstring
		}
	}

	return maxSubstring
}

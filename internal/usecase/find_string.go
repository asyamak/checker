package usecase

type StringFinderUsecase interface {
	FindMaxSubstring(s string) string
}

type StringFinder struct{}

func NewStringFinder() *StringFinder {
	return &StringFinder{}
}

func (u *StringFinder) FindMaxSubstring(s string) string {
	var maxSubstring string
	var currentSubstring string

	charMap := make(map[rune]bool)

	for _, c := range s {
		if charMap[c] {
			for i, char := range currentSubstring {
				if char == c {
					currentSubstring = currentSubstring[i+1:]
					break
				}
				delete(charMap, char)
			}
		}

		currentSubstring += string(c)
		charMap[c] = true

		if len(currentSubstring) > len(maxSubstring) {
			maxSubstring = currentSubstring
		}
	}

	return maxSubstring
}

package asciiHandler

import (
	"os"
	"strings"
)

func WordSplit(s string) []string {
	str := ""
	slice := []string{}
	for i := 0; i < len(s); i++ {
		// 	if i+1 < len(s) && s[i] == '\\' && s[i+1] == 'n' {
		// 		if len(str) != 0 {
		// 			slice = append(slice, str)
		// 			str = ""
		// 		}
		// 		i = i + 1
		// 		slice = append(slice, "\n")
		if s[i] == 10 {
			if len(str) != 0 {
				slice = append(slice, str)
				str = ""
			}
			slice = append(slice, "\n")
		} else {
			str = str + string(s[i])
		}
	}

	if len(str) != 0 {
		slice = append(slice, str)
		str = ""
	}

	return slice
}

func Mapa(fileBanner string, inputText string) (string, error) {
	bytes, err := os.ReadFile(fileBanner)
	if err != nil {
		return "", err
	}

	strarr := strings.Split(string(bytes), "\n")

	mapa := make(map[rune]int)

	ind := 1
	result := ""
	for i := ' '; i <= '~'; i++ {
		mapa[i] = ind
		ind += 9
	}
	words := WordSplit(inputText)
	for _, word := range words {
		if word == "\n" {
			result += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for _, v := range word {
					result += strarr[mapa[v]+i]
				}
				if i != 7 {
					result += "\n"
				}

			}
		}
	}
	if words[len(words)-1] != "\n" {
		result += "\n"
	}
	return result, nil
}

func IsPrintable(str string) bool {
	for i := 0; i < len(str); i++ {
		if (str[i] < 32 || str[i] > 126) && str[i] != 10 && str[i] != 13 {
			return false
		}
	}
	return true
}

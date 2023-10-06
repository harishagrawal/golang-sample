package reloaded

import (
	"regexp"
)

func back_line(arr []string, i int) int {
	j := 0
	reg := regexp.MustCompile(`[?:;,.!\n()'{}\[\]"]`)
	if i > 0 {
		j = i - 1
	}
	for ; j > 0; j-- {
		if !reg.MatchString(arr[j]) {
			break
		}
	}
	return j
}

func ArticleFix(arr []string) []string {
	reg := regexp.MustCompile(`[ayeoihAYEHIO]`)
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != "\n" {
			if reg.MatchString(arr[i][0:1]) {
				j = back_line(arr, i)
				if arr[j] == "a" || arr[j] == "A" {
					arr[j] = arr[j] + "n"
				}
			} else {
				j = back_line(arr, i)
				if arr[j] == "an" || arr[j] == "AN" || arr[j] == "An" || arr[j] == "aN" {
					arr[j] = arr[j][0:1]
				}

			}
		}
		// if arr[i] == "a" || arr[i] == "A" {
		// 	for j := i + 1; ; {
		// 		fmt.Println(arr[i], "it is artiocle fix")
		// 		if arr[j] == "'" || arr[j] == "\"" || arr[j] == "\n" {
		// 			j++
		// 		} else {
		// 			for _, v := range vowels {
		// 				if arr[j][0] == byte(v) {
		// 					arr[i] = arr[i] + "n"
		// 					fmt.Println("++++++")
		// 				}
		// 				break
		// 			}
		// 			break
		// 		}
		// 	}
		// }eak{
		// 			if arr[j] == "'" || arr[j] == "\"" || arr[j] == "\n" {
		// 				j++
		// 			} else {
		// 				for _, v := range vowels {
		// 					if arr[j][0] != byte(v) {
		// 						arr[i] = arr[i][:1]
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}
		// }
	}
	return arr
}

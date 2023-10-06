package reloaded

func CommandFix(input []string) []string {
	start_ind := 0
	end_ind := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "(" && (input[i+1] == "up" || input[i+1] == "hex" || input[i+1] == "bin" || input[i+1] == "low" || input[i+1] == "cap") {
			start_ind = i
			for j := i + 1; j < len(input); {
				if input[j] == "\n" {
					j++
					continue
				}
				if input[j] != ")" {
					input[i] += input[j]
					j++
				} else {
					input[i] += input[j]
					end_ind = j
					break
				}
			}
			input = append(input[:start_ind+1], input[end_ind+1:]...)

		}
	}


	// (up) (up, jddc)
	// (up,

	// up := regexp.MustCompile(`\(([^)]*up[^)]*)\)`)

	// res := up.ReplaceAllFunc(input, func(match []byte) []byte {
	// 	if bytes.Contains(match, []byte("up")) {
	// 		return bytes.ReplaceAll(match, []byte(" "), []byte(""))
	// 	}
	// 	return match
	// })

	// hex := regexp.MustCompile(`\(([^)]*hex[^)]*)\)`)

	// res = hex.ReplaceAllFunc(res, func(match []byte) []byte {
	// 	if bytes.Contains(match, []byte("hex")) {
	// 		return bytes.ReplaceAll(match, []byte(" ")2
	// })

	// bin := regexp.MustCompile(`\(([^)]*bin[^)]*)\)`)

	// res = bin.ReplaceAllFunc(res, func(match []byte) []byte {
	// 	if bytes.Contains(match, []byte("bin")) {
	// 		return bytes.ReplaceAll(match, []byte(" "), []byte(""))
	// 	}
	// 	return match
	// })

	// low := regexp.MustCompile(`\(([^)]*low[^)]*)\)`)

	// res = low.ReplaceAllFunc(res, func(match []byte) []byte {
	// 	if bytes.Contains(match, []byte("low")) {
	// 		return bytes.ReplaceAll(match, []byte(" "), []byte(""))
	// 	}
	// 	return match
	// })

	// cap := regexp.MustCompile(`\(([^)]*cap[^)]*)\)`)

	// res = cap.ReplaceAllFunc(res, func(match []byte) []byte {
	// 	if bytes.Contains(match, []byte("cap")) {
	// 		return bytes.ReplaceAll(match, []byte(" "), []byte(""))
	// 	}
	// 	return match
	// })
	return input
}

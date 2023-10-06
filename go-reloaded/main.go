package main

import (
	"fmt"
	"os"
	reloaded "reloaded/functions"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Error: invalid number of arguments")
		return
	}

	if args[0] == args[1] {
		fmt.Println("Error: invalid filenames. Please enter different arguments")
		return
	}

	if !strings.Contains(args[0], ".txt") || !strings.Contains(args[1], ".txt") {
		fmt.Println("Error: invalid datatype of arguments. Please use .txt notation")
		return

	}

	_, err := os.Stat(args[0])
	if err != nil {
		os.Create(args[0])
	}

	bytes, err := os.ReadFile(args[0])
	if err != nil {
		err.Error()
	}

	strarr := reloaded.BytesToStrArr(bytes)

	strarr = reloaded.CommandFix(strarr)

	strarr = reloaded.ArticleFix(strarr)

	command := ""
	index := 0
	res := 0

	for ind, str := range strarr {
		if reloaded.IsCommand(str) {
			if ind == 0 {
				fmt.Println("Command can be apllied only to word before it. Please change the command position")
				return
			}
			command, index = reloaded.CommandInfo(str)

			if command == "bin" && index != 1 {
				fmt.Println("Error: bin command can not handle any parameters")
				return
			}
			if command == "hex" && index != 1 {
				fmt.Println("Error: hex command can not handle any parameters")
				return
			}

			if index == 0 {
				fmt.Println("Error: used invalid command parameter. Please check if parameter exists or it is positive number")
				fmt.Println("Check COMMAND: == ", command, " ==")
				return
			} else if index > ind {
				fmt.Println("Error: used invalid command parameter. Please check if parameter is bigger than words quantity before command")
				fmt.Println("Check COMMAND: == ", command, " == ;  PARAM: ==", index, "== ")
				return
			} else {
				for i := ind; i > 0 && index > 0; i-- {
					if !reloaded.IsPunc(strarr[i-1]) {
						if command == "up" {
							strarr[i-1] = strings.ToUpper(strarr[i-1])
							// strarr = append(strarr[:i], strarr[i+1:]...)
							// i--
							index--
						}

						if command == "low" {
							strarr[i-1] = strings.ToLower(strarr[i-1])
							// strarr = append(strarr[:i], strarr[i+1:]...)
							// i--
							index--
						}

						if command == "cap" {
							strarr[i-1] = strings.ToLower(strarr[i-1])
							strarr[i-1] = strings.Title(strarr[i-1])
							// strarr = append(strarr[:i], strarr[i+1:]...)
							// i--
							index--
						}

						// fmt.Println(reloaded.AtoiBase("10", "01"))

						// if command == "hex" {
						// 	strarr[i-1] = strconv.Itoa(reloaded.AtoiBase(strarr[i-1], "0123456789ABCDEF"))
						// 	// strarr = append(strarr[:i], strarr[i+1:]...)
						// 	// i--
						// 	index--
						// }

						if command == "bin" {
							res = reloaded.AtoiBase(strarr[i-1], "01")
							if res != 0 {
								strarr[i-1] = strconv.Itoa(res)
								index--
							} else {
								fmt.Println("Error: used invalid command argument\nCommand: bin")
								return
							}

						}

						if command == "hex" {
							res = reloaded.AtoiBase(strarr[i-1], "0123456789ABCDEF")
							if res != 0 {
								strarr[i-1] = strconv.Itoa(res)
								index--
							} else {
								fmt.Println("Error: used invalid command argument\nCommand: hex")
								return
							}

						}

					}
				}
			}

		}
		// fmt.Println("startstartsatrtsatrt")

		// for i, w := range strarr {
		// 	fmt.Printf("%v.%v\n", i, w)
		// }

		// fmt.Println()

		// fmt.Println(strarr)
	}

	if reloaded.QuotesPair(strarr) {
		strarr = reloaded.QuotesFix(strarr)
	} else {
		fmt.Println("Can not be done, qoutes are not paired")
		return
	}

	strarr = reloaded.BracketsFix(strarr)

	strarr = reloaded.Punc(strarr)

	// for i, w := range strarr {
	// 	fmt.Printf("%v.%v\n", i, w)
	// }

	// fmt.Println()

	// fmt.Println(strarr)

	towrite := ""

	for _, v := range strarr {
		if !reloaded.IsCommand(v) {
			towrite += v
			if v != "\n" {
				towrite += " "
			}
		}
	}

	os.WriteFile(args[1], []byte(towrite), 0644)
}

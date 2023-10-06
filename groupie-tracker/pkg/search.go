package pkg

import (
	"strconv"
	"strings"
)

func SearchOut(list []Artists, targetStr string) []Artists {
	var newList []Artists
	target := strings.ToLower(targetStr)
	target_arr := strings.Split(target, " -> ")
	flag := false
	for _, val := range list {
		flag = false
		// search by name
		if strings.Contains(strings.ToLower(val.Name), target_arr[0]) {
			flag = true
		}
		// search by members
		for _, memb := range val.Members {
			if strings.Contains(strings.ToLower(memb), target_arr[0]) {
				flag = true
			}
		}
		// search by creation date
		if strings.Contains(strings.ToLower(strconv.Itoa(val.CreationDate)), target_arr[0]) {
			flag = true
		}
		// first albume date
		if strings.Contains(strings.ToLower(val.FirstAlbum), target_arr[0]) {
			flag = true
		}
		// dates and locations search
		for key, element := range val.DatesLocations {
			if strings.Contains(strings.ToLower(key), target_arr[0]) {
				flag = true
			}
			for _, vals := range element {
				if strings.Contains(strings.ToLower(vals), target_arr[0]) {
					flag = true
				}
			}
		}
		if flag {
			newList = append(newList, val)
		}
	}
	return newList
}

func Suggestions(dataset []Artists) map[string]string {
	opt := make(map[string]string)

	for _, val := range dataset {

		opt[val.Name] = " -> name"
		// search by members

		for _, mems := range val.Members {
			opt[mems] = " -> member"
		}
		opt[strconv.Itoa(val.CreationDate)] = " -> creation date"

		opt[val.FirstAlbum] = " -> first album"
		// dates and locations search

		for key, element := range val.DatesLocations {
			opt[key] = " -> concert location"

			for _, vals := range element {
				opt[vals] = " -> concert date"
			}
		}
	}

	return opt
}

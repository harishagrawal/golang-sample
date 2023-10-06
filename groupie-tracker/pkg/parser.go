package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func Unmarshal() ([]Artists, error) {
	art_prs, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer art_prs.Body.Close()
	body, err := io.ReadAll(art_prs.Body)
	if err != nil {
		return nil, err
	}
	var Art []Artists
	err = json.Unmarshal(body, &Art)
	if err != nil {
		return nil, err
	}
	rel_prs, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer rel_prs.Body.Close()
	body_rel, err := io.ReadAll(rel_prs.Body)
	if err != nil {
		return nil, err
	}
	var Rel Relation
	err = json.Unmarshal(body_rel, &Rel)
	if err != nil {
		return nil, err
	}
	for i := range Rel.Index {
		newMap := make(map[string][]string)
		for key, val := range Rel.Index[i].DatesLocations {
			newKey := LocMidify(key)
			newMap[newKey] = val
		}
		Rel.Index[i].DatesLocations = newMap
	}
	for i := range Art {
		Art[i].DatesLocations = Rel.Index[i].DatesLocations
	}
	return Art, nil
}

func LocMidify(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", ", ")
	s = strings.ToTitle(s)
	return s
}

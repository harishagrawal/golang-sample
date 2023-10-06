package pkg

type Artists struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Search struct {
	Artists  []Artists
	Datalist map[string]string
}

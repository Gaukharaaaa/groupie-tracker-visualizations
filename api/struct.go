package api

type Artists struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

type Artist struct {
	Id             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	LocationsUrl   string              `json:"locations"`
	ConcertDateUrl string              `json:"concertDates"`
	Details        map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Id             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

const (
	AllLinks  = "https://groupietrackers.herokuapp.com/api/artists"
	Relations = "https://groupietrackers.herokuapp.com/api/relation/"
)

package models

type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationdate"`
	FirstAlbum     string   `json:"firstalbum"`
	DatesLocations map[string][]string
}
type Relations struct {
	Location []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"Index"`
}
type AllArtist struct {
	Artist      []Artist
	FoundArtist []Artist
}

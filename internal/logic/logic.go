package logic

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie/models"
)

var (
	linkArtist   = "https://groupietrackers.herokuapp.com/api/artists"
	linkRelation = "https://groupietrackers.herokuapp.com/api/relation"
	artist       models.AllArtist
	relation     models.Relations
	file         = "json/artist.json"
	fileRel      = "json/relation.json"
)

// function GetData receives data from API link and saves it into structs.
func GetData() (models.AllArtist, error) {
	err := GetJson(linkArtist, &artist.Artist, file)
	if err != nil {
		log.Printf("Error getting json in GetArtist function: %s\n", err.Error())
		return artist, err
	}
	err = GetJson(linkRelation, &relation, fileRel)
	if err != nil {
		log.Printf("Error getting json in GetArtist function: %s\n", err.Error())
		return artist, err
	}
	for i, w := range relation.Location {
		artist.Artist[i].DatesLocations = w.DatesLocations
	}
	return artist, nil
}

var errReceveUrl = errors.New("error receving data through API link")

// function GetJson receives an API url and unparsing it into given struct. Also saving uparsed
// JSON data to a file.
func GetJson(url string, target interface{}, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error with receiving data: %v\n", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error with receiving data through url : %d\n", resp.StatusCode)
		return errReceveUrl
	}
	if err = json.NewDecoder(resp.Body).Decode(&target); err != nil {
		return err
	}
	file, _ := json.MarshalIndent(&target, "", "\t")
	if err = ioutil.WriteFile(filename, file, 0o644); err != nil {
		log.Printf("Error saving json file: %v\n", err)
		return err
	}
	return nil
}

// function SearchBar receive query , which should process and compare it with Artist's data
// if any Artist consist the same information with query -  it will returns Artist's data.
func SearchBar(search string) (*models.AllArtist, error) {
	var all []models.Artist
	entries, err := GetData()
	if err != nil {
		log.Printf("Error with receiving data: %d\n", err)
		return nil, err
	}
	for _, query := range entries.Artist {
		var found bool
		if Compare(query.Name, search) {
			all = append(all, query)
			continue
		}
		if Compare(query.FirstAlbum, search) {
			all = append(all, query)
			continue
		}
		if Compare(strconv.Itoa(query.CreationDate), search) {
			all = append(all, query)
			continue
		}
		for _, m := range query.Members {
			if Compare(m, search) {
				all = append(all, query)
				found = true
				break // LOOP
			}
		}
		if !found {
			for key := range query.DatesLocations {
				if Compare(key, search) {
					all = append(all, query)
					break
				}
			}
		}
	}
	Result := &models.AllArtist{
		Artist: all,
	}
	return Result, nil
}

// function Compare lowers both receiving strings and compares them. Returns true if string query consists
// itself same pattern from compare string.
func Compare(compare, query string) bool {
	return strings.Contains(
		strings.ToLower(compare),
		strings.ToLower(query),
	)
}

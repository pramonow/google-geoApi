package geomap

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type GoogleGeocodeResponse struct {
	Results []struct {
		AddressComponents []AddressComponent `json:"address_components"`
		FormattedAddress  string             `json:"formatted_address"`
		Geometry          GoogleGeometry     `json:"geometry"`
		PlaceID           string             `json:"place_id"`
		PlusCode          GooglePlusCode     `json:"plus_code"`
		Types             []string           `json:"types"`
	} `json:"results"`
	Status string `json:"status"`
}

type GooglePlaceSearchResponse struct {
	Candidates []Candidate `json:"candidates"`
	Status     string      `json:"status"`
}

type GoogleNearbySearchResponse struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Results          []struct {
		Geometry         GoogleGeometry `json:"geometry"`
		Icon             string         `json:"icon"`
		ID               string         `json:"id"`
		Name             string         `json:"name"`
		OpeningHours     OpeningHour    `json:"opening_hours"`
		Photos           []Photo        `json:"photos"`
		PlaceID          string         `json:"place_id"`
		PlusCode         GooglePlusCode `json:"plus_code"`
		PriceLevel       int            `json:"price_level,omitempty"`
		Rating           float64        `json:"rating"`
		Reference        string         `json:"reference"`
		Scope            string         `json:"scope"`
		Types            []string       `json:"types"`
		UserRatingsTotal int            `json:"user_ratings_total"`
		Vicinity         string         `json:"vicinity"`
	} `json:"results"`
	Status string `json:"status"`
}

type OpeningHour struct {
	OpenNow bool `json:"open_now"`
}

type Candidate struct {
	FormattedAddress string  `json:"formatted_address"`
	Name             string  `json:"name"`
	Photos           []Photo `json:"photos"`
	Rating           int     `json:"rating"`
}

type Photo struct {
	Height           int      `json:"height"`
	HTMLAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}

type GooglePlusCode struct {
	CompoundCode string `json:"compound_code"`
	GlobalCode   string `json:"global_code"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type GoogleGeometry struct {
	Location     GoogleLocation `json:"location"`
	LocationType string         `json:"location_type,omitempty"`
	Viewport     GoogleViewport `json:"viewport"`
}

type GoogleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type GoogleViewport struct {
	Northeast GoogleLocation `json:"northeast"`
	SouthWest GoogleLocation `json:"southwest"`
}

var (
	client *http.Client
)

func init() {

	client = &http.Client{}
}

/*
	GetReverseGeoCode will return GoogleReverseGeocodeResponse on success
	the example of usage is sending params that contains "address" and "key" (both of them are required)
	Key is obtained in config.GoogleMap.Key
	more references https://developers.google.com/maps/documentation/geocoding/intro#Geocoding
*/
func GetGeocode(ctx context.Context, params map[string]string) (GoogleGeocodeResponse, error) {

	var googleGeocodeResponse GoogleGeocodeResponse

	//Generating url for geocode
	reqURL := "https://maps.googleapis.com/maps/api/geocode/json"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return googleGeocodeResponse, err
	}

	//Insert the query mapping into the request
	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return googleGeocodeResponse, err
	}

	if resp.StatusCode != http.StatusOK {
		return googleGeocodeResponse, errors.New("Status not OK")
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return googleGeocodeResponse, err
	}

	//Unmarshal the contents
	err = json.Unmarshal(contents, &googleGeocodeResponse)
	if err != nil {
		return googleGeocodeResponse, err
	}

	return googleGeocodeResponse, nil
}

func FindPlace(ctx context.Context, params map[string]string) (GooglePlaceSearchResponse, error) {

	var googleFindPlaceResponse GooglePlaceSearchResponse

	//Generating url for geocode
	reqURL := "https://maps.googleapis.com/maps/api/place/findplacefromtext/json"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return googleFindPlaceResponse, err
	}

	//Insert the query mapping into the request
	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return googleFindPlaceResponse, err
	}

	if resp.StatusCode != http.StatusOK {
		return googleFindPlaceResponse, errors.New("Status not OK")
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return googleFindPlaceResponse, err
	}

	//Unmarshal the contents
	err = json.Unmarshal(contents, &googleFindPlaceResponse)
	if err != nil {
		return googleFindPlaceResponse, err
	}

	return googleFindPlaceResponse, nil
}

/*
	By default, when a user selects a place, Nearby Search returns all of the available data fields for the selected place,
	and you will be billed accordingly. There is no way to constrain Nearby Search requests to only return specific fields.
	To keep from requesting (and paying for) data that you don't need, use a Find Place request instead.
*/
func PlaceNearby(ctx context.Context, params map[string]string) (GoogleNearbySearchResponse, error) {

	var googleNearbySearchResponse GoogleNearbySearchResponse

	//Generating url for geocode
	reqURL := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return googleNearbySearchResponse, err
	}

	//Insert the query mapping into the request
	q := req.URL.Query()
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return googleNearbySearchResponse, err
	}

	if resp.StatusCode != http.StatusOK {
		return googleNearbySearchResponse, errors.New("Status not OK")
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return googleNearbySearchResponse, err
	}

	//Unmarshal the contents
	err = json.Unmarshal(contents, &googleNearbySearchResponse)
	if err != nil {
		return googleNearbySearchResponse, err
	}

	return googleNearbySearchResponse, nil
}

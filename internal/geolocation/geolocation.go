package geolocation

import "fmt"

type GeolocationResult struct {
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

func (g *GeolocationResult) DisplayDescribe() string {

	return fmt.Sprintf("%s (%s, %s)", g.Name, g.Country, g.State)

}

package geolocation

import "fmt"

type GeolocationResult struct {
	Name    string  `json:"name"`
	Lat     float32 `json:"lat"`
	Lng     float32 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}

func (g *GeolocationResult) Describe() string {

	return fmt.Sprintf("%s (%s, %s)", g.Name, g.Country, g.State)

}

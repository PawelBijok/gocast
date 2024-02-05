package weather

import "github.com/pafello/gocast/internal/units"

type Forecast struct {
	List           []Weather `json:"list"`
	UnitSystemUsed units.UnitSystem
}

package geometry

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Coordinate [2]float64

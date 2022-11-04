package dao

import "github.com/uptrace/bun"

type UnomCoordinate struct {
	bun.BaseModel `bun:"table:unom_coordinates,alias:uc"`

	UNOM      int     `bun:"unom,pk"`
	Latitude  float64 `bun:"latitude"`
	Longitude float64 `bun:"longitude"`
}

package ebitenLDTK

import (
	"fmt"
)

type Defs struct {
	Tilesets      []Tileset      `json:"tilesets"`
	Enums         []Enum         `json:"enums"`
	ExternalEnums []ExternalEnum `json:"externalEnums"`
	LevelFields   []LevelField   `json:"levelFields"`
	LayerDefs     []LayerDef     `json:"layers"`
}

type Enum struct {
	// TBA
}

type ExternalEnum struct {
	// TBA
}

type LevelField struct {
	// TBA
}

type LayerDef struct {
	Name          string         `json:"identifier"`
	IntGridValues []IntGridValue `json:"intGridValues"`
}

type IntGridValue struct {
	Value      int    `json:"value"`
	Identifier string `json:"identifier"`
}

func (ld *LayerDef) GetIntGridValue(identifier string) int {
	for _, intGridValue := range ld.IntGridValues {
		if intGridValue.Identifier == identifier {
			return intGridValue.Value
		}
	}
	return -1
}

func (d *Defs) GetTilesetByUid(uid int) (Tileset, error) {
	for _, tileset := range d.Tilesets {
		if tileset.Uid == uid {
			return tileset, nil
		}
	}
	return Tileset{}, fmt.Errorf("tileset with uid [%d] was not found", uid)
}

package ebitenLDTK

import (
	"fmt"
	"slices"
)

type Defs struct {
	Tilesets      []Tileset    `json:"tilesets"`
	Enums         []EnumDef    `json:"enums"`
	ExternalEnums []EnumDef    `json:"externalEnums"`
	LevelFields   []LevelField `json:"levelFields"`
	LayerDefs     []LayerDef   `json:"layers"`
}

type EnumDef struct {
	Name   string      `json:"identifier"`
	Values []EnumValue `json:"values"`
}

type EnumValue struct {
	Id string `json:"id"`
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

func (d *Defs) GetEnum(name string) (EnumDef, error) {
	for _, enum := range slices.Concat(d.Enums, d.ExternalEnums) {
		if enum.Name == name {
			return enum, nil
		}
	}
	return EnumDef{}, fmt.Errorf("enum with name [%d] was not found", name)
}

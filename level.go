package ebitenLDTK

import (
	"fmt"
)

type Level struct {
	Name       string  `json:"identifier"`
	Iid        string  `json:"iid"`
	Uid        int     `json:"uid"`
	WorldX     int     `json:"worldX"`
	WorldY     int     `json:"worldY"`
	WorldDepth int     `json:"worldDepth"`
	PxWid      float64 `json:"pxWid"`
	PxHei      float64 `json:"pxHei"`
	BgColorHex string  `json:"__bgColor"`
	Layers     []Layer `json:"layerInstances"`
	Fields     []Field `json:"fieldInstances"`
}

func (l *Level) GetLayerByName(name string) (Layer, error) {
	for _, layer := range l.Layers {
		if layer.Name == name {
			return layer, nil
		}
	}
	return Layer{}, fmt.Errorf("layer with name [%s] was not found", name)
}

func (l *Level) GetEntityByIid(iid string) (Entity, error) {
	for _, layer := range l.Layers {
		for _, entity := range layer.Entities {
			if entity.Iid == iid {
				return entity, nil
			}
		}

	}
	return Entity{}, fmt.Errorf("entity with iid [%s] was not found", iid)
}

func (l *Level) GetFieldByName(name string) (Field, error) {
	for _, field := range l.Fields {
		if field.Name == name {
			return field, nil
		}
	}
	return Field{}, fmt.Errorf("field with name [%s] was not found", name)
}

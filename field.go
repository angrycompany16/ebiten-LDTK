package ebitenLDTK

// TODO: Write the integration test
// Right now I'm too lazy though

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FieldType string

const (
	FieldTypeEntityRef      = "EntityRef"
	FieldTypeEntityRefArray = "Array<EntityRef>"
	FieldTypeFloat          = "Float"
	FieldTypeString         = "String"
	FieldTypeBiome          = "LocalEnum.Biome"
	FieldTypePoint          = "Point"
)

type Field struct {
	Name           string
	Type           FieldType
	EntityRef      EntityRefValue
	EntityRefArray []EntityRefValue
	Float          float64
	String         string
	Biome          string
	Point          PointValue
}

type EntityRefValue struct {
	EntityIid string `json:"entityIid"`
	LevelIid  string `json:"levelIid"`
}

type PointValue struct {
	X, Y float64
}

// They call me the programming wizard
func (f *Field) UnmarshalJSON(data []byte) error {
	var result map[string]any
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	name, ok := result["__identifier"].(string)
	if !ok {
		return errors.New("could not cast result[\"__identifier\"] to string")
	}
	f.Name = name

	fieldType, ok := result["__type"].(string)
	if !ok {
		return errors.New("could not cast result[\"__type \"] to string")
	}
	f.Type = FieldType(fieldType)

	switch f.Type {
	case FieldTypeFloat:
		float, ok := result["__value"].(float64)
		if !ok {
			return errors.New("could not cast [\"__value\"] to float64")
		}
		f.Float = float
	case FieldTypeString:
		string, ok := result["__value"].(string)
		if !ok {
			return errors.New("could not cast [\"__value\"] to string")
		}
		f.String = string
	case FieldTypeEntityRef:
		entityRefValue := EntityRefValue{}
		entityRef, ok := result["__value"].(map[string]any)
		if !ok {
			return errors.New("could not cast result[\"__value\"] to map[string]any")
		}
		JSONdata, err := json.Marshal(entityRef)
		if err != nil {
			return errors.New("could not Marshal entityRef")
		}

		err = json.Unmarshal(JSONdata, &entityRefValue)
		if err != nil {
			return errors.New("could not Unmarshal entityRefValue")
		}
		f.EntityRef = entityRefValue
	case FieldTypeEntityRefArray:
		entityRefArray, ok := result["__value"].([]any)
		if !ok {
			return errors.New("could not cast result[\"__value\"] to []any")
		}

		for _, _entityRef := range entityRefArray {
			var entityRefValue EntityRefValue
			entityRef, ok := _entityRef.(map[string]any)

			if !ok {
				return errors.New("could not cast entityRef to map[string]any")
			}

			JSONdata, err := json.Marshal(entityRef)
			if err != nil {
				return errors.New("could not Marshal entityRefArray")
			}

			err = json.Unmarshal(JSONdata, &entityRefValue)
			if err != nil {
				return errors.New("could not Unmarshal entityRefValue")
			}
			f.EntityRefArray = append(f.EntityRefArray, entityRefValue)
		}
	case FieldTypeBiome:
		if result["__value"] == nil {
			f.Biome = ""
			fmt.Println("Assigning empty biome to level with missing biome field")
		} else {
			biome := result["__value"].(string)
			f.Biome = biome
		}
	case FieldTypePoint:
		point := result["__value"]
		if point == nil {
			f.Point = PointValue{
				X: 0,
				Y: 0,
			}
			return nil
		}
		pointMap, ok := point.(map[string]any)
		if !ok {
			return errors.New("could not cast [\"__value\"] to map[string]any")
		}
		f.Point = PointValue{
			X: pointMap["cx"].(float64),
			Y: pointMap["cy"].(float64),
		}
	}
	return nil
}

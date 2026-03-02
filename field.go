package ebitenLDTK

// TODO: Write the integration test
// Right now I'm too lazy though

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// It's possible that this is a bit unnecessary
type FieldTypeV2 int

const (
	FieldTypeNone FieldTypeV2 = iota - 1
	FieldTypeInt
	FieldTypeFloat
	FieldTypeBool
	FieldTypeString
	FieldTypeMultilines
	FieldTypeColor
	FieldTypeEnum
	FieldTypeFilePath
	FieldTypeTile
	FieldTypeEntityRef
	FieldTypePoint
)

func stringToFieldType(fieldType string) (FieldTypeV2, bool) {
	isArray := false
	if strings.HasPrefix(fieldType, "Array") {
		isArray = true
		fieldType = strings.TrimPrefix(fieldType, "Array<")
		fieldType = strings.TrimSuffix(fieldType, ">")
	}

	if fieldType == "Int" {
		return FieldTypeInt, isArray
	} else if fieldType == "Float" {
		return FieldTypeFloat, isArray
	} else if fieldType == "Bool" {
		return FieldTypeBool, isArray
	} else if fieldType == "String" {
		return FieldTypeString, isArray
	} else if fieldType == "Multilines" {
		return FieldTypeMultilines, isArray
	} else if fieldType == "Color" {
		return FieldTypeColor, isArray
	} else if strings.HasPrefix(fieldType, "LocalEnum") || strings.HasPrefix(fieldType, "ExternEnum") {
		return FieldTypeEnum, isArray
	} else if fieldType == "FilePath" {
		return FieldTypeFilePath, isArray
	} else if fieldType == "Tile" {
		return FieldTypeTile, isArray
	} else if fieldType == "EntityRef" {
		return FieldTypeEntityRef, isArray
	} else if fieldType == "Point" {
		return FieldTypePoint, isArray
	}
	return FieldTypeNone, isArray
}

type Field struct {
	isArray bool
	value   any
	Name    string
}

type Color struct {
	R uint8
	G uint8
	B uint8
}

type EntityRef struct {
	EntityIid string
	LevelIid  string
}

type Point struct {
	X, Y float64
}

type Enum struct {
	Name  string
	Value string
}

type TileFieldValue struct {
	TileSetUID int
	X          float64
	Y          float64
	W          float64
	H          float64
}

// Gets the field value for fields of type Single
func As[T any](f Field) T {
	var zero T
	if f.value == nil {
		return zero
	}
	if f.isArray {
		fmt.Println("The field represents an array. Please is AsArray instead.")
		return zero
	}

	if val, ok := f.value.(T); ok {
		return val
	}

	fmt.Printf("Error converting value [%v] to [%T]. Try converting to [%T] instead.\n", f.value, zero, f.value)
	return zero
}

// Gets the value of fields of type array
func AsArray[T any](f Field) []T {
	if !f.isArray {
		fmt.Println("The field does not represent an array. Please use As instead.")
		return make([]T, 0)
	}
	valueSlice := f.value.([]any)
	formattedSlice := make([]T, len(valueSlice))
	var zero T
	for i, value := range valueSlice {
		if value == nil {
			formattedSlice[i] = zero
			continue
		}

		if val, ok := value.(T); ok {
			formattedSlice[i] = val
			continue
		}

		fmt.Printf("Error converting value [%v] to [%T]. Try converting to [%T] instead.\n", value, zero, value)
		formattedSlice[i] = zero
	}
	return formattedSlice
}

func (f *Field) UnmarshalJSON(data []byte) error {
	var result map[string]any
	err := json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	name, ok := result["__identifier"].(string)
	if !ok {
		return errors.New("failed to cast result[\"__identifier\"] to string")
	}
	f.Name = name

	fieldTypeStr, ok := result["__type"].(string)
	if !ok {
		return errors.New("failed to cast result[\"__type \"] to string")
	}
	fieldType, isArray := stringToFieldType(fieldTypeStr)

	if isArray {
		valueArray := result["__value"].([]any)
		values := make([]any, len(valueArray))
		for i, value := range valueArray {
			values[i] = makeValue(value, fieldType, strings.TrimSuffix(fieldTypeStr, ">"))
		}
		f.value = values
	} else {
		f.value = makeValue(result["__value"], fieldType, fieldTypeStr)
	}

	f.isArray = isArray
	return nil
}

func makeValue(valueData any, fieldType FieldTypeV2, fieldTypeStr string) any {
	value := valueData
	if fieldType == FieldTypeInt {
		value = int(value.(float64))
	} else if fieldType == FieldTypeColor {
		hex := valueData.(string)
		value, _ = parseHexString(hex)
	} else if fieldType == FieldTypeEnum {
		enumName := getEnumName(fieldTypeStr)
		valueStr, _ := valueData.(string)
		value = Enum{
			Name:  enumName,
			Value: valueStr,
		}
	} else if fieldType == FieldTypeTile {
		if tile, ok := toMap(valueData); ok {
			value = TileFieldValue{
				TileSetUID: int(tile["tilesetUid"].(float64)),
				X:          tile["x"].(float64),
				Y:          tile["y"].(float64),
				W:          tile["w"].(float64),
				H:          tile["h"].(float64),
			}
		}
	} else if fieldType == FieldTypeEntityRef {
		if entityRef, ok := toMap(valueData); ok {
			value = EntityRef{
				EntityIid: entityRef["entityIid"].(string),
				LevelIid:  entityRef["levelIid"].(string),
			}
		}
	} else if fieldType == FieldTypePoint {
		if point, ok := toMap(valueData); ok {
			value = Point{
				X: point["cx"].(float64),
				Y: point["cy"].(float64),
			}
		}
	}
	return value
}

func toMap(value any) (map[string]any, bool) {
	if value == nil {
		return make(map[string]any, 0), false
	}
	res, ok := value.(map[string]any)
	if !ok {
		return make(map[string]any, 0), false
	}
	return res, true
}

func parseHexString(hex string) (Color, error) {
	var color Color
	values, err := strconv.ParseUint(string(hex[1:]), 16, 32)

	if err != nil {
		return Color{}, err
	}

	color = Color{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
	}

	return color, nil
}

func getEnumName(enumData string) string {
	return strings.Split(enumData, ".")[1]
}

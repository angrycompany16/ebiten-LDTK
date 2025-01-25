package ebitenLDTK

/*
func setField(
	targetStruct *interface{},
	fieldName string,
	fieldType reflect.Kind,
	JSONdata map[string]interface{},
	fieldNameJSON string,
) error {
	field1 := reflect.ValueOf(targetStruct).Elem().FieldByName(fieldName)

	if field1 == reflect.Zero() {

	}
	if field1.Kind() != fieldType {
		return fmt.Errorf("error parsing LDTK struct fields: field name [%s] and field type [%s] do not correspond", fieldName, fieldType)
	}

	switch field1.Kind() {
	case reflect.Float64:
		value, ok := JSONdata[fieldNameJSON]
		if !ok {
			return fmt.Errorf("error parsing LDTK JSON fields: field name [%s] does not exist", fieldNameJSON)
		}
		_value, _ok := value.(float64)
		if !_ok {
			return fmt.Errorf("error parsing LDTK JSON fields: field name %s and field type %s do not correspond", fieldNameJSON, fieldType)
		}
		field1.SetFloat(_value)
	}
	return nil
}
*/

/*
func LoadLDTK(path string) (LDTKWorld, error) {
	fields := []struct {
		fieldName string
		fieldType reflect.Kind
		key       string
	}{
		{"GridWidth", reflect.Float64, "worldGridWidth"},
		{"GridHeight", reflect.Float64, "worldGridHeight"},
	}

	w := LDTKWorld{}
	jsonData, err := os.ReadFile(path)
	if err != nil {
		return w, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return w, err
	}

	for _, f := range fields {
		if err := w.setField(f.fieldName, f.fieldType, data, f.key); err != nil {
			return w, err
		}
	}

	return w, nil
}

func testLDTKJSON(path, field string) {
	jsonData, _ := os.ReadFile(path)
	var data map[string]interface{}
	_ = json.Unmarshal(jsonData, &data)
	fmt.Println(reflect.TypeOf(data[field]))
}*/

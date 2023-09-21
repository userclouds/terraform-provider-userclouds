package apitypes

// Map represents a map from string to ValueType.
type Map struct {
	ValueType APIType
}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *Map) TFModelType() string {
	return "types.Map"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *Map) TFSchemaAttributeType() string {
	return `types.MapType {
		ElemType: ` + t.ValueType.TFSchemaAttributeType() + `,
	}`
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *Map) TFSchemaAttributeText(extraFields map[string]string) string {
	if child, ok := t.ValueType.(*Object); ok {
		return `schema.MapNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: ` + child.TFSchemaAttributesMapName + `,
			},
			` + fieldsToString(extraFields) + `
		}`
	}
	return `schema.MapAttribute{
		ElementType: ` + t.ValueType.TFSchemaAttributeType() + `,
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *Map) JSONClientModelType() string {
	return "map[string]" + t.ValueType.JSONClientModelType()
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *Map) TFModelToJSONClientFunc() string {
	return `func (val *` + t.TFModelType() + `) (*` + t.JSONClientModelType() + `, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		out := ` + t.JSONClientModelType() + `{}
		for k, v := range val.Elements() {
			vTyped, ok := v.(` + t.ValueType.TFModelType() + `)
			if !ok {
				return nil, ucerr.Errorf("unexpected value type %s in map", v.Type(context.Background()).String())
			}
			converted, err := ` + t.ValueType.TFModelToJSONClientFunc() + `(&vTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out[k] = *converted
		}
		return &out, nil
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *Map) JSONClientModelToTFFunc() string {
	return `func (val *` + t.JSONClientModelType() + `) (` + t.TFModelType() + `, error) {
		valueAttrType := ` + t.ValueType.TFSchemaAttributeType() + `
		if val == nil {
			return types.MapNull(valueAttrType), nil
		}
		var out = map[string]attr.Value{}
		for k, v := range *val {
			converted, err := ` + t.ValueType.JSONClientModelToTFFunc() + `(&v)
			if err != nil {
				return types.MapNull(valueAttrType), ucerr.Wrap(err)
			}
			out[k] = converted
		}
		return types.MapValueMust(valueAttrType, out), nil
	}`
}

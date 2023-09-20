package apitypes

// Float represents a float value.
type Float struct{}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *Float) TFModelType() string {
	return "types.Float64"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *Float) TFSchemaAttributeType() string {
	return "types.Float64Type"
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *Float) TFSchemaAttributeText(extraFields map[string]string) string {
	return `schema.Float64Attribute{
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *Float) JSONClientModelType() string {
	return "float64"
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *Float) TFModelToJSONClientFunc() string {
	return `func (val *types.Float64) (*float64, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueFloat64()
		return &converted, nil
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *Float) JSONClientModelToTFFunc() string {
	return `func (val *float64) (types.Float64, error) {
		return types.Float64PointerValue(val), nil
	}`
}

package apitypes

import "fmt"

// StringEnum represents an enum of strings.
type StringEnum struct {
	Values []string
}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *StringEnum) TFModelType() string {
	return "types.String"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *StringEnum) TFSchemaAttributeType() string {
	return "types.StringType"
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *StringEnum) TFSchemaAttributeText(extraFields map[string]string) string {
	return `schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf(` + fmt.Sprintf("%#v", t.Values) + `...),
		},
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *StringEnum) JSONClientModelType() string {
	return "string"
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *StringEnum) TFModelToJSONClientFunc() string {
	return `func (val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *StringEnum) JSONClientModelToTFFunc() string {
	return `func (val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}`
}

package apitypes

import "fmt"

// Bool represents a boolean value.
type Bool struct {
	Description *string
}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *Bool) TFModelType() string {
	return "types.Bool"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *Bool) TFSchemaAttributeType() string {
	return "types.BoolType"
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *Bool) TFSchemaAttributeText(extraFields map[string]string) string {
	return `schema.BoolAttribute{
		Description: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		MarkdownDescription: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *Bool) JSONClientModelType() string {
	return "bool"
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *Bool) TFModelToJSONClientFunc() string {
	return `func (val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *Bool) JSONClientModelToTFFunc() string {
	return `func (val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}`
}

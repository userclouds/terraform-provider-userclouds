package apitypes

import "fmt"

// Int represents an integer value.
type Int struct {
	Description *string
}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *Int) TFModelType() string {
	return "types.Int64"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *Int) TFSchemaAttributeType() string {
	return "types.Int64Type"
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *Int) TFSchemaAttributeText(extraFields map[string]string) string {
	return `schema.Int64Attribute{
		Description: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		MarkdownDescription: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *Int) JSONClientModelType() string {
	return "int64"
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *Int) TFModelToJSONClientFunc() string {
	return `func (val *types.Int64) (*int64, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueInt64()
		return &converted, nil
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *Int) JSONClientModelToTFFunc() string {
	return `func (val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}`
}

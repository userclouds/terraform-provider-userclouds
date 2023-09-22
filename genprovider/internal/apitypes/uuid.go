package apitypes

import "fmt"

const uuidRegex = `(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$`

// UUID represents a UUIDv4 string.
type UUID struct {
	Description *string
}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *UUID) TFModelType() string {
	return "types.String"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *UUID) TFSchemaAttributeType() string {
	return "types.StringType"
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *UUID) TFSchemaAttributeText(extraFields map[string]string) string {
	return `schema.StringAttribute{
		Description: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		MarkdownDescription: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("` + uuidRegex + `"),
				"invalid UUIDv4 format",
			),
		},
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *UUID) JSONClientModelType() string {
	return "uuid.UUID"
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *UUID) TFModelToJSONClientFunc() string {
	return `func (val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *UUID) JSONClientModelToTFFunc() string {
	return `func (val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}`
}

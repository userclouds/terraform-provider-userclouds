package apitypes

import (
	"fmt"
	"strings"
)

// StringEnum represents an enum of strings.
type StringEnum struct {
	Description *string
	Values      []string
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

func (t *StringEnum) description() string {
	desc := "Valid values: "
	for i, v := range t.Values {
		desc += "`" + v + "`"
		if i < len(t.Values)-1 {
			desc += ", "
		}
	}
	if t.Description != nil && *t.Description != "" {
		if !strings.HasSuffix(*t.Description, ".") {
			desc = ". " + desc
		}
		desc = *t.Description + desc
	}
	return desc
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *StringEnum) TFSchemaAttributeText(extraFields map[string]string) string {
	return `schema.StringAttribute{
		Description: ` + fmt.Sprintf("%#v", t.description()) + `,
		MarkdownDescription: ` + fmt.Sprintf("%#v", t.description()) + `,
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

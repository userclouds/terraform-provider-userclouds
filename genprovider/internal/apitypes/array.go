package apitypes

import "fmt"

// Array represents an array of a single type.
type Array struct {
	Description *string
	ChildType   APIType
}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *Array) TFModelType() string {
	return "types.List"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *Array) TFSchemaAttributeType() string {
	return `types.ListType {
		ElemType: ` + t.ChildType.TFSchemaAttributeType() + `,
	}`
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *Array) TFSchemaAttributeText(extraFields map[string]string) string {
	if child, ok := t.ChildType.(*Object); ok {
		return `schema.ListNestedAttribute{
			Description: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
			MarkdownDescription: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
			NestedObject: schema.NestedAttributeObject{
				Attributes: ` + child.TFSchemaAttributesMapName + `,
			},
			` + fieldsToString(extraFields) + `
		}`
	}
	return `schema.ListAttribute{
		Description: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		MarkdownDescription: ` + fmt.Sprintf("%#v", defaultDescription(t.Description)) + `,
		ElementType: ` + t.ChildType.TFSchemaAttributeType() + `,
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *Array) JSONClientModelType() string {
	return "[]" + t.ChildType.JSONClientModelType()
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *Array) TFModelToJSONClientFunc() string {
	return `func (val *` + t.TFModelType() + `) (*` + t.JSONClientModelType() + `, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out ` + t.JSONClientModelType() + `
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(` + t.ChildType.TFModelType() + `)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := ` + t.ChildType.TFModelToJSONClientFunc() + `(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *Array) JSONClientModelToTFFunc() string {
	return `func (val *` + t.JSONClientModelType() + `) (` + t.TFModelType() + `, error) {
		childAttrType := ` + t.ChildType.TFSchemaAttributeType() + `
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := ` + t.ChildType.JSONClientModelToTFFunc() + `(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}`
}

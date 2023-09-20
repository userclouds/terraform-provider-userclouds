package apitypes

// Object represents an object with known property types.
type Object struct {
	// We store the object properties in structs generated by schemas.go. We
	// need to reference those structs in the code generated here
	TFModelStructName         string
	TFSchemaAttrTypeMapName   string
	TFSchemaAttributesMapName string
	JSONClientModelStructName string
	// We also generate functions in schemas.go that iterate over the object
	// properties in order to convert between model types
	TFModelToJSONClientFuncName string
	JSONClientModelToTFFuncName string
}

// TFModelType returns the type that should be used to represent this type in a Terraform model.
func (t *Object) TFModelType() string {
	return "types.Object"
}

// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
// object in the Terraform schema.
func (t *Object) TFSchemaAttributeType() string {
	return `types.ObjectType{
		AttrTypes: ` + t.TFSchemaAttrTypeMapName + `,
	}`
}

// TFSchemaAttributeText returns the text of the code for instantiating this type as a Terraform
// schema attribute.
func (t *Object) TFSchemaAttributeText(extraFields map[string]string) string {
	return `schema.SingleNestedAttribute{
		Attributes: ` + t.TFSchemaAttributesMapName + `,
		` + fieldsToString(extraFields) + `
	}`
}

// JSONClientModelType returns the type that should be used to represent this type in a jsonclient
// request/response struct.
func (t *Object) JSONClientModelType() string {
	return t.JSONClientModelStructName
}

// TFModelToJSONClientFunc returns the text of a function for converting a Terraform model struct to
// a jsonclient model struct.
func (t *Object) TFModelToJSONClientFunc() string {
	return `func (val *` + t.TFModelType() + `) (*` + t.JSONClientModelType() + `, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := ` + t.TFModelStructName + `{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return ` + t.TFModelToJSONClientFuncName + `(&tfModel)
	}`
}

// JSONClientModelToTFFunc returns the text of a function for converting a jsonclient model struct
// to a Terraform model struct.
func (t *Object) JSONClientModelToTFFunc() string {
	return `func (val *` + t.JSONClientModelType() + `) (` + t.TFModelType() + `, error) {
		attrTypes := ` + t.TFSchemaAttrTypeMapName + `

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := ` + t.JSONClientModelToTFFuncName + `(val)
		if err != nil {
			return types.ObjectNull(attrTypes), ucerr.Wrap(err)
		}

		v := reflect.ValueOf(tfModel)

		attrVals := map[string]attr.Value{}
		for i := 0; i < v.NumField(); i++ {
			attrVals[v.Type().Field(i).Tag.Get("tfsdk")] = v.Field(i).Interface().(attr.Value)
		}

		objVal, diag := types.ObjectValue(attrTypes, attrVals)
		if diag.ErrorsCount() > 0 {
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert ` + t.TFModelStructName + ` to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}`
}

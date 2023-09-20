package apitypes

import (
	"testing"
)

func TestString(t *testing.T) {
	runTestProgram(t, data{
		T:                     &String{},
		SampleJSONClientValue: `"hello"`,
		SampleTFModelValue:    `types.StringValue("hello")`,
	})
}

func TestInt(t *testing.T) {
	runTestProgram(t, data{
		T:                     &Int{},
		SampleJSONClientValue: `int64(123)`,
		SampleTFModelValue:    `types.Int64Value(123)`,
	})
}

func TestFloat(t *testing.T) {
	runTestProgram(t, data{
		T:                     &Float{},
		SampleJSONClientValue: `float64(1.23)`,
		SampleTFModelValue:    `types.Float64Value(1.23)`,
	})
}

func TestBool(t *testing.T) {
	runTestProgram(t, data{
		T:                     &Bool{},
		SampleJSONClientValue: `true`,
		SampleTFModelValue:    `types.BoolValue(true)`,
	})
}

var objectSampleCode = `
		type ObjTFModel struct {
			Name types.String ` + "`tfsdk:\"name\"`" + `
		}

		type ObjJSONClientModel struct {
			Name *string ` + "`json:\"name\"`" + `
		}

		var ObjAttributes = map[string]schema.Attribute {
			"name": schema.StringAttribute{},
		}

		var ObjSchemaAttrTypes = map[string]attr.Type {
			"name": types.StringType,
		}

		func ObjTFModelToJSONClient(val *ObjTFModel) (*ObjJSONClientModel, error) {
			return &ObjJSONClientModel{Name: val.Name.ValueStringPointer()}, nil
		}

		func ObjJSONClientModelToTF(val *ObjJSONClientModel) (ObjTFModel, error) {
			return ObjTFModel{Name: types.StringPointerValue(val.Name)}, nil
		}

		// define sample string so that we can get a pointer to it
		var sampleString = "hello"
		var sampleJSONClientValue = ObjJSONClientModel{Name: &sampleString}
		var sampleTFModelAttrTypes = map[string]attr.Type{"name": basetypes.StringType{}}
		var sampleTFModelValue = types.ObjectValueMust(sampleTFModelAttrTypes, map[string]attr.Value{"name": types.StringValue(sampleString)})
`

func TestObject(t *testing.T) {
	runTestProgram(t, data{
		ExtraCode: objectSampleCode,
		T: &Object{
			TFModelStructName:           "ObjTFModel",
			TFSchemaAttributesMapName:   "ObjAttributes",
			TFSchemaAttrTypeMapName:     "ObjSchemaAttrTypes",
			JSONClientModelStructName:   "ObjJSONClientModel",
			TFModelToJSONClientFuncName: "ObjTFModelToJSONClient",
			JSONClientModelToTFFuncName: "ObjJSONClientModelToTF",
		},
		SampleJSONClientValue: "sampleJSONClientValue",
		SampleTFModelValue:    "sampleTFModelValue",
	})
}

func TestStringArray(t *testing.T) {
	runTestProgram(t, data{
		T:                     &Array{ChildType: &String{}},
		SampleJSONClientValue: `[]string{"hello", "world"}`,
		SampleTFModelValue:    `types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hello"), types.StringValue("world")})`,
	})
}

func TestObjectArray(t *testing.T) {
	runTestProgram(t, data{
		ExtraCode: objectSampleCode,
		T: &Array{
			ChildType: &Object{
				TFModelStructName:           "ObjTFModel",
				TFSchemaAttributesMapName:   "ObjAttributes",
				TFSchemaAttrTypeMapName:     "ObjSchemaAttrTypes",
				JSONClientModelStructName:   "ObjJSONClientModel",
				TFModelToJSONClientFuncName: "ObjTFModelToJSONClient",
				JSONClientModelToTFFuncName: "ObjJSONClientModelToTF",
			},
		},
		SampleJSONClientValue: "[]ObjJSONClientModel{sampleJSONClientValue}",
		SampleTFModelValue:    `types.ListValueMust(sampleTFModelValue.Type(context.Background()), []attr.Value{sampleTFModelValue})`,
	})
}

func TestStringMap(t *testing.T) {
	runTestProgram(t, data{
		T:                     &Map{ValueType: &String{}},
		SampleJSONClientValue: `map[string]string{"a": "A", "b": "B"}`,
		SampleTFModelValue:    `types.MapValueMust(types.StringType, map[string]attr.Value{"a": types.StringValue("A"), "b": types.StringValue("B")})`,
	})
}

func TestObjectMap(t *testing.T) {
	runTestProgram(t, data{
		ExtraCode: objectSampleCode,
		T: &Map{
			ValueType: &Object{
				TFModelStructName:           "ObjTFModel",
				TFSchemaAttributesMapName:   "ObjAttributes",
				TFSchemaAttrTypeMapName:     "ObjSchemaAttrTypes",
				JSONClientModelStructName:   "ObjJSONClientModel",
				TFModelToJSONClientFuncName: "ObjTFModelToJSONClient",
				JSONClientModelToTFFuncName: "ObjJSONClientModelToTF",
			},
		},
		SampleJSONClientValue: `map[string]ObjJSONClientModel{"a": sampleJSONClientValue}`,
		SampleTFModelValue:    `types.MapValueMust(sampleTFModelValue.Type(context.Background()), map[string]attr.Value{"a": sampleTFModelValue})`,
	})
}

func TestUUID(t *testing.T) {
	runTestProgram(t, data{
		T:                     &UUID{},
		SampleJSONClientValue: `uuid.Must(uuid.FromString("123e4567-e89b-12d3-a456-426655440000"))`,
		SampleTFModelValue:    `types.StringValue("123e4567-e89b-12d3-a456-426655440000")`,
	})
}

func TestStringEnum(t *testing.T) {
	// NOTE: this just tests that the generated StringEnum code compiles. It
	// doesn't test the enum validation logic, since that's specified as a
	// Terraform attribute validator, and that's harder to test.
	runTestProgram(t, data{
		T:                     &StringEnum{Values: []string{"a", "b", "c"}},
		SampleJSONClientValue: `"hello"`,
		SampleTFModelValue:    `types.StringValue("hello")`,
	})
}

var userstoreResourceIDSampleCode = `
		type UserstoreResourceIDJSONClientModel struct {
			ID   *uuid.UUID
			Name *string
		}

		// define sample string so that we can get a pointer to it
		var sampleID = uuid.Must(uuid.FromString("123e4567-e89b-12d3-a456-426655440000"))
		var sampleJSONClientValue = UserstoreResourceIDJSONClientModel{ID: &sampleID}
		var sampleTFModelValue = types.StringValue(sampleID.String())
`

func TestUserstoreResourceID(t *testing.T) {
	runTestProgram(t, data{
		ExtraCode: userstoreResourceIDSampleCode,
		T: &UserstoreResourceID{
			JSONClientModelStructName: "UserstoreResourceIDJSONClientModel",
		},
		SampleJSONClientValue: "sampleJSONClientValue",
		SampleTFModelValue:    "sampleTFModelValue",
	})
}

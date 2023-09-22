package apitypes

import (
	"sort"
	"strings"
)

// APIType represents a type for use with the UserClouds API. It contains methods to simplify
// mapping between various Terraform types and Go types.
type APIType interface {
	// TFModelType returns the name of the corresponding type in the
	// terraform-provider-userclouds/types package, used for in-memory storage
	TFModelType() string
	// TFSchemaAttributeType returns an attr.Type, which is a struct representing the type of this
	// object in the Terraform schema.
	TFSchemaAttributeType() string
	// TFSchemaAttributeText generates the declaration for an attribute storing
	// this type in a Terraform schema
	TFSchemaAttributeText(extraFields map[string]string) string
	// JSONClientModelType returns the name of the type for use with
	// userclouds/jsonclient (most likely simple Go native types)
	JSONClientModelType() string
	// TFModelToJSONClientFunc should return the text of a function with
	// signature "func (val *TFModelType) (*JSONClientModelType, error)"
	TFModelToJSONClientFunc() string
	// TFModelToJSONClientFunc should return the text of a function with
	// signature "func (val *JSONClientModelType) (TFModelType, error)"
	JSONClientModelToTFFunc() string
}

func fieldsToString(fields map[string]string) string {
	keys := []string{}
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := ""
	for _, k := range keys {
		out += k + ": " + fields[k] + ",\n"
	}
	return strings.TrimSpace(out)
}

func defaultDescription(description *string) string {
	if description == nil {
		return ""
	}
	return *description
}

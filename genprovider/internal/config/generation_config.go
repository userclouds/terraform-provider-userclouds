package config

// ResourceConfig describes a resource that should be generated.
type ResourceConfig struct {
	TypeNameSuffix      string `json:"type_name_suffix"`
	Description         string `json:"description"`
	MarkdownDescription string `json:"markdown_description"`
	OpenapiSchema       string `json:"openapi_schema"`
	RestCollectionPath  string `json:"rest_collection_path"`
	RestResourcePath    string `json:"rest_resource_path"`
	// Map path param names (e.g. "id") to the name of the OpenAPI schema
	// property whose value should be used to replace the path param in URLs.
	// Schema property names should be lower_snake_case
	PathParamsToSchemaProperty map[string]string `json:"path_params_to_schema_property"`
}

// SchemaOverride allows us to tweak OpenAPI schemas before we generate Terraform schemas from them.
type SchemaOverride struct {
	// Mark these schema properties as read-only. In theory, we should just mark everything readonly
	// via OpenAPI, but this doesn't work with object references:
	// https://usercloudsworkspace.slack.com/archives/C02A3HELPPU/p1695418020075049
	// Furthermore, swaggest doesn't currently support a `readOnly` struct tag :(
	ReadonlyProperties []string `json:"readonly_properties"`
}

// Spec configures the generation of resources from a single OpenAPI spec.
type Spec struct {
	File                 string                    `json:"file"`
	GeneratedFilePackage string                    `json:"generated_file_package"`
	Resources            []ResourceConfig          `json:"resources"`
	SchemaOverrides      map[string]SchemaOverride `json:"schema_overrides"`
}

// GenerationConfig configures the generation of the Terraform provider.
type GenerationConfig struct {
	Specs []Spec `json:"specs"`
}

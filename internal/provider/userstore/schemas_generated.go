// NOTE: automatically generated file -- DO NOT EDIT

package userstore

import (
	"context"
	"reflect"
	"regexp"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/userclouds/terraform-provider-userclouds/internal/provider/planmodifiers"

	"userclouds.com/infra/ucerr"
)

// Note: revive is complaining about stuttering in the generated schema names (e.g. an OpenAPI
// schema might be called "UserstoreColumnTFModel", and then we generate it in the "userstore"
// package, so it becomes "userstore.UserstoreColumnTFModel"), but these names are coming from the
// OpenAPI spec and no one is going to be reading this generated code anyways, so we should just
// silence the rule.
//revive:disable:exported

// IdpColumnConsentedPurposesTFModel is a Terraform model struct for the IdpColumnConsentedPurposesAttributes schema.
type IdpColumnConsentedPurposesTFModel struct {
	Column            types.String `tfsdk:"column"`
	ConsentedPurposes types.List   `tfsdk:"consented_purposes"`
}

// IdpColumnConsentedPurposesJSONClientModel stores data for use with jsonclient for making API requests.
type IdpColumnConsentedPurposesJSONClientModel struct {
	Column            *UserstoreResourceIDJSONClientModel   `json:"column,omitempty"`
	ConsentedPurposes *[]UserstoreResourceIDJSONClientModel `json:"consented_purposes,omitempty"`
}

// IdpColumnConsentedPurposesAttrTypes defines the attribute types for the IdpColumnConsentedPurposesAttributes schema.
var IdpColumnConsentedPurposesAttrTypes = map[string]attr.Type{
	"column": types.StringType,
	"consented_purposes": types.ListType{
		ElemType: types.StringType,
	},
}

// IdpColumnConsentedPurposesAttributes defines the Terraform attributes schema.
var IdpColumnConsentedPurposesAttributes = map[string]schema.Attribute{
	"column": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"consented_purposes": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpColumnConsentedPurposesTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpColumnConsentedPurposesTFModelToJSONClient(in *IdpColumnConsentedPurposesTFModel) (*IdpColumnConsentedPurposesJSONClientModel, error) {
	out := IdpColumnConsentedPurposesJSONClientModel{}
	var err error
	out.Column, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.Column)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	out.ConsentedPurposes, err = func(val *types.List) (*[]UserstoreResourceIDJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []UserstoreResourceIDJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.String)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
				if val.IsNull() || val.IsUnknown() {
					return nil, nil
				}
				converted, err := uuid.FromString(val.ValueString())
				if err != nil {
					return nil, ucerr.Errorf("failed to parse uuid: %v", err)
				}
				s := UserstoreResourceIDJSONClientModel{
					ID: &converted,
				}
				return &s, nil
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.ConsentedPurposes)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"consented_purposes\" field: %+v", err)
	}
	return &out, nil
}

// IdpColumnConsentedPurposesJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpColumnConsentedPurposesJSONClientModelToTF(in *IdpColumnConsentedPurposesJSONClientModel) (IdpColumnConsentedPurposesTFModel, error) {
	out := IdpColumnConsentedPurposesTFModel{}
	var err error
	out.Column, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.Column)
	if err != nil {
		return IdpColumnConsentedPurposesTFModel{}, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	out.ConsentedPurposes, err = func(val *[]UserstoreResourceIDJSONClientModel) (types.List, error) {
		childAttrType := types.StringType
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
				if val == nil {
					return types.StringNull(), nil
				}
				// We should only need to convert jsonclient models to TF models when receiving API
				// responses, and API responses should always have the ID set.
				// Sometimes we receive nil UUIDs here because of how the server
				// serializes empty values, so we should only freak out if we see a
				// name provided but not an ID.
				if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
					return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
				}
				if val.ID == nil || val.ID.IsNil() {
					return types.StringNull(), nil
				}
				return types.StringValue(val.ID.String()), nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.ConsentedPurposes)
	if err != nil {
		return IdpColumnConsentedPurposesTFModel{}, ucerr.Errorf("failed to convert \"consented_purposes\" field: %+v", err)
	}
	return out, nil
}

// IdpColumnRetentionDurationTFModel is a Terraform model struct for the IdpColumnRetentionDurationAttributes schema.
type IdpColumnRetentionDurationTFModel struct {
	ColumnID        types.String `tfsdk:"column_id"`
	DefaultDuration types.Object `tfsdk:"default_duration"`
	Duration        types.Object `tfsdk:"duration"`
	DurationType    types.String `tfsdk:"duration_type"`
	ID              types.String `tfsdk:"id"`
	PurposeID       types.String `tfsdk:"purpose_id"`
	PurposeName     types.String `tfsdk:"purpose_name"`
	UseDefault      types.Bool   `tfsdk:"use_default"`
	Version         types.Int64  `tfsdk:"version"`
}

// IdpColumnRetentionDurationJSONClientModel stores data for use with jsonclient for making API requests.
type IdpColumnRetentionDurationJSONClientModel struct {
	ColumnID        *uuid.UUID                           `json:"column_id,omitempty"`
	DefaultDuration *IdpRetentionDurationJSONClientModel `json:"default_duration,omitempty"`
	Duration        *IdpRetentionDurationJSONClientModel `json:"duration,omitempty"`
	DurationType    *string                              `json:"duration_type,omitempty"`
	ID              *uuid.UUID                           `json:"id,omitempty"`
	PurposeID       *uuid.UUID                           `json:"purpose_id,omitempty"`
	PurposeName     *string                              `json:"purpose_name,omitempty"`
	UseDefault      *bool                                `json:"use_default,omitempty"`
	Version         *int64                               `json:"version,omitempty"`
}

// IdpColumnRetentionDurationAttrTypes defines the attribute types for the IdpColumnRetentionDurationAttributes schema.
var IdpColumnRetentionDurationAttrTypes = map[string]attr.Type{
	"column_id": types.StringType,
	"default_duration": types.ObjectType{
		AttrTypes: IdpRetentionDurationAttrTypes,
	},
	"duration": types.ObjectType{
		AttrTypes: IdpRetentionDurationAttrTypes,
	},
	"duration_type": types.StringType,
	"id":            types.StringType,
	"purpose_id":    types.StringType,
	"purpose_name":  types.StringType,
	"use_default":   types.BoolType,
	"version":       types.Int64Type,
}

// IdpColumnRetentionDurationAttributes defines the Terraform attributes schema.
var IdpColumnRetentionDurationAttributes = map[string]schema.Attribute{
	"column_id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"default_duration": schema.SingleNestedAttribute{
		Attributes:          IdpRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"duration": schema.SingleNestedAttribute{
		Attributes:          IdpRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"duration_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"", "live", "postdelete", "predelete", "softdeleted"}...),
		},
		Computed:            true,
		Description:         "Valid values: ``, `live`, `postdelete`, `predelete`, `softdeleted`",
		MarkdownDescription: "Valid values: ``, `live`, `postdelete`, `predelete`, `softdeleted`",
		Optional:            true,
	},
	"id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"purpose_id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"purpose_name": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"use_default": schema.BoolAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"version": schema.Int64Attribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		PlanModifiers: []planmodifier.Int64{
			planmodifiers.IncrementOnUpdate(),
		},
	},
}

// IdpColumnRetentionDurationTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpColumnRetentionDurationTFModelToJSONClient(in *IdpColumnRetentionDurationTFModel) (*IdpColumnRetentionDurationJSONClientModel, error) {
	out := IdpColumnRetentionDurationJSONClientModel{}
	var err error
	out.ColumnID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.ColumnID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"column_id\" field: %+v", err)
	}
	out.DefaultDuration, err = func(val *types.Object) (*IdpRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.DefaultDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"default_duration\" field: %+v", err)
	}
	out.Duration, err = func(val *types.Object) (*IdpRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.Duration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"duration\" field: %+v", err)
	}
	out.DurationType, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.DurationType)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"duration_type\" field: %+v", err)
	}
	out.ID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.ID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.PurposeID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.PurposeID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"purpose_id\" field: %+v", err)
	}
	out.PurposeName, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.PurposeName)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"purpose_name\" field: %+v", err)
	}
	out.UseDefault, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.UseDefault)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"use_default\" field: %+v", err)
	}
	out.Version, err = func(val *types.Int64) (*int64, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueInt64()
		return &converted, nil
	}(&in.Version)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
	return &out, nil
}

// IdpColumnRetentionDurationJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpColumnRetentionDurationJSONClientModelToTF(in *IdpColumnRetentionDurationJSONClientModel) (IdpColumnRetentionDurationTFModel, error) {
	out := IdpColumnRetentionDurationTFModel{}
	var err error
	out.ColumnID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ColumnID)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"column_id\" field: %+v", err)
	}
	out.DefaultDuration, err = func(val *IdpRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.DefaultDuration)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"default_duration\" field: %+v", err)
	}
	out.Duration, err = func(val *IdpRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Duration)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"duration\" field: %+v", err)
	}
	out.DurationType, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.DurationType)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"duration_type\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.PurposeID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.PurposeID)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"purpose_id\" field: %+v", err)
	}
	out.PurposeName, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.PurposeName)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"purpose_name\" field: %+v", err)
	}
	out.UseDefault, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.UseDefault)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"use_default\" field: %+v", err)
	}
	out.Version, err = func(val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}(in.Version)
	if err != nil {
		return IdpColumnRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
	return out, nil
}

// IdpColumnRetentionDurationResponseTFModel is a Terraform model struct for the IdpColumnRetentionDurationResponseAttributes schema.
type IdpColumnRetentionDurationResponseTFModel struct {
	MaxDuration       types.Object `tfsdk:"max_duration"`
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// IdpColumnRetentionDurationResponseJSONClientModel stores data for use with jsonclient for making API requests.
type IdpColumnRetentionDurationResponseJSONClientModel struct {
	MaxDuration       *IdpRetentionDurationJSONClientModel       `json:"max_duration,omitempty"`
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// IdpColumnRetentionDurationResponseAttrTypes defines the attribute types for the IdpColumnRetentionDurationResponseAttributes schema.
var IdpColumnRetentionDurationResponseAttrTypes = map[string]attr.Type{
	"max_duration": types.ObjectType{
		AttrTypes: IdpRetentionDurationAttrTypes,
	},
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// IdpColumnRetentionDurationResponseAttributes defines the Terraform attributes schema.
var IdpColumnRetentionDurationResponseAttributes = map[string]schema.Attribute{
	"max_duration": schema.SingleNestedAttribute{
		Attributes:          IdpRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpColumnRetentionDurationResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpColumnRetentionDurationResponseTFModelToJSONClient(in *IdpColumnRetentionDurationResponseTFModel) (*IdpColumnRetentionDurationResponseJSONClientModel, error) {
	out := IdpColumnRetentionDurationResponseJSONClientModel{}
	var err error
	out.MaxDuration, err = func(val *types.Object) (*IdpRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.MaxDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"max_duration\" field: %+v", err)
	}
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// IdpColumnRetentionDurationResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpColumnRetentionDurationResponseJSONClientModelToTF(in *IdpColumnRetentionDurationResponseJSONClientModel) (IdpColumnRetentionDurationResponseTFModel, error) {
	out := IdpColumnRetentionDurationResponseTFModel{}
	var err error
	out.MaxDuration, err = func(val *IdpRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.MaxDuration)
	if err != nil {
		return IdpColumnRetentionDurationResponseTFModel{}, ucerr.Errorf("failed to convert \"max_duration\" field: %+v", err)
	}
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return IdpColumnRetentionDurationResponseTFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// IdpColumnRetentionDurationsResponseTFModel is a Terraform model struct for the IdpColumnRetentionDurationsResponseAttributes schema.
type IdpColumnRetentionDurationsResponseTFModel struct {
	MaxDuration        types.Object `tfsdk:"max_duration"`
	RetentionDurations types.List   `tfsdk:"retention_durations"`
}

// IdpColumnRetentionDurationsResponseJSONClientModel stores data for use with jsonclient for making API requests.
type IdpColumnRetentionDurationsResponseJSONClientModel struct {
	MaxDuration        *IdpRetentionDurationJSONClientModel         `json:"max_duration,omitempty"`
	RetentionDurations *[]IdpColumnRetentionDurationJSONClientModel `json:"retention_durations,omitempty"`
}

// IdpColumnRetentionDurationsResponseAttrTypes defines the attribute types for the IdpColumnRetentionDurationsResponseAttributes schema.
var IdpColumnRetentionDurationsResponseAttrTypes = map[string]attr.Type{
	"max_duration": types.ObjectType{
		AttrTypes: IdpRetentionDurationAttrTypes,
	},
	"retention_durations": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: IdpColumnRetentionDurationAttrTypes,
		},
	},
}

// IdpColumnRetentionDurationsResponseAttributes defines the Terraform attributes schema.
var IdpColumnRetentionDurationsResponseAttributes = map[string]schema.Attribute{
	"max_duration": schema.SingleNestedAttribute{
		Attributes:          IdpRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"retention_durations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IdpColumnRetentionDurationAttributes,
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpColumnRetentionDurationsResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpColumnRetentionDurationsResponseTFModelToJSONClient(in *IdpColumnRetentionDurationsResponseTFModel) (*IdpColumnRetentionDurationsResponseJSONClientModel, error) {
	out := IdpColumnRetentionDurationsResponseJSONClientModel{}
	var err error
	out.MaxDuration, err = func(val *types.Object) (*IdpRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.MaxDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"max_duration\" field: %+v", err)
	}
	out.RetentionDurations, err = func(val *types.List) (*[]IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []IdpColumnRetentionDurationJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := IdpColumnRetentionDurationTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.RetentionDurations)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_durations\" field: %+v", err)
	}
	return &out, nil
}

// IdpColumnRetentionDurationsResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpColumnRetentionDurationsResponseJSONClientModelToTF(in *IdpColumnRetentionDurationsResponseJSONClientModel) (IdpColumnRetentionDurationsResponseTFModel, error) {
	out := IdpColumnRetentionDurationsResponseTFModel{}
	var err error
	out.MaxDuration, err = func(val *IdpRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.MaxDuration)
	if err != nil {
		return IdpColumnRetentionDurationsResponseTFModel{}, ucerr.Errorf("failed to convert \"max_duration\" field: %+v", err)
	}
	out.RetentionDurations, err = func(val *[]IdpColumnRetentionDurationJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: IdpColumnRetentionDurationAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
				attrTypes := IdpColumnRetentionDurationAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.RetentionDurations)
	if err != nil {
		return IdpColumnRetentionDurationsResponseTFModel{}, ucerr.Errorf("failed to convert \"retention_durations\" field: %+v", err)
	}
	return out, nil
}

// IdpCreateAccessorRequestTFModel is a Terraform model struct for the IdpCreateAccessorRequestAttributes schema.
type IdpCreateAccessorRequestTFModel struct {
	Accessor types.Object `tfsdk:"accessor"`
}

// IdpCreateAccessorRequestJSONClientModel stores data for use with jsonclient for making API requests.
type IdpCreateAccessorRequestJSONClientModel struct {
	Accessor *UserstoreAccessorJSONClientModel `json:"accessor,omitempty"`
}

// IdpCreateAccessorRequestAttrTypes defines the attribute types for the IdpCreateAccessorRequestAttributes schema.
var IdpCreateAccessorRequestAttrTypes = map[string]attr.Type{
	"accessor": types.ObjectType{
		AttrTypes: UserstoreAccessorAttrTypes,
	},
}

// IdpCreateAccessorRequestAttributes defines the Terraform attributes schema.
var IdpCreateAccessorRequestAttributes = map[string]schema.Attribute{
	"accessor": schema.SingleNestedAttribute{
		Attributes:          UserstoreAccessorAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpCreateAccessorRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpCreateAccessorRequestTFModelToJSONClient(in *IdpCreateAccessorRequestTFModel) (*IdpCreateAccessorRequestJSONClientModel, error) {
	out := IdpCreateAccessorRequestJSONClientModel{}
	var err error
	out.Accessor, err = func(val *types.Object) (*UserstoreAccessorJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreAccessorTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreAccessorTFModelToJSONClient(&tfModel)
	}(&in.Accessor)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"accessor\" field: %+v", err)
	}
	return &out, nil
}

// IdpCreateAccessorRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpCreateAccessorRequestJSONClientModelToTF(in *IdpCreateAccessorRequestJSONClientModel) (IdpCreateAccessorRequestTFModel, error) {
	out := IdpCreateAccessorRequestTFModel{}
	var err error
	out.Accessor, err = func(val *UserstoreAccessorJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreAccessorAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreAccessorJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreAccessorTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Accessor)
	if err != nil {
		return IdpCreateAccessorRequestTFModel{}, ucerr.Errorf("failed to convert \"accessor\" field: %+v", err)
	}
	return out, nil
}

// IdpCreateColumnRequestTFModel is a Terraform model struct for the IdpCreateColumnRequestAttributes schema.
type IdpCreateColumnRequestTFModel struct {
	Column types.Object `tfsdk:"column"`
}

// IdpCreateColumnRequestJSONClientModel stores data for use with jsonclient for making API requests.
type IdpCreateColumnRequestJSONClientModel struct {
	Column *UserstoreColumnJSONClientModel `json:"column,omitempty"`
}

// IdpCreateColumnRequestAttrTypes defines the attribute types for the IdpCreateColumnRequestAttributes schema.
var IdpCreateColumnRequestAttrTypes = map[string]attr.Type{
	"column": types.ObjectType{
		AttrTypes: UserstoreColumnAttrTypes,
	},
}

// IdpCreateColumnRequestAttributes defines the Terraform attributes schema.
var IdpCreateColumnRequestAttributes = map[string]schema.Attribute{
	"column": schema.SingleNestedAttribute{
		Attributes:          UserstoreColumnAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpCreateColumnRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpCreateColumnRequestTFModelToJSONClient(in *IdpCreateColumnRequestTFModel) (*IdpCreateColumnRequestJSONClientModel, error) {
	out := IdpCreateColumnRequestJSONClientModel{}
	var err error
	out.Column, err = func(val *types.Object) (*UserstoreColumnJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreColumnTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreColumnTFModelToJSONClient(&tfModel)
	}(&in.Column)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	return &out, nil
}

// IdpCreateColumnRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpCreateColumnRequestJSONClientModelToTF(in *IdpCreateColumnRequestJSONClientModel) (IdpCreateColumnRequestTFModel, error) {
	out := IdpCreateColumnRequestTFModel{}
	var err error
	out.Column, err = func(val *UserstoreColumnJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreColumnAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreColumnJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreColumnTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Column)
	if err != nil {
		return IdpCreateColumnRequestTFModel{}, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	return out, nil
}

// IdpCreateMutatorRequestTFModel is a Terraform model struct for the IdpCreateMutatorRequestAttributes schema.
type IdpCreateMutatorRequestTFModel struct {
	Mutator types.Object `tfsdk:"mutator"`
}

// IdpCreateMutatorRequestJSONClientModel stores data for use with jsonclient for making API requests.
type IdpCreateMutatorRequestJSONClientModel struct {
	Mutator *UserstoreMutatorJSONClientModel `json:"mutator,omitempty"`
}

// IdpCreateMutatorRequestAttrTypes defines the attribute types for the IdpCreateMutatorRequestAttributes schema.
var IdpCreateMutatorRequestAttrTypes = map[string]attr.Type{
	"mutator": types.ObjectType{
		AttrTypes: UserstoreMutatorAttrTypes,
	},
}

// IdpCreateMutatorRequestAttributes defines the Terraform attributes schema.
var IdpCreateMutatorRequestAttributes = map[string]schema.Attribute{
	"mutator": schema.SingleNestedAttribute{
		Attributes:          UserstoreMutatorAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpCreateMutatorRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpCreateMutatorRequestTFModelToJSONClient(in *IdpCreateMutatorRequestTFModel) (*IdpCreateMutatorRequestJSONClientModel, error) {
	out := IdpCreateMutatorRequestJSONClientModel{}
	var err error
	out.Mutator, err = func(val *types.Object) (*UserstoreMutatorJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreMutatorTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreMutatorTFModelToJSONClient(&tfModel)
	}(&in.Mutator)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"mutator\" field: %+v", err)
	}
	return &out, nil
}

// IdpCreateMutatorRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpCreateMutatorRequestJSONClientModelToTF(in *IdpCreateMutatorRequestJSONClientModel) (IdpCreateMutatorRequestTFModel, error) {
	out := IdpCreateMutatorRequestTFModel{}
	var err error
	out.Mutator, err = func(val *UserstoreMutatorJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreMutatorAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreMutatorJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreMutatorTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Mutator)
	if err != nil {
		return IdpCreateMutatorRequestTFModel{}, ucerr.Errorf("failed to convert \"mutator\" field: %+v", err)
	}
	return out, nil
}

// IdpCreatePurposeRequestTFModel is a Terraform model struct for the IdpCreatePurposeRequestAttributes schema.
type IdpCreatePurposeRequestTFModel struct {
	Purpose types.Object `tfsdk:"purpose"`
}

// IdpCreatePurposeRequestJSONClientModel stores data for use with jsonclient for making API requests.
type IdpCreatePurposeRequestJSONClientModel struct {
	Purpose *UserstorePurposeJSONClientModel `json:"purpose,omitempty"`
}

// IdpCreatePurposeRequestAttrTypes defines the attribute types for the IdpCreatePurposeRequestAttributes schema.
var IdpCreatePurposeRequestAttrTypes = map[string]attr.Type{
	"purpose": types.ObjectType{
		AttrTypes: UserstorePurposeAttrTypes,
	},
}

// IdpCreatePurposeRequestAttributes defines the Terraform attributes schema.
var IdpCreatePurposeRequestAttributes = map[string]schema.Attribute{
	"purpose": schema.SingleNestedAttribute{
		Attributes:          UserstorePurposeAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpCreatePurposeRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpCreatePurposeRequestTFModelToJSONClient(in *IdpCreatePurposeRequestTFModel) (*IdpCreatePurposeRequestJSONClientModel, error) {
	out := IdpCreatePurposeRequestJSONClientModel{}
	var err error
	out.Purpose, err = func(val *types.Object) (*UserstorePurposeJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstorePurposeTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstorePurposeTFModelToJSONClient(&tfModel)
	}(&in.Purpose)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"purpose\" field: %+v", err)
	}
	return &out, nil
}

// IdpCreatePurposeRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpCreatePurposeRequestJSONClientModelToTF(in *IdpCreatePurposeRequestJSONClientModel) (IdpCreatePurposeRequestTFModel, error) {
	out := IdpCreatePurposeRequestTFModel{}
	var err error
	out.Purpose, err = func(val *UserstorePurposeJSONClientModel) (types.Object, error) {
		attrTypes := UserstorePurposeAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstorePurposeJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstorePurposeTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Purpose)
	if err != nil {
		return IdpCreatePurposeRequestTFModel{}, ucerr.Errorf("failed to convert \"purpose\" field: %+v", err)
	}
	return out, nil
}

// IdpDurationUnitTFModel is a Terraform model struct for the IdpDurationUnitAttributes schema.
type IdpDurationUnitTFModel struct {
}

// IdpDurationUnitJSONClientModel stores data for use with jsonclient for making API requests.
type IdpDurationUnitJSONClientModel struct {
}

// IdpDurationUnitAttrTypes defines the attribute types for the IdpDurationUnitAttributes schema.
var IdpDurationUnitAttrTypes = map[string]attr.Type{}

// IdpDurationUnitAttributes defines the Terraform attributes schema.
var IdpDurationUnitAttributes = map[string]schema.Attribute{}

// IdpDurationUnitTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpDurationUnitTFModelToJSONClient(in *IdpDurationUnitTFModel) (*IdpDurationUnitJSONClientModel, error) {
	out := IdpDurationUnitJSONClientModel{}
	return &out, nil
}

// IdpDurationUnitJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpDurationUnitJSONClientModelToTF(in *IdpDurationUnitJSONClientModel) (IdpDurationUnitTFModel, error) {
	out := IdpDurationUnitTFModel{}
	return out, nil
}

// IdpExecuteAccessorResponseTFModel is a Terraform model struct for the IdpExecuteAccessorResponseAttributes schema.
type IdpExecuteAccessorResponseTFModel struct {
	Data types.List `tfsdk:"data"`
}

// IdpExecuteAccessorResponseJSONClientModel stores data for use with jsonclient for making API requests.
type IdpExecuteAccessorResponseJSONClientModel struct {
	Data *[]string `json:"data,omitempty"`
}

// IdpExecuteAccessorResponseAttrTypes defines the attribute types for the IdpExecuteAccessorResponseAttributes schema.
var IdpExecuteAccessorResponseAttrTypes = map[string]attr.Type{
	"data": types.ListType{
		ElemType: types.StringType,
	},
}

// IdpExecuteAccessorResponseAttributes defines the Terraform attributes schema.
var IdpExecuteAccessorResponseAttributes = map[string]schema.Attribute{
	"data": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpExecuteAccessorResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpExecuteAccessorResponseTFModelToJSONClient(in *IdpExecuteAccessorResponseTFModel) (*IdpExecuteAccessorResponseJSONClientModel, error) {
	out := IdpExecuteAccessorResponseJSONClientModel{}
	var err error
	out.Data, err = func(val *types.List) (*[]string, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []string{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.String)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.String) (*string, error) {
				if val.IsNull() || val.IsUnknown() {
					return nil, nil
				}
				converted := val.ValueString()
				return &converted, nil
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Data)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	return &out, nil
}

// IdpExecuteAccessorResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpExecuteAccessorResponseJSONClientModelToTF(in *IdpExecuteAccessorResponseJSONClientModel) (IdpExecuteAccessorResponseTFModel, error) {
	out := IdpExecuteAccessorResponseTFModel{}
	var err error
	out.Data, err = func(val *[]string) (types.List, error) {
		childAttrType := types.StringType
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *string) (types.String, error) {
				return types.StringPointerValue(val), nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Data)
	if err != nil {
		return IdpExecuteAccessorResponseTFModel{}, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	return out, nil
}

// IdpExecuteMutatorResponseTFModel is a Terraform model struct for the IdpExecuteMutatorResponseAttributes schema.
type IdpExecuteMutatorResponseTFModel struct {
	UserIds types.List `tfsdk:"user_ids"`
}

// IdpExecuteMutatorResponseJSONClientModel stores data for use with jsonclient for making API requests.
type IdpExecuteMutatorResponseJSONClientModel struct {
	UserIds *[]uuid.UUID `json:"user_ids,omitempty"`
}

// IdpExecuteMutatorResponseAttrTypes defines the attribute types for the IdpExecuteMutatorResponseAttributes schema.
var IdpExecuteMutatorResponseAttrTypes = map[string]attr.Type{
	"user_ids": types.ListType{
		ElemType: types.StringType,
	},
}

// IdpExecuteMutatorResponseAttributes defines the Terraform attributes schema.
var IdpExecuteMutatorResponseAttributes = map[string]schema.Attribute{
	"user_ids": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpExecuteMutatorResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpExecuteMutatorResponseTFModelToJSONClient(in *IdpExecuteMutatorResponseTFModel) (*IdpExecuteMutatorResponseJSONClientModel, error) {
	out := IdpExecuteMutatorResponseJSONClientModel{}
	var err error
	out.UserIds, err = func(val *types.List) (*[]uuid.UUID, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []uuid.UUID{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.String)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.String) (*uuid.UUID, error) {
				if val.IsNull() || val.IsUnknown() {
					return nil, nil
				}
				converted, err := uuid.FromString(val.ValueString())
				if err != nil {
					return nil, ucerr.Errorf("failed to parse uuid: %v", err)
				}
				return &converted, nil
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.UserIds)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"user_ids\" field: %+v", err)
	}
	return &out, nil
}

// IdpExecuteMutatorResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpExecuteMutatorResponseJSONClientModelToTF(in *IdpExecuteMutatorResponseJSONClientModel) (IdpExecuteMutatorResponseTFModel, error) {
	out := IdpExecuteMutatorResponseTFModel{}
	var err error
	out.UserIds, err = func(val *[]uuid.UUID) (types.List, error) {
		childAttrType := types.StringType
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *uuid.UUID) (types.String, error) {
				if val == nil {
					return types.StringNull(), nil
				}
				return types.StringValue(val.String()), nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.UserIds)
	if err != nil {
		return IdpExecuteMutatorResponseTFModel{}, ucerr.Errorf("failed to convert \"user_ids\" field: %+v", err)
	}
	return out, nil
}

// IdpGetConsentedPurposesForUserRequestTFModel is a Terraform model struct for the IdpGetConsentedPurposesForUserRequestAttributes schema.
type IdpGetConsentedPurposesForUserRequestTFModel struct {
	Columns types.List   `tfsdk:"columns"`
	UserID  types.String `tfsdk:"user_id"`
}

// IdpGetConsentedPurposesForUserRequestJSONClientModel stores data for use with jsonclient for making API requests.
type IdpGetConsentedPurposesForUserRequestJSONClientModel struct {
	Columns *[]UserstoreResourceIDJSONClientModel `json:"columns,omitempty"`
	UserID  *uuid.UUID                            `json:"user_id,omitempty"`
}

// IdpGetConsentedPurposesForUserRequestAttrTypes defines the attribute types for the IdpGetConsentedPurposesForUserRequestAttributes schema.
var IdpGetConsentedPurposesForUserRequestAttrTypes = map[string]attr.Type{
	"columns": types.ListType{
		ElemType: types.StringType,
	},
	"user_id": types.StringType,
}

// IdpGetConsentedPurposesForUserRequestAttributes defines the Terraform attributes schema.
var IdpGetConsentedPurposesForUserRequestAttributes = map[string]schema.Attribute{
	"columns": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"user_id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpGetConsentedPurposesForUserRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpGetConsentedPurposesForUserRequestTFModelToJSONClient(in *IdpGetConsentedPurposesForUserRequestTFModel) (*IdpGetConsentedPurposesForUserRequestJSONClientModel, error) {
	out := IdpGetConsentedPurposesForUserRequestJSONClientModel{}
	var err error
	out.Columns, err = func(val *types.List) (*[]UserstoreResourceIDJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []UserstoreResourceIDJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.String)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
				if val.IsNull() || val.IsUnknown() {
					return nil, nil
				}
				converted, err := uuid.FromString(val.ValueString())
				if err != nil {
					return nil, ucerr.Errorf("failed to parse uuid: %v", err)
				}
				s := UserstoreResourceIDJSONClientModel{
					ID: &converted,
				}
				return &s, nil
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Columns)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"columns\" field: %+v", err)
	}
	out.UserID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.UserID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"user_id\" field: %+v", err)
	}
	return &out, nil
}

// IdpGetConsentedPurposesForUserRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpGetConsentedPurposesForUserRequestJSONClientModelToTF(in *IdpGetConsentedPurposesForUserRequestJSONClientModel) (IdpGetConsentedPurposesForUserRequestTFModel, error) {
	out := IdpGetConsentedPurposesForUserRequestTFModel{}
	var err error
	out.Columns, err = func(val *[]UserstoreResourceIDJSONClientModel) (types.List, error) {
		childAttrType := types.StringType
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
				if val == nil {
					return types.StringNull(), nil
				}
				// We should only need to convert jsonclient models to TF models when receiving API
				// responses, and API responses should always have the ID set.
				// Sometimes we receive nil UUIDs here because of how the server
				// serializes empty values, so we should only freak out if we see a
				// name provided but not an ID.
				if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
					return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
				}
				if val.ID == nil || val.ID.IsNil() {
					return types.StringNull(), nil
				}
				return types.StringValue(val.ID.String()), nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Columns)
	if err != nil {
		return IdpGetConsentedPurposesForUserRequestTFModel{}, ucerr.Errorf("failed to convert \"columns\" field: %+v", err)
	}
	out.UserID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.UserID)
	if err != nil {
		return IdpGetConsentedPurposesForUserRequestTFModel{}, ucerr.Errorf("failed to convert \"user_id\" field: %+v", err)
	}
	return out, nil
}

// IdpGetConsentedPurposesForUserResponseTFModel is a Terraform model struct for the IdpGetConsentedPurposesForUserResponseAttributes schema.
type IdpGetConsentedPurposesForUserResponseTFModel struct {
	Data types.List `tfsdk:"data"`
}

// IdpGetConsentedPurposesForUserResponseJSONClientModel stores data for use with jsonclient for making API requests.
type IdpGetConsentedPurposesForUserResponseJSONClientModel struct {
	Data *[]IdpColumnConsentedPurposesJSONClientModel `json:"data,omitempty"`
}

// IdpGetConsentedPurposesForUserResponseAttrTypes defines the attribute types for the IdpGetConsentedPurposesForUserResponseAttributes schema.
var IdpGetConsentedPurposesForUserResponseAttrTypes = map[string]attr.Type{
	"data": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: IdpColumnConsentedPurposesAttrTypes,
		},
	},
}

// IdpGetConsentedPurposesForUserResponseAttributes defines the Terraform attributes schema.
var IdpGetConsentedPurposesForUserResponseAttributes = map[string]schema.Attribute{
	"data": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IdpColumnConsentedPurposesAttributes,
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpGetConsentedPurposesForUserResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpGetConsentedPurposesForUserResponseTFModelToJSONClient(in *IdpGetConsentedPurposesForUserResponseTFModel) (*IdpGetConsentedPurposesForUserResponseJSONClientModel, error) {
	out := IdpGetConsentedPurposesForUserResponseJSONClientModel{}
	var err error
	out.Data, err = func(val *types.List) (*[]IdpColumnConsentedPurposesJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []IdpColumnConsentedPurposesJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*IdpColumnConsentedPurposesJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := IdpColumnConsentedPurposesTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return IdpColumnConsentedPurposesTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Data)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	return &out, nil
}

// IdpGetConsentedPurposesForUserResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpGetConsentedPurposesForUserResponseJSONClientModelToTF(in *IdpGetConsentedPurposesForUserResponseJSONClientModel) (IdpGetConsentedPurposesForUserResponseTFModel, error) {
	out := IdpGetConsentedPurposesForUserResponseTFModel{}
	var err error
	out.Data, err = func(val *[]IdpColumnConsentedPurposesJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: IdpColumnConsentedPurposesAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *IdpColumnConsentedPurposesJSONClientModel) (types.Object, error) {
				attrTypes := IdpColumnConsentedPurposesAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := IdpColumnConsentedPurposesJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnConsentedPurposesTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Data)
	if err != nil {
		return IdpGetConsentedPurposesForUserResponseTFModel{}, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	return out, nil
}

// IdpRetentionDurationTFModel is a Terraform model struct for the IdpRetentionDurationAttributes schema.
type IdpRetentionDurationTFModel struct {
	Duration types.Int64  `tfsdk:"duration"`
	Unit     types.String `tfsdk:"unit"`
}

// IdpRetentionDurationJSONClientModel stores data for use with jsonclient for making API requests.
type IdpRetentionDurationJSONClientModel struct {
	Duration *int64  `json:"duration,omitempty"`
	Unit     *string `json:"unit,omitempty"`
}

// IdpRetentionDurationAttrTypes defines the attribute types for the IdpRetentionDurationAttributes schema.
var IdpRetentionDurationAttrTypes = map[string]attr.Type{
	"duration": types.Int64Type,
	"unit":     types.StringType,
}

// IdpRetentionDurationAttributes defines the Terraform attributes schema.
var IdpRetentionDurationAttributes = map[string]schema.Attribute{
	"duration": schema.Int64Attribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"unit": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"day", "hour", "indefinite", "month", "week", "year"}...),
		},
		Computed:            true,
		Description:         "Valid values: `day`, `hour`, `indefinite`, `month`, `week`, `year`",
		MarkdownDescription: "Valid values: `day`, `hour`, `indefinite`, `month`, `week`, `year`",
		Optional:            true,
	},
}

// IdpRetentionDurationTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpRetentionDurationTFModelToJSONClient(in *IdpRetentionDurationTFModel) (*IdpRetentionDurationJSONClientModel, error) {
	out := IdpRetentionDurationJSONClientModel{}
	var err error
	out.Duration, err = func(val *types.Int64) (*int64, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueInt64()
		return &converted, nil
	}(&in.Duration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"duration\" field: %+v", err)
	}
	out.Unit, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Unit)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"unit\" field: %+v", err)
	}
	return &out, nil
}

// IdpRetentionDurationJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpRetentionDurationJSONClientModelToTF(in *IdpRetentionDurationJSONClientModel) (IdpRetentionDurationTFModel, error) {
	out := IdpRetentionDurationTFModel{}
	var err error
	out.Duration, err = func(val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}(in.Duration)
	if err != nil {
		return IdpRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"duration\" field: %+v", err)
	}
	out.Unit, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Unit)
	if err != nil {
		return IdpRetentionDurationTFModel{}, ucerr.Errorf("failed to convert \"unit\" field: %+v", err)
	}
	return out, nil
}

// IdpUpdateColumnRetentionDurationRequestTFModel is a Terraform model struct for the IdpUpdateColumnRetentionDurationRequestAttributes schema.
type IdpUpdateColumnRetentionDurationRequestTFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// IdpUpdateColumnRetentionDurationRequestJSONClientModel stores data for use with jsonclient for making API requests.
type IdpUpdateColumnRetentionDurationRequestJSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// IdpUpdateColumnRetentionDurationRequestAttrTypes defines the attribute types for the IdpUpdateColumnRetentionDurationRequestAttributes schema.
var IdpUpdateColumnRetentionDurationRequestAttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// IdpUpdateColumnRetentionDurationRequestAttributes defines the Terraform attributes schema.
var IdpUpdateColumnRetentionDurationRequestAttributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// IdpUpdateColumnRetentionDurationRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func IdpUpdateColumnRetentionDurationRequestTFModelToJSONClient(in *IdpUpdateColumnRetentionDurationRequestTFModel) (*IdpUpdateColumnRetentionDurationRequestJSONClientModel, error) {
	out := IdpUpdateColumnRetentionDurationRequestJSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// IdpUpdateColumnRetentionDurationRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func IdpUpdateColumnRetentionDurationRequestJSONClientModelToTF(in *IdpUpdateColumnRetentionDurationRequestJSONClientModel) (IdpUpdateColumnRetentionDurationRequestTFModel, error) {
	out := IdpUpdateColumnRetentionDurationRequestTFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return IdpUpdateColumnRetentionDurationRequestTFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// PaginationCursorTFModel is a Terraform model struct for the PaginationCursorAttributes schema.
type PaginationCursorTFModel struct {
}

// PaginationCursorJSONClientModel stores data for use with jsonclient for making API requests.
type PaginationCursorJSONClientModel struct {
}

// PaginationCursorAttrTypes defines the attribute types for the PaginationCursorAttributes schema.
var PaginationCursorAttrTypes = map[string]attr.Type{}

// PaginationCursorAttributes defines the Terraform attributes schema.
var PaginationCursorAttributes = map[string]schema.Attribute{}

// PaginationCursorTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PaginationCursorTFModelToJSONClient(in *PaginationCursorTFModel) (*PaginationCursorJSONClientModel, error) {
	out := PaginationCursorJSONClientModel{}
	return &out, nil
}

// PaginationCursorJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PaginationCursorJSONClientModelToTF(in *PaginationCursorJSONClientModel) (PaginationCursorTFModel, error) {
	out := PaginationCursorTFModel{}
	return out, nil
}

// PolicyClientContextTFModel is a Terraform model struct for the PolicyClientContextAttributes schema.
type PolicyClientContextTFModel struct {
}

// PolicyClientContextJSONClientModel stores data for use with jsonclient for making API requests.
type PolicyClientContextJSONClientModel struct {
}

// PolicyClientContextAttrTypes defines the attribute types for the PolicyClientContextAttributes schema.
var PolicyClientContextAttrTypes = map[string]attr.Type{}

// PolicyClientContextAttributes defines the Terraform attributes schema.
var PolicyClientContextAttributes = map[string]schema.Attribute{}

// PolicyClientContextTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PolicyClientContextTFModelToJSONClient(in *PolicyClientContextTFModel) (*PolicyClientContextJSONClientModel, error) {
	out := PolicyClientContextJSONClientModel{}
	return &out, nil
}

// PolicyClientContextJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PolicyClientContextJSONClientModelToTF(in *PolicyClientContextJSONClientModel) (PolicyClientContextTFModel, error) {
	out := PolicyClientContextTFModel{}
	return out, nil
}

// UserstoreAccessorTFModel is a Terraform model struct for the UserstoreAccessorAttributes schema.
type UserstoreAccessorTFModel struct {
	AccessPolicy       types.String `tfsdk:"access_policy"`
	Columns            types.List   `tfsdk:"columns"`
	DataLifeCycleState types.String `tfsdk:"data_life_cycle_state"`
	Description        types.String `tfsdk:"description"`
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Purposes           types.List   `tfsdk:"purposes"`
	SelectorConfig     types.Object `tfsdk:"selector_config"`
	TokenAccessPolicy  types.String `tfsdk:"token_access_policy"`
	Version            types.Int64  `tfsdk:"version"`
}

// UserstoreAccessorJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreAccessorJSONClientModel struct {
	AccessPolicy       *UserstoreResourceIDJSONClientModel           `json:"access_policy,omitempty"`
	Columns            *[]UserstoreColumnOutputConfigJSONClientModel `json:"columns,omitempty"`
	DataLifeCycleState *string                                       `json:"data_life_cycle_state,omitempty"`
	Description        *string                                       `json:"description,omitempty"`
	ID                 *uuid.UUID                                    `json:"id,omitempty"`
	Name               *string                                       `json:"name,omitempty"`
	Purposes           *[]UserstoreResourceIDJSONClientModel         `json:"purposes,omitempty"`
	SelectorConfig     *UserstoreUserSelectorConfigJSONClientModel   `json:"selector_config,omitempty"`
	TokenAccessPolicy  *UserstoreResourceIDJSONClientModel           `json:"token_access_policy,omitempty"`
	Version            *int64                                        `json:"version,omitempty"`
}

// UserstoreAccessorAttrTypes defines the attribute types for the UserstoreAccessorAttributes schema.
var UserstoreAccessorAttrTypes = map[string]attr.Type{
	"access_policy": types.StringType,
	"columns": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: UserstoreColumnOutputConfigAttrTypes,
		},
	},
	"data_life_cycle_state": types.StringType,
	"description":           types.StringType,
	"id":                    types.StringType,
	"name":                  types.StringType,
	"purposes": types.ListType{
		ElemType: types.StringType,
	},
	"selector_config": types.ObjectType{
		AttrTypes: UserstoreUserSelectorConfigAttrTypes,
	},
	"token_access_policy": types.StringType,
	"version":             types.Int64Type,
}

// UserstoreAccessorAttributes defines the Terraform attributes schema.
var UserstoreAccessorAttributes = map[string]schema.Attribute{
	"access_policy": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"columns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UserstoreColumnOutputConfigAttributes,
		},
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"data_life_cycle_state": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"", "live", "postdelete", "predelete", "softdeleted"}...),
		},
		Computed:            true,
		Description:         "Valid values: ``, `live`, `postdelete`, `predelete`, `softdeleted`",
		MarkdownDescription: "Valid values: ``, `live`, `postdelete`, `predelete`, `softdeleted`",
		Optional:            true,
	},
	"description": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"purposes": schema.ListAttribute{
		ElementType:         types.StringType,
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"selector_config": schema.SingleNestedAttribute{
		Attributes:          UserstoreUserSelectorConfigAttributes,
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"token_access_policy": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"version": schema.Int64Attribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		PlanModifiers: []planmodifier.Int64{
			planmodifiers.IncrementOnUpdate(),
		},
	},
}

// UserstoreAccessorTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreAccessorTFModelToJSONClient(in *UserstoreAccessorTFModel) (*UserstoreAccessorJSONClientModel, error) {
	out := UserstoreAccessorJSONClientModel{}
	var err error
	out.AccessPolicy, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.AccessPolicy)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	out.Columns, err = func(val *types.List) (*[]UserstoreColumnOutputConfigJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []UserstoreColumnOutputConfigJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*UserstoreColumnOutputConfigJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := UserstoreColumnOutputConfigTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return UserstoreColumnOutputConfigTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Columns)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"columns\" field: %+v", err)
	}
	out.DataLifeCycleState, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.DataLifeCycleState)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"data_life_cycle_state\" field: %+v", err)
	}
	out.Description, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Description)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.ID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.ID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Name)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.Purposes, err = func(val *types.List) (*[]UserstoreResourceIDJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []UserstoreResourceIDJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.String)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
				if val.IsNull() || val.IsUnknown() {
					return nil, nil
				}
				converted, err := uuid.FromString(val.ValueString())
				if err != nil {
					return nil, ucerr.Errorf("failed to parse uuid: %v", err)
				}
				s := UserstoreResourceIDJSONClientModel{
					ID: &converted,
				}
				return &s, nil
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Purposes)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"purposes\" field: %+v", err)
	}
	out.SelectorConfig, err = func(val *types.Object) (*UserstoreUserSelectorConfigJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreUserSelectorConfigTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreUserSelectorConfigTFModelToJSONClient(&tfModel)
	}(&in.SelectorConfig)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"selector_config\" field: %+v", err)
	}
	out.TokenAccessPolicy, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.TokenAccessPolicy)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"token_access_policy\" field: %+v", err)
	}
	out.Version, err = func(val *types.Int64) (*int64, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueInt64()
		return &converted, nil
	}(&in.Version)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreAccessorJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreAccessorJSONClientModelToTF(in *UserstoreAccessorJSONClientModel) (UserstoreAccessorTFModel, error) {
	out := UserstoreAccessorTFModel{}
	var err error
	out.AccessPolicy, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.AccessPolicy)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	out.Columns, err = func(val *[]UserstoreColumnOutputConfigJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: UserstoreColumnOutputConfigAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *UserstoreColumnOutputConfigJSONClientModel) (types.Object, error) {
				attrTypes := UserstoreColumnOutputConfigAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := UserstoreColumnOutputConfigJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreColumnOutputConfigTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Columns)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"columns\" field: %+v", err)
	}
	out.DataLifeCycleState, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.DataLifeCycleState)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"data_life_cycle_state\" field: %+v", err)
	}
	out.Description, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Description)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.Purposes, err = func(val *[]UserstoreResourceIDJSONClientModel) (types.List, error) {
		childAttrType := types.StringType
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
				if val == nil {
					return types.StringNull(), nil
				}
				// We should only need to convert jsonclient models to TF models when receiving API
				// responses, and API responses should always have the ID set.
				// Sometimes we receive nil UUIDs here because of how the server
				// serializes empty values, so we should only freak out if we see a
				// name provided but not an ID.
				if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
					return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
				}
				if val.ID == nil || val.ID.IsNil() {
					return types.StringNull(), nil
				}
				return types.StringValue(val.ID.String()), nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Purposes)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"purposes\" field: %+v", err)
	}
	out.SelectorConfig, err = func(val *UserstoreUserSelectorConfigJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreUserSelectorConfigAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreUserSelectorConfigJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreUserSelectorConfigTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.SelectorConfig)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"selector_config\" field: %+v", err)
	}
	out.TokenAccessPolicy, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.TokenAccessPolicy)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"token_access_policy\" field: %+v", err)
	}
	out.Version, err = func(val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}(in.Version)
	if err != nil {
		return UserstoreAccessorTFModel{}, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
	return out, nil
}

// UserstoreColumnTFModel is a Terraform model struct for the UserstoreColumnAttributes schema.
type UserstoreColumnTFModel struct {
	DefaultValue types.String `tfsdk:"default_value"`
	ID           types.String `tfsdk:"id"`
	IndexType    types.String `tfsdk:"index_type"`
	IsArray      types.Bool   `tfsdk:"is_array"`
	Name         types.String `tfsdk:"name"`
	Type         types.String `tfsdk:"type"`
}

// UserstoreColumnJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreColumnJSONClientModel struct {
	DefaultValue *string    `json:"default_value,omitempty"`
	ID           *uuid.UUID `json:"id,omitempty"`
	IndexType    *string    `json:"index_type,omitempty"`
	IsArray      *bool      `json:"is_array,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Type         *string    `json:"type,omitempty"`
}

// UserstoreColumnAttrTypes defines the attribute types for the UserstoreColumnAttributes schema.
var UserstoreColumnAttrTypes = map[string]attr.Type{
	"default_value": types.StringType,
	"id":            types.StringType,
	"index_type":    types.StringType,
	"is_array":      types.BoolType,
	"name":          types.StringType,
	"type":          types.StringType,
}

// UserstoreColumnAttributes defines the Terraform attributes schema.
var UserstoreColumnAttributes = map[string]schema.Attribute{
	"default_value": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"index_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"indexed", "none", "unique"}...),
		},
		Description:         "Valid values: `indexed`, `none`, `unique`",
		MarkdownDescription: "Valid values: `indexed`, `none`, `unique`",
		Required:            true,
	},
	"is_array": schema.BoolAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"name": schema.StringAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"address", "birthdate", "boolean", "date", "e164_phonenumber", "email", "integer", "phonenumber", "ssn", "string", "timestamp", "uuid"}...),
		},
		Description:         "Valid values: `address`, `birthdate`, `boolean`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		MarkdownDescription: "Valid values: `address`, `birthdate`, `boolean`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		Required:            true,
	},
}

// UserstoreColumnTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreColumnTFModelToJSONClient(in *UserstoreColumnTFModel) (*UserstoreColumnJSONClientModel, error) {
	out := UserstoreColumnJSONClientModel{}
	var err error
	out.DefaultValue, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.DefaultValue)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"default_value\" field: %+v", err)
	}
	out.ID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.ID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.IndexType, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.IndexType)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"index_type\" field: %+v", err)
	}
	out.IsArray, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.IsArray)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"is_array\" field: %+v", err)
	}
	out.Name, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Name)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.Type, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Type)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"type\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreColumnJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreColumnJSONClientModelToTF(in *UserstoreColumnJSONClientModel) (UserstoreColumnTFModel, error) {
	out := UserstoreColumnTFModel{}
	var err error
	out.DefaultValue, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.DefaultValue)
	if err != nil {
		return UserstoreColumnTFModel{}, ucerr.Errorf("failed to convert \"default_value\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return UserstoreColumnTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.IndexType, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.IndexType)
	if err != nil {
		return UserstoreColumnTFModel{}, ucerr.Errorf("failed to convert \"index_type\" field: %+v", err)
	}
	out.IsArray, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.IsArray)
	if err != nil {
		return UserstoreColumnTFModel{}, ucerr.Errorf("failed to convert \"is_array\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return UserstoreColumnTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.Type, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Type)
	if err != nil {
		return UserstoreColumnTFModel{}, ucerr.Errorf("failed to convert \"type\" field: %+v", err)
	}
	return out, nil
}

// UserstoreColumnIndexTypeTFModel is a Terraform model struct for the UserstoreColumnIndexTypeAttributes schema.
type UserstoreColumnIndexTypeTFModel struct {
}

// UserstoreColumnIndexTypeJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreColumnIndexTypeJSONClientModel struct {
}

// UserstoreColumnIndexTypeAttrTypes defines the attribute types for the UserstoreColumnIndexTypeAttributes schema.
var UserstoreColumnIndexTypeAttrTypes = map[string]attr.Type{}

// UserstoreColumnIndexTypeAttributes defines the Terraform attributes schema.
var UserstoreColumnIndexTypeAttributes = map[string]schema.Attribute{}

// UserstoreColumnIndexTypeTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreColumnIndexTypeTFModelToJSONClient(in *UserstoreColumnIndexTypeTFModel) (*UserstoreColumnIndexTypeJSONClientModel, error) {
	out := UserstoreColumnIndexTypeJSONClientModel{}
	return &out, nil
}

// UserstoreColumnIndexTypeJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreColumnIndexTypeJSONClientModelToTF(in *UserstoreColumnIndexTypeJSONClientModel) (UserstoreColumnIndexTypeTFModel, error) {
	out := UserstoreColumnIndexTypeTFModel{}
	return out, nil
}

// UserstoreColumnInputConfigTFModel is a Terraform model struct for the UserstoreColumnInputConfigAttributes schema.
type UserstoreColumnInputConfigTFModel struct {
	Column    types.String `tfsdk:"column"`
	Validator types.String `tfsdk:"validator"`
}

// UserstoreColumnInputConfigJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreColumnInputConfigJSONClientModel struct {
	Column    *UserstoreResourceIDJSONClientModel `json:"column,omitempty"`
	Validator *UserstoreResourceIDJSONClientModel `json:"validator,omitempty"`
}

// UserstoreColumnInputConfigAttrTypes defines the attribute types for the UserstoreColumnInputConfigAttributes schema.
var UserstoreColumnInputConfigAttrTypes = map[string]attr.Type{
	"column":    types.StringType,
	"validator": types.StringType,
}

// UserstoreColumnInputConfigAttributes defines the Terraform attributes schema.
var UserstoreColumnInputConfigAttributes = map[string]schema.Attribute{
	"column": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"validator": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreColumnInputConfigTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreColumnInputConfigTFModelToJSONClient(in *UserstoreColumnInputConfigTFModel) (*UserstoreColumnInputConfigJSONClientModel, error) {
	out := UserstoreColumnInputConfigJSONClientModel{}
	var err error
	out.Column, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.Column)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	out.Validator, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.Validator)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"validator\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreColumnInputConfigJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreColumnInputConfigJSONClientModelToTF(in *UserstoreColumnInputConfigJSONClientModel) (UserstoreColumnInputConfigTFModel, error) {
	out := UserstoreColumnInputConfigTFModel{}
	var err error
	out.Column, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.Column)
	if err != nil {
		return UserstoreColumnInputConfigTFModel{}, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	out.Validator, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.Validator)
	if err != nil {
		return UserstoreColumnInputConfigTFModel{}, ucerr.Errorf("failed to convert \"validator\" field: %+v", err)
	}
	return out, nil
}

// UserstoreColumnOutputConfigTFModel is a Terraform model struct for the UserstoreColumnOutputConfigAttributes schema.
type UserstoreColumnOutputConfigTFModel struct {
	Column      types.String `tfsdk:"column"`
	Transformer types.String `tfsdk:"transformer"`
}

// UserstoreColumnOutputConfigJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreColumnOutputConfigJSONClientModel struct {
	Column      *UserstoreResourceIDJSONClientModel `json:"column,omitempty"`
	Transformer *UserstoreResourceIDJSONClientModel `json:"transformer,omitempty"`
}

// UserstoreColumnOutputConfigAttrTypes defines the attribute types for the UserstoreColumnOutputConfigAttributes schema.
var UserstoreColumnOutputConfigAttrTypes = map[string]attr.Type{
	"column":      types.StringType,
	"transformer": types.StringType,
}

// UserstoreColumnOutputConfigAttributes defines the Terraform attributes schema.
var UserstoreColumnOutputConfigAttributes = map[string]schema.Attribute{
	"column": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"transformer": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreColumnOutputConfigTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreColumnOutputConfigTFModelToJSONClient(in *UserstoreColumnOutputConfigTFModel) (*UserstoreColumnOutputConfigJSONClientModel, error) {
	out := UserstoreColumnOutputConfigJSONClientModel{}
	var err error
	out.Column, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.Column)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	out.Transformer, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.Transformer)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"transformer\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreColumnOutputConfigJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreColumnOutputConfigJSONClientModelToTF(in *UserstoreColumnOutputConfigJSONClientModel) (UserstoreColumnOutputConfigTFModel, error) {
	out := UserstoreColumnOutputConfigTFModel{}
	var err error
	out.Column, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.Column)
	if err != nil {
		return UserstoreColumnOutputConfigTFModel{}, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	out.Transformer, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.Transformer)
	if err != nil {
		return UserstoreColumnOutputConfigTFModel{}, ucerr.Errorf("failed to convert \"transformer\" field: %+v", err)
	}
	return out, nil
}

// UserstoreDataLifeCycleStateTFModel is a Terraform model struct for the UserstoreDataLifeCycleStateAttributes schema.
type UserstoreDataLifeCycleStateTFModel struct {
}

// UserstoreDataLifeCycleStateJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreDataLifeCycleStateJSONClientModel struct {
}

// UserstoreDataLifeCycleStateAttrTypes defines the attribute types for the UserstoreDataLifeCycleStateAttributes schema.
var UserstoreDataLifeCycleStateAttrTypes = map[string]attr.Type{}

// UserstoreDataLifeCycleStateAttributes defines the Terraform attributes schema.
var UserstoreDataLifeCycleStateAttributes = map[string]schema.Attribute{}

// UserstoreDataLifeCycleStateTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreDataLifeCycleStateTFModelToJSONClient(in *UserstoreDataLifeCycleStateTFModel) (*UserstoreDataLifeCycleStateJSONClientModel, error) {
	out := UserstoreDataLifeCycleStateJSONClientModel{}
	return &out, nil
}

// UserstoreDataLifeCycleStateJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreDataLifeCycleStateJSONClientModelToTF(in *UserstoreDataLifeCycleStateJSONClientModel) (UserstoreDataLifeCycleStateTFModel, error) {
	out := UserstoreDataLifeCycleStateTFModel{}
	return out, nil
}

// UserstoreDataTypeTFModel is a Terraform model struct for the UserstoreDataTypeAttributes schema.
type UserstoreDataTypeTFModel struct {
}

// UserstoreDataTypeJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreDataTypeJSONClientModel struct {
}

// UserstoreDataTypeAttrTypes defines the attribute types for the UserstoreDataTypeAttributes schema.
var UserstoreDataTypeAttrTypes = map[string]attr.Type{}

// UserstoreDataTypeAttributes defines the Terraform attributes schema.
var UserstoreDataTypeAttributes = map[string]schema.Attribute{}

// UserstoreDataTypeTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreDataTypeTFModelToJSONClient(in *UserstoreDataTypeTFModel) (*UserstoreDataTypeJSONClientModel, error) {
	out := UserstoreDataTypeJSONClientModel{}
	return &out, nil
}

// UserstoreDataTypeJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreDataTypeJSONClientModelToTF(in *UserstoreDataTypeJSONClientModel) (UserstoreDataTypeTFModel, error) {
	out := UserstoreDataTypeTFModel{}
	return out, nil
}

// UserstoreMutatorTFModel is a Terraform model struct for the UserstoreMutatorAttributes schema.
type UserstoreMutatorTFModel struct {
	AccessPolicy   types.String `tfsdk:"access_policy"`
	Columns        types.List   `tfsdk:"columns"`
	Description    types.String `tfsdk:"description"`
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	SelectorConfig types.Object `tfsdk:"selector_config"`
	Version        types.Int64  `tfsdk:"version"`
}

// UserstoreMutatorJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreMutatorJSONClientModel struct {
	AccessPolicy   *UserstoreResourceIDJSONClientModel          `json:"access_policy,omitempty"`
	Columns        *[]UserstoreColumnInputConfigJSONClientModel `json:"columns,omitempty"`
	Description    *string                                      `json:"description,omitempty"`
	ID             *uuid.UUID                                   `json:"id,omitempty"`
	Name           *string                                      `json:"name,omitempty"`
	SelectorConfig *UserstoreUserSelectorConfigJSONClientModel  `json:"selector_config,omitempty"`
	Version        *int64                                       `json:"version,omitempty"`
}

// UserstoreMutatorAttrTypes defines the attribute types for the UserstoreMutatorAttributes schema.
var UserstoreMutatorAttrTypes = map[string]attr.Type{
	"access_policy": types.StringType,
	"columns": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: UserstoreColumnInputConfigAttrTypes,
		},
	},
	"description": types.StringType,
	"id":          types.StringType,
	"name":        types.StringType,
	"selector_config": types.ObjectType{
		AttrTypes: UserstoreUserSelectorConfigAttrTypes,
	},
	"version": types.Int64Type,
}

// UserstoreMutatorAttributes defines the Terraform attributes schema.
var UserstoreMutatorAttributes = map[string]schema.Attribute{
	"access_policy": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"columns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UserstoreColumnInputConfigAttributes,
		},
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"description": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"selector_config": schema.SingleNestedAttribute{
		Attributes:          UserstoreUserSelectorConfigAttributes,
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"version": schema.Int64Attribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		PlanModifiers: []planmodifier.Int64{
			planmodifiers.IncrementOnUpdate(),
		},
	},
}

// UserstoreMutatorTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreMutatorTFModelToJSONClient(in *UserstoreMutatorTFModel) (*UserstoreMutatorJSONClientModel, error) {
	out := UserstoreMutatorJSONClientModel{}
	var err error
	out.AccessPolicy, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		s := UserstoreResourceIDJSONClientModel{
			ID: &converted,
		}
		return &s, nil
	}(&in.AccessPolicy)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	out.Columns, err = func(val *types.List) (*[]UserstoreColumnInputConfigJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []UserstoreColumnInputConfigJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*UserstoreColumnInputConfigJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := UserstoreColumnInputConfigTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return UserstoreColumnInputConfigTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Columns)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"columns\" field: %+v", err)
	}
	out.Description, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Description)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.ID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.ID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Name)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.SelectorConfig, err = func(val *types.Object) (*UserstoreUserSelectorConfigJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreUserSelectorConfigTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreUserSelectorConfigTFModelToJSONClient(&tfModel)
	}(&in.SelectorConfig)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"selector_config\" field: %+v", err)
	}
	out.Version, err = func(val *types.Int64) (*int64, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueInt64()
		return &converted, nil
	}(&in.Version)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreMutatorJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreMutatorJSONClientModelToTF(in *UserstoreMutatorJSONClientModel) (UserstoreMutatorTFModel, error) {
	out := UserstoreMutatorTFModel{}
	var err error
	out.AccessPolicy, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		// We should only need to convert jsonclient models to TF models when receiving API
		// responses, and API responses should always have the ID set.
		// Sometimes we receive nil UUIDs here because of how the server
		// serializes empty values, so we should only freak out if we see a
		// name provided but not an ID.
		if val.Name != nil && *val.Name != "" && (val.ID == nil || val.ID.IsNil()) {
			return types.StringNull(), ucerr.Errorf("got nil ID field in UserstoreResourceID. this is an issue with the UserClouds Terraform provider")
		}
		if val.ID == nil || val.ID.IsNil() {
			return types.StringNull(), nil
		}
		return types.StringValue(val.ID.String()), nil
	}(in.AccessPolicy)
	if err != nil {
		return UserstoreMutatorTFModel{}, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	out.Columns, err = func(val *[]UserstoreColumnInputConfigJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: UserstoreColumnInputConfigAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *UserstoreColumnInputConfigJSONClientModel) (types.Object, error) {
				attrTypes := UserstoreColumnInputConfigAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := UserstoreColumnInputConfigJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreColumnInputConfigTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Columns)
	if err != nil {
		return UserstoreMutatorTFModel{}, ucerr.Errorf("failed to convert \"columns\" field: %+v", err)
	}
	out.Description, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Description)
	if err != nil {
		return UserstoreMutatorTFModel{}, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return UserstoreMutatorTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return UserstoreMutatorTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.SelectorConfig, err = func(val *UserstoreUserSelectorConfigJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreUserSelectorConfigAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreUserSelectorConfigJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreUserSelectorConfigTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.SelectorConfig)
	if err != nil {
		return UserstoreMutatorTFModel{}, ucerr.Errorf("failed to convert \"selector_config\" field: %+v", err)
	}
	out.Version, err = func(val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}(in.Version)
	if err != nil {
		return UserstoreMutatorTFModel{}, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
	return out, nil
}

// UserstorePurposeTFModel is a Terraform model struct for the UserstorePurposeAttributes schema.
type UserstorePurposeTFModel struct {
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
}

// UserstorePurposeJSONClientModel stores data for use with jsonclient for making API requests.
type UserstorePurposeJSONClientModel struct {
	Description *string    `json:"description,omitempty"`
	ID          *uuid.UUID `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
}

// UserstorePurposeAttrTypes defines the attribute types for the UserstorePurposeAttributes schema.
var UserstorePurposeAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"id":          types.StringType,
	"name":        types.StringType,
}

// UserstorePurposeAttributes defines the Terraform attributes schema.
var UserstorePurposeAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
}

// UserstorePurposeTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstorePurposeTFModelToJSONClient(in *UserstorePurposeTFModel) (*UserstorePurposeJSONClientModel, error) {
	out := UserstorePurposeJSONClientModel{}
	var err error
	out.Description, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Description)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.ID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.ID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Name)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	return &out, nil
}

// UserstorePurposeJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstorePurposeJSONClientModelToTF(in *UserstorePurposeJSONClientModel) (UserstorePurposeTFModel, error) {
	out := UserstorePurposeTFModel{}
	var err error
	out.Description, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Description)
	if err != nil {
		return UserstorePurposeTFModel{}, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return UserstorePurposeTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return UserstorePurposeTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	return out, nil
}

// UserstoreResourceIDTFModel is a Terraform model struct for the UserstoreResourceIDAttributes schema.
type UserstoreResourceIDTFModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// UserstoreResourceIDJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreResourceIDJSONClientModel struct {
	ID   *uuid.UUID `json:"id,omitempty"`
	Name *string    `json:"name,omitempty"`
}

// UserstoreResourceIDAttrTypes defines the attribute types for the UserstoreResourceIDAttributes schema.
var UserstoreResourceIDAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"name": types.StringType,
}

// UserstoreResourceIDAttributes defines the Terraform attributes schema.
var UserstoreResourceIDAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.RegexMatches(
				regexp.MustCompile("(?i)^[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$"),
				"invalid UUIDv4 format",
			),
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreResourceIDTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreResourceIDTFModelToJSONClient(in *UserstoreResourceIDTFModel) (*UserstoreResourceIDJSONClientModel, error) {
	out := UserstoreResourceIDJSONClientModel{}
	var err error
	out.ID, err = func(val *types.String) (*uuid.UUID, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted, err := uuid.FromString(val.ValueString())
		if err != nil {
			return nil, ucerr.Errorf("failed to parse uuid: %v", err)
		}
		return &converted, nil
	}(&in.ID)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Name)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreResourceIDJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreResourceIDJSONClientModelToTF(in *UserstoreResourceIDJSONClientModel) (UserstoreResourceIDTFModel, error) {
	out := UserstoreResourceIDTFModel{}
	var err error
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return UserstoreResourceIDTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return UserstoreResourceIDTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateAccessorRequestTFModel is a Terraform model struct for the UserstoreUpdateAccessorRequestAttributes schema.
type UserstoreUpdateAccessorRequestTFModel struct {
	Accessor types.Object `tfsdk:"accessor"`
}

// UserstoreUpdateAccessorRequestJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateAccessorRequestJSONClientModel struct {
	Accessor *UserstoreAccessorJSONClientModel `json:"accessor,omitempty"`
}

// UserstoreUpdateAccessorRequestAttrTypes defines the attribute types for the UserstoreUpdateAccessorRequestAttributes schema.
var UserstoreUpdateAccessorRequestAttrTypes = map[string]attr.Type{
	"accessor": types.ObjectType{
		AttrTypes: UserstoreAccessorAttrTypes,
	},
}

// UserstoreUpdateAccessorRequestAttributes defines the Terraform attributes schema.
var UserstoreUpdateAccessorRequestAttributes = map[string]schema.Attribute{
	"accessor": schema.SingleNestedAttribute{
		Attributes:          UserstoreAccessorAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateAccessorRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateAccessorRequestTFModelToJSONClient(in *UserstoreUpdateAccessorRequestTFModel) (*UserstoreUpdateAccessorRequestJSONClientModel, error) {
	out := UserstoreUpdateAccessorRequestJSONClientModel{}
	var err error
	out.Accessor, err = func(val *types.Object) (*UserstoreAccessorJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreAccessorTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreAccessorTFModelToJSONClient(&tfModel)
	}(&in.Accessor)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"accessor\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateAccessorRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateAccessorRequestJSONClientModelToTF(in *UserstoreUpdateAccessorRequestJSONClientModel) (UserstoreUpdateAccessorRequestTFModel, error) {
	out := UserstoreUpdateAccessorRequestTFModel{}
	var err error
	out.Accessor, err = func(val *UserstoreAccessorJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreAccessorAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreAccessorJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreAccessorTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Accessor)
	if err != nil {
		return UserstoreUpdateAccessorRequestTFModel{}, ucerr.Errorf("failed to convert \"accessor\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRequestTFModel is a Terraform model struct for the UserstoreUpdateColumnRequestAttributes schema.
type UserstoreUpdateColumnRequestTFModel struct {
	Column types.Object `tfsdk:"column"`
}

// UserstoreUpdateColumnRequestJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRequestJSONClientModel struct {
	Column *UserstoreColumnJSONClientModel `json:"column,omitempty"`
}

// UserstoreUpdateColumnRequestAttrTypes defines the attribute types for the UserstoreUpdateColumnRequestAttributes schema.
var UserstoreUpdateColumnRequestAttrTypes = map[string]attr.Type{
	"column": types.ObjectType{
		AttrTypes: UserstoreColumnAttrTypes,
	},
}

// UserstoreUpdateColumnRequestAttributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRequestAttributes = map[string]schema.Attribute{
	"column": schema.SingleNestedAttribute{
		Attributes:          UserstoreColumnAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRequestTFModelToJSONClient(in *UserstoreUpdateColumnRequestTFModel) (*UserstoreUpdateColumnRequestJSONClientModel, error) {
	out := UserstoreUpdateColumnRequestJSONClientModel{}
	var err error
	out.Column, err = func(val *types.Object) (*UserstoreColumnJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreColumnTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreColumnTFModelToJSONClient(&tfModel)
	}(&in.Column)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRequestJSONClientModelToTF(in *UserstoreUpdateColumnRequestJSONClientModel) (UserstoreUpdateColumnRequestTFModel, error) {
	out := UserstoreUpdateColumnRequestTFModel{}
	var err error
	out.Column, err = func(val *UserstoreColumnJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreColumnAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreColumnJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreColumnTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Column)
	if err != nil {
		return UserstoreUpdateColumnRequestTFModel{}, ucerr.Errorf("failed to convert \"column\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestTFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestAttributes schema.
type UserstoreUpdateColumnRetentionDurationRequestTFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestJSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestAttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestAttributes schema.
var UserstoreUpdateColumnRetentionDurationRequestAttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestAttributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestAttributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestTFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestTFModel) (*UserstoreUpdateColumnRetentionDurationRequestJSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestJSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestJSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestJSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestTFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestTFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestTFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType2TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestType2Attributes schema.
type UserstoreUpdateColumnRetentionDurationRequestType2TFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestType2JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestType2JSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestType2AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestType2Attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType2AttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType2Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType2Attributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType2TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestType2TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestType2TFModel) (*UserstoreUpdateColumnRetentionDurationRequestType2JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType2JSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType2JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestType2JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestType2JSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestType2TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType2TFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestType2TFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType3TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestType3Attributes schema.
type UserstoreUpdateColumnRetentionDurationRequestType3TFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestType3JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestType3JSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestType3AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestType3Attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType3AttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType3Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType3Attributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType3TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestType3TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestType3TFModel) (*UserstoreUpdateColumnRetentionDurationRequestType3JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType3JSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType3JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestType3JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestType3JSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestType3TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType3TFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestType3TFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType4TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestType4Attributes schema.
type UserstoreUpdateColumnRetentionDurationRequestType4TFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestType4JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestType4JSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestType4AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestType4Attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType4AttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType4Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType4Attributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType4TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestType4TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestType4TFModel) (*UserstoreUpdateColumnRetentionDurationRequestType4JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType4JSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType4JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestType4JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestType4JSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestType4TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType4TFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestType4TFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType5TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestType5Attributes schema.
type UserstoreUpdateColumnRetentionDurationRequestType5TFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestType5JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestType5JSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestType5AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestType5Attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType5AttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType5Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType5Attributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType5TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestType5TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestType5TFModel) (*UserstoreUpdateColumnRetentionDurationRequestType5JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType5JSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType5JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestType5JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestType5JSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestType5TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType5TFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestType5TFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType6TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestType6Attributes schema.
type UserstoreUpdateColumnRetentionDurationRequestType6TFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestType6JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestType6JSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestType6AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestType6Attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType6AttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType6Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType6Attributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType6TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestType6TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestType6TFModel) (*UserstoreUpdateColumnRetentionDurationRequestType6JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType6JSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType6JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestType6JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestType6JSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestType6TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType6TFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestType6TFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType7TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestType7Attributes schema.
type UserstoreUpdateColumnRetentionDurationRequestType7TFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestType7JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestType7JSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestType7AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestType7Attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType7AttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType7Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType7Attributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType7TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestType7TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestType7TFModel) (*UserstoreUpdateColumnRetentionDurationRequestType7JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType7JSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType7JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestType7JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestType7JSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestType7TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType7TFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestType7TFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType8TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationRequestType8Attributes schema.
type UserstoreUpdateColumnRetentionDurationRequestType8TFModel struct {
	RetentionDuration types.Object `tfsdk:"retention_duration"`
}

// UserstoreUpdateColumnRetentionDurationRequestType8JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationRequestType8JSONClientModel struct {
	RetentionDuration *IdpColumnRetentionDurationJSONClientModel `json:"retention_duration,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationRequestType8AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationRequestType8Attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType8AttrTypes = map[string]attr.Type{
	"retention_duration": types.ObjectType{
		AttrTypes: IdpColumnRetentionDurationAttrTypes,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType8Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationRequestType8Attributes = map[string]schema.Attribute{
	"retention_duration": schema.SingleNestedAttribute{
		Attributes:          IdpColumnRetentionDurationAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationRequestType8TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationRequestType8TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationRequestType8TFModel) (*UserstoreUpdateColumnRetentionDurationRequestType8JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType8JSONClientModel{}
	var err error
	out.RetentionDuration, err = func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := IdpColumnRetentionDurationTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
	}(&in.RetentionDuration)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationRequestType8JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationRequestType8JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationRequestType8JSONClientModel) (UserstoreUpdateColumnRetentionDurationRequestType8TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationRequestType8TFModel{}
	var err error
	out.RetentionDuration, err = func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
		attrTypes := IdpColumnRetentionDurationAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.RetentionDuration)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationRequestType8TFModel{}, ucerr.Errorf("failed to convert \"retention_duration\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationsRequestTFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationsRequestAttributes schema.
type UserstoreUpdateColumnRetentionDurationsRequestTFModel struct {
	RetentionDurations types.List `tfsdk:"retention_durations"`
}

// UserstoreUpdateColumnRetentionDurationsRequestJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationsRequestJSONClientModel struct {
	RetentionDurations *[]IdpColumnRetentionDurationJSONClientModel `json:"retention_durations,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationsRequestAttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationsRequestAttributes schema.
var UserstoreUpdateColumnRetentionDurationsRequestAttrTypes = map[string]attr.Type{
	"retention_durations": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: IdpColumnRetentionDurationAttrTypes,
		},
	},
}

// UserstoreUpdateColumnRetentionDurationsRequestAttributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationsRequestAttributes = map[string]schema.Attribute{
	"retention_durations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IdpColumnRetentionDurationAttributes,
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationsRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationsRequestTFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationsRequestTFModel) (*UserstoreUpdateColumnRetentionDurationsRequestJSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationsRequestJSONClientModel{}
	var err error
	out.RetentionDurations, err = func(val *types.List) (*[]IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []IdpColumnRetentionDurationJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := IdpColumnRetentionDurationTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.RetentionDurations)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_durations\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationsRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationsRequestJSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationsRequestJSONClientModel) (UserstoreUpdateColumnRetentionDurationsRequestTFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationsRequestTFModel{}
	var err error
	out.RetentionDurations, err = func(val *[]IdpColumnRetentionDurationJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: IdpColumnRetentionDurationAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
				attrTypes := IdpColumnRetentionDurationAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.RetentionDurations)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationsRequestTFModel{}, ucerr.Errorf("failed to convert \"retention_durations\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateColumnRetentionDurationsRequestType2TFModel is a Terraform model struct for the UserstoreUpdateColumnRetentionDurationsRequestType2Attributes schema.
type UserstoreUpdateColumnRetentionDurationsRequestType2TFModel struct {
	RetentionDurations types.List `tfsdk:"retention_durations"`
}

// UserstoreUpdateColumnRetentionDurationsRequestType2JSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateColumnRetentionDurationsRequestType2JSONClientModel struct {
	RetentionDurations *[]IdpColumnRetentionDurationJSONClientModel `json:"retention_durations,omitempty"`
}

// UserstoreUpdateColumnRetentionDurationsRequestType2AttrTypes defines the attribute types for the UserstoreUpdateColumnRetentionDurationsRequestType2Attributes schema.
var UserstoreUpdateColumnRetentionDurationsRequestType2AttrTypes = map[string]attr.Type{
	"retention_durations": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: IdpColumnRetentionDurationAttrTypes,
		},
	},
}

// UserstoreUpdateColumnRetentionDurationsRequestType2Attributes defines the Terraform attributes schema.
var UserstoreUpdateColumnRetentionDurationsRequestType2Attributes = map[string]schema.Attribute{
	"retention_durations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IdpColumnRetentionDurationAttributes,
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateColumnRetentionDurationsRequestType2TFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateColumnRetentionDurationsRequestType2TFModelToJSONClient(in *UserstoreUpdateColumnRetentionDurationsRequestType2TFModel) (*UserstoreUpdateColumnRetentionDurationsRequestType2JSONClientModel, error) {
	out := UserstoreUpdateColumnRetentionDurationsRequestType2JSONClientModel{}
	var err error
	out.RetentionDurations, err = func(val *types.List) (*[]IdpColumnRetentionDurationJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []IdpColumnRetentionDurationJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*IdpColumnRetentionDurationJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := IdpColumnRetentionDurationTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return IdpColumnRetentionDurationTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.RetentionDurations)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"retention_durations\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateColumnRetentionDurationsRequestType2JSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateColumnRetentionDurationsRequestType2JSONClientModelToTF(in *UserstoreUpdateColumnRetentionDurationsRequestType2JSONClientModel) (UserstoreUpdateColumnRetentionDurationsRequestType2TFModel, error) {
	out := UserstoreUpdateColumnRetentionDurationsRequestType2TFModel{}
	var err error
	out.RetentionDurations, err = func(val *[]IdpColumnRetentionDurationJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: IdpColumnRetentionDurationAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *IdpColumnRetentionDurationJSONClientModel) (types.Object, error) {
				attrTypes := IdpColumnRetentionDurationAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := IdpColumnRetentionDurationJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert IdpColumnRetentionDurationTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.RetentionDurations)
	if err != nil {
		return UserstoreUpdateColumnRetentionDurationsRequestType2TFModel{}, ucerr.Errorf("failed to convert \"retention_durations\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdateMutatorRequestTFModel is a Terraform model struct for the UserstoreUpdateMutatorRequestAttributes schema.
type UserstoreUpdateMutatorRequestTFModel struct {
	Mutator types.Object `tfsdk:"mutator"`
}

// UserstoreUpdateMutatorRequestJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdateMutatorRequestJSONClientModel struct {
	Mutator *UserstoreMutatorJSONClientModel `json:"mutator,omitempty"`
}

// UserstoreUpdateMutatorRequestAttrTypes defines the attribute types for the UserstoreUpdateMutatorRequestAttributes schema.
var UserstoreUpdateMutatorRequestAttrTypes = map[string]attr.Type{
	"mutator": types.ObjectType{
		AttrTypes: UserstoreMutatorAttrTypes,
	},
}

// UserstoreUpdateMutatorRequestAttributes defines the Terraform attributes schema.
var UserstoreUpdateMutatorRequestAttributes = map[string]schema.Attribute{
	"mutator": schema.SingleNestedAttribute{
		Attributes:          UserstoreMutatorAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdateMutatorRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdateMutatorRequestTFModelToJSONClient(in *UserstoreUpdateMutatorRequestTFModel) (*UserstoreUpdateMutatorRequestJSONClientModel, error) {
	out := UserstoreUpdateMutatorRequestJSONClientModel{}
	var err error
	out.Mutator, err = func(val *types.Object) (*UserstoreMutatorJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreMutatorTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreMutatorTFModelToJSONClient(&tfModel)
	}(&in.Mutator)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"mutator\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdateMutatorRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdateMutatorRequestJSONClientModelToTF(in *UserstoreUpdateMutatorRequestJSONClientModel) (UserstoreUpdateMutatorRequestTFModel, error) {
	out := UserstoreUpdateMutatorRequestTFModel{}
	var err error
	out.Mutator, err = func(val *UserstoreMutatorJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreMutatorAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreMutatorJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreMutatorTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Mutator)
	if err != nil {
		return UserstoreUpdateMutatorRequestTFModel{}, ucerr.Errorf("failed to convert \"mutator\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUpdatePurposeRequestTFModel is a Terraform model struct for the UserstoreUpdatePurposeRequestAttributes schema.
type UserstoreUpdatePurposeRequestTFModel struct {
	Purpose types.Object `tfsdk:"purpose"`
}

// UserstoreUpdatePurposeRequestJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUpdatePurposeRequestJSONClientModel struct {
	Purpose *UserstorePurposeJSONClientModel `json:"purpose,omitempty"`
}

// UserstoreUpdatePurposeRequestAttrTypes defines the attribute types for the UserstoreUpdatePurposeRequestAttributes schema.
var UserstoreUpdatePurposeRequestAttrTypes = map[string]attr.Type{
	"purpose": types.ObjectType{
		AttrTypes: UserstorePurposeAttrTypes,
	},
}

// UserstoreUpdatePurposeRequestAttributes defines the Terraform attributes schema.
var UserstoreUpdatePurposeRequestAttributes = map[string]schema.Attribute{
	"purpose": schema.SingleNestedAttribute{
		Attributes:          UserstorePurposeAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUpdatePurposeRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUpdatePurposeRequestTFModelToJSONClient(in *UserstoreUpdatePurposeRequestTFModel) (*UserstoreUpdatePurposeRequestJSONClientModel, error) {
	out := UserstoreUpdatePurposeRequestJSONClientModel{}
	var err error
	out.Purpose, err = func(val *types.Object) (*UserstorePurposeJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstorePurposeTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstorePurposeTFModelToJSONClient(&tfModel)
	}(&in.Purpose)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"purpose\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUpdatePurposeRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUpdatePurposeRequestJSONClientModelToTF(in *UserstoreUpdatePurposeRequestJSONClientModel) (UserstoreUpdatePurposeRequestTFModel, error) {
	out := UserstoreUpdatePurposeRequestTFModel{}
	var err error
	out.Purpose, err = func(val *UserstorePurposeJSONClientModel) (types.Object, error) {
		attrTypes := UserstorePurposeAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstorePurposeJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstorePurposeTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Purpose)
	if err != nil {
		return UserstoreUpdatePurposeRequestTFModel{}, ucerr.Errorf("failed to convert \"purpose\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUserSelectorConfigTFModel is a Terraform model struct for the UserstoreUserSelectorConfigAttributes schema.
type UserstoreUserSelectorConfigTFModel struct {
	WhereClause types.String `tfsdk:"where_clause"`
}

// UserstoreUserSelectorConfigJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUserSelectorConfigJSONClientModel struct {
	WhereClause *string `json:"where_clause,omitempty"`
}

// UserstoreUserSelectorConfigAttrTypes defines the attribute types for the UserstoreUserSelectorConfigAttributes schema.
var UserstoreUserSelectorConfigAttrTypes = map[string]attr.Type{
	"where_clause": types.StringType,
}

// UserstoreUserSelectorConfigAttributes defines the Terraform attributes schema.
var UserstoreUserSelectorConfigAttributes = map[string]schema.Attribute{
	"where_clause": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// UserstoreUserSelectorConfigTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUserSelectorConfigTFModelToJSONClient(in *UserstoreUserSelectorConfigTFModel) (*UserstoreUserSelectorConfigJSONClientModel, error) {
	out := UserstoreUserSelectorConfigJSONClientModel{}
	var err error
	out.WhereClause, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.WhereClause)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"where_clause\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreUserSelectorConfigJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUserSelectorConfigJSONClientModelToTF(in *UserstoreUserSelectorConfigJSONClientModel) (UserstoreUserSelectorConfigTFModel, error) {
	out := UserstoreUserSelectorConfigTFModel{}
	var err error
	out.WhereClause, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.WhereClause)
	if err != nil {
		return UserstoreUserSelectorConfigTFModel{}, ucerr.Errorf("failed to convert \"where_clause\" field: %+v", err)
	}
	return out, nil
}

// UserstoreUserSelectorValuesTFModel is a Terraform model struct for the UserstoreUserSelectorValuesAttributes schema.
type UserstoreUserSelectorValuesTFModel struct {
}

// UserstoreUserSelectorValuesJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreUserSelectorValuesJSONClientModel struct {
}

// UserstoreUserSelectorValuesAttrTypes defines the attribute types for the UserstoreUserSelectorValuesAttributes schema.
var UserstoreUserSelectorValuesAttrTypes = map[string]attr.Type{}

// UserstoreUserSelectorValuesAttributes defines the Terraform attributes schema.
var UserstoreUserSelectorValuesAttributes = map[string]schema.Attribute{}

// UserstoreUserSelectorValuesTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreUserSelectorValuesTFModelToJSONClient(in *UserstoreUserSelectorValuesTFModel) (*UserstoreUserSelectorValuesJSONClientModel, error) {
	out := UserstoreUserSelectorValuesJSONClientModel{}
	return &out, nil
}

// UserstoreUserSelectorValuesJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreUserSelectorValuesJSONClientModelToTF(in *UserstoreUserSelectorValuesJSONClientModel) (UserstoreUserSelectorValuesTFModel, error) {
	out := UserstoreUserSelectorValuesTFModel{}
	return out, nil
}

// UUIDUuidTFModel is a Terraform model struct for the UUIDUuidAttributes schema.
type UUIDUuidTFModel struct {
}

// UUIDUuidJSONClientModel stores data for use with jsonclient for making API requests.
type UUIDUuidJSONClientModel struct {
}

// UUIDUuidAttrTypes defines the attribute types for the UUIDUuidAttributes schema.
var UUIDUuidAttrTypes = map[string]attr.Type{}

// UUIDUuidAttributes defines the Terraform attributes schema.
var UUIDUuidAttributes = map[string]schema.Attribute{}

// UUIDUuidTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UUIDUuidTFModelToJSONClient(in *UUIDUuidTFModel) (*UUIDUuidJSONClientModel, error) {
	out := UUIDUuidJSONClientModel{}
	return &out, nil
}

// UUIDUuidJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UUIDUuidJSONClientModelToTF(in *UUIDUuidJSONClientModel) (UUIDUuidTFModel, error) {
	out := UUIDUuidTFModel{}
	return out, nil
}

// NOTE: automatically generated file -- DO NOT EDIT

package tokenizer

import (
	"context"
	"reflect"
	"regexp"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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

// boolplanmodifier is used in userstore schemas but not tokenizer
var _ = boolplanmodifier.RequiresReplace

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

// PolicyAccessPolicyTFModel is a Terraform model struct for the PolicyAccessPolicyAttributes schema.
type PolicyAccessPolicyTFModel struct {
	Components  types.List   `tfsdk:"components"`
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	PolicyType  types.String `tfsdk:"policy_type"`
	TagIds      types.List   `tfsdk:"tag_ids"`
	Version     types.Int64  `tfsdk:"version"`
}

// PolicyAccessPolicyJSONClientModel stores data for use with jsonclient for making API requests.
type PolicyAccessPolicyJSONClientModel struct {
	Components  *[]PolicyAccessPolicyComponentJSONClientModel `json:"components,omitempty"`
	Description *string                                       `json:"description,omitempty"`
	ID          *uuid.UUID                                    `json:"id,omitempty"`
	Name        *string                                       `json:"name,omitempty"`
	PolicyType  *string                                       `json:"policy_type,omitempty"`
	TagIds      *[]uuid.UUID                                  `json:"tag_ids,omitempty"`
	Version     *int64                                        `json:"version,omitempty"`
}

// PolicyAccessPolicyAttrTypes defines the attribute types for the PolicyAccessPolicyAttributes schema.
var PolicyAccessPolicyAttrTypes = map[string]attr.Type{
	"components": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: PolicyAccessPolicyComponentAttrTypes,
		},
	},
	"description": types.StringType,
	"id":          types.StringType,
	"name":        types.StringType,
	"policy_type": types.StringType,
	"tag_ids": types.ListType{
		ElemType: types.StringType,
	},
	"version": types.Int64Type,
}

// PolicyAccessPolicyAttributes defines the Terraform attributes schema.
var PolicyAccessPolicyAttributes = map[string]schema.Attribute{
	"components": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: PolicyAccessPolicyComponentAttributes,
		},
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
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
	"policy_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"composite_and", "compositeintersection", "composite_or", "compositeunion"}...),
		},
		Description:         "Valid values: `composite_and`, `compositeintersection`, `composite_or`, `compositeunion`",
		MarkdownDescription: "Valid values: `composite_and`, `compositeintersection`, `composite_or`, `compositeunion`",
		Required:            true,
	},
	"tag_ids": schema.ListAttribute{
		ElementType:         types.StringType,
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

// PolicyAccessPolicyTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PolicyAccessPolicyTFModelToJSONClient(in *PolicyAccessPolicyTFModel) (*PolicyAccessPolicyJSONClientModel, error) {
	out := PolicyAccessPolicyJSONClientModel{}
	var err error
	out.Components, err = func(val *types.List) (*[]PolicyAccessPolicyComponentJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []PolicyAccessPolicyComponentJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*PolicyAccessPolicyComponentJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := PolicyAccessPolicyComponentTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return PolicyAccessPolicyComponentTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Components)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"components\" field: %+v", err)
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
	out.PolicyType, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.PolicyType)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"policy_type\" field: %+v", err)
	}
	out.TagIds, err = func(val *types.List) (*[]uuid.UUID, error) {
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
	}(&in.TagIds)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"tag_ids\" field: %+v", err)
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

// PolicyAccessPolicyJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PolicyAccessPolicyJSONClientModelToTF(in *PolicyAccessPolicyJSONClientModel) (PolicyAccessPolicyTFModel, error) {
	out := PolicyAccessPolicyTFModel{}
	var err error
	out.Components, err = func(val *[]PolicyAccessPolicyComponentJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: PolicyAccessPolicyComponentAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *PolicyAccessPolicyComponentJSONClientModel) (types.Object, error) {
				attrTypes := PolicyAccessPolicyComponentAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := PolicyAccessPolicyComponentJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyAccessPolicyComponentTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Components)
	if err != nil {
		return PolicyAccessPolicyTFModel{}, ucerr.Errorf("failed to convert \"components\" field: %+v", err)
	}
	out.Description, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Description)
	if err != nil {
		return PolicyAccessPolicyTFModel{}, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return PolicyAccessPolicyTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return PolicyAccessPolicyTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.PolicyType, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.PolicyType)
	if err != nil {
		return PolicyAccessPolicyTFModel{}, ucerr.Errorf("failed to convert \"policy_type\" field: %+v", err)
	}
	out.TagIds, err = func(val *[]uuid.UUID) (types.List, error) {
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
	}(in.TagIds)
	if err != nil {
		return PolicyAccessPolicyTFModel{}, ucerr.Errorf("failed to convert \"tag_ids\" field: %+v", err)
	}
	out.Version, err = func(val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}(in.Version)
	if err != nil {
		return PolicyAccessPolicyTFModel{}, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
	return out, nil
}

// PolicyAccessPolicyComponentTFModel is a Terraform model struct for the PolicyAccessPolicyComponentAttributes schema.
type PolicyAccessPolicyComponentTFModel struct {
	Policy             types.String `tfsdk:"policy"`
	Template           types.String `tfsdk:"template"`
	TemplateParameters types.String `tfsdk:"template_parameters"`
}

// PolicyAccessPolicyComponentJSONClientModel stores data for use with jsonclient for making API requests.
type PolicyAccessPolicyComponentJSONClientModel struct {
	Policy             *UserstoreResourceIDJSONClientModel `json:"policy,omitempty"`
	Template           *UserstoreResourceIDJSONClientModel `json:"template,omitempty"`
	TemplateParameters *string                             `json:"template_parameters,omitempty"`
}

// PolicyAccessPolicyComponentAttrTypes defines the attribute types for the PolicyAccessPolicyComponentAttributes schema.
var PolicyAccessPolicyComponentAttrTypes = map[string]attr.Type{
	"policy":              types.StringType,
	"template":            types.StringType,
	"template_parameters": types.StringType,
}

// PolicyAccessPolicyComponentAttributes defines the Terraform attributes schema.
var PolicyAccessPolicyComponentAttributes = map[string]schema.Attribute{
	"policy": schema.StringAttribute{
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
	"template": schema.StringAttribute{
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
	"template_parameters": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// PolicyAccessPolicyComponentTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PolicyAccessPolicyComponentTFModelToJSONClient(in *PolicyAccessPolicyComponentTFModel) (*PolicyAccessPolicyComponentJSONClientModel, error) {
	out := PolicyAccessPolicyComponentJSONClientModel{}
	var err error
	out.Policy, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.Policy)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"policy\" field: %+v", err)
	}
	out.Template, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.Template)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"template\" field: %+v", err)
	}
	out.TemplateParameters, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.TemplateParameters)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"template_parameters\" field: %+v", err)
	}
	return &out, nil
}

// PolicyAccessPolicyComponentJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PolicyAccessPolicyComponentJSONClientModelToTF(in *PolicyAccessPolicyComponentJSONClientModel) (PolicyAccessPolicyComponentTFModel, error) {
	out := PolicyAccessPolicyComponentTFModel{}
	var err error
	out.Policy, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
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
	}(in.Policy)
	if err != nil {
		return PolicyAccessPolicyComponentTFModel{}, ucerr.Errorf("failed to convert \"policy\" field: %+v", err)
	}
	out.Template, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
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
	}(in.Template)
	if err != nil {
		return PolicyAccessPolicyComponentTFModel{}, ucerr.Errorf("failed to convert \"template\" field: %+v", err)
	}
	out.TemplateParameters, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.TemplateParameters)
	if err != nil {
		return PolicyAccessPolicyComponentTFModel{}, ucerr.Errorf("failed to convert \"template_parameters\" field: %+v", err)
	}
	return out, nil
}

// PolicyAccessPolicyTemplateTFModel is a Terraform model struct for the PolicyAccessPolicyTemplateAttributes schema.
type PolicyAccessPolicyTemplateTFModel struct {
	Created     types.String `tfsdk:"created"`
	Deleted     types.String `tfsdk:"deleted"`
	Description types.String `tfsdk:"description"`
	Function    types.String `tfsdk:"function"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Updated     types.String `tfsdk:"updated"`
	Version     types.Int64  `tfsdk:"version"`
}

// PolicyAccessPolicyTemplateJSONClientModel stores data for use with jsonclient for making API requests.
type PolicyAccessPolicyTemplateJSONClientModel struct {
	Created     *string    `json:"created,omitempty"`
	Deleted     *string    `json:"deleted,omitempty"`
	Description *string    `json:"description,omitempty"`
	Function    *string    `json:"function,omitempty"`
	ID          *uuid.UUID `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Updated     *string    `json:"updated,omitempty"`
	Version     *int64     `json:"version,omitempty"`
}

// PolicyAccessPolicyTemplateAttrTypes defines the attribute types for the PolicyAccessPolicyTemplateAttributes schema.
var PolicyAccessPolicyTemplateAttrTypes = map[string]attr.Type{
	"created":     types.StringType,
	"deleted":     types.StringType,
	"description": types.StringType,
	"function":    types.StringType,
	"id":          types.StringType,
	"name":        types.StringType,
	"updated":     types.StringType,
	"version":     types.Int64Type,
}

// PolicyAccessPolicyTemplateAttributes defines the Terraform attributes schema.
var PolicyAccessPolicyTemplateAttributes = map[string]schema.Attribute{
	"created": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"deleted": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"description": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"function": schema.StringAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
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
	"updated": schema.StringAttribute{
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

// PolicyAccessPolicyTemplateTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PolicyAccessPolicyTemplateTFModelToJSONClient(in *PolicyAccessPolicyTemplateTFModel) (*PolicyAccessPolicyTemplateJSONClientModel, error) {
	out := PolicyAccessPolicyTemplateJSONClientModel{}
	var err error
	out.Created, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Created)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"created\" field: %+v", err)
	}
	out.Deleted, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Deleted)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"deleted\" field: %+v", err)
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
	out.Function, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Function)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"function\" field: %+v", err)
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
	out.Updated, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Updated)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"updated\" field: %+v", err)
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

// PolicyAccessPolicyTemplateJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PolicyAccessPolicyTemplateJSONClientModelToTF(in *PolicyAccessPolicyTemplateJSONClientModel) (PolicyAccessPolicyTemplateTFModel, error) {
	out := PolicyAccessPolicyTemplateTFModel{}
	var err error
	out.Created, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Created)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"created\" field: %+v", err)
	}
	out.Deleted, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Deleted)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"deleted\" field: %+v", err)
	}
	out.Description, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Description)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.Function, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Function)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"function\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.Updated, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Updated)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"updated\" field: %+v", err)
	}
	out.Version, err = func(val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}(in.Version)
	if err != nil {
		return PolicyAccessPolicyTemplateTFModel{}, ucerr.Errorf("failed to convert \"version\" field: %+v", err)
	}
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

// PolicyPolicyTypeTFModel is a Terraform model struct for the PolicyPolicyTypeAttributes schema.
type PolicyPolicyTypeTFModel struct {
}

// PolicyPolicyTypeJSONClientModel stores data for use with jsonclient for making API requests.
type PolicyPolicyTypeJSONClientModel struct {
}

// PolicyPolicyTypeAttrTypes defines the attribute types for the PolicyPolicyTypeAttributes schema.
var PolicyPolicyTypeAttrTypes = map[string]attr.Type{}

// PolicyPolicyTypeAttributes defines the Terraform attributes schema.
var PolicyPolicyTypeAttributes = map[string]schema.Attribute{}

// PolicyPolicyTypeTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PolicyPolicyTypeTFModelToJSONClient(in *PolicyPolicyTypeTFModel) (*PolicyPolicyTypeJSONClientModel, error) {
	out := PolicyPolicyTypeJSONClientModel{}
	return &out, nil
}

// PolicyPolicyTypeJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PolicyPolicyTypeJSONClientModelToTF(in *PolicyPolicyTypeJSONClientModel) (PolicyPolicyTypeTFModel, error) {
	out := PolicyPolicyTypeTFModel{}
	return out, nil
}

// PolicyTransformTypeTFModel is a Terraform model struct for the PolicyTransformTypeAttributes schema.
type PolicyTransformTypeTFModel struct {
}

// PolicyTransformTypeJSONClientModel stores data for use with jsonclient for making API requests.
type PolicyTransformTypeJSONClientModel struct {
}

// PolicyTransformTypeAttrTypes defines the attribute types for the PolicyTransformTypeAttributes schema.
var PolicyTransformTypeAttrTypes = map[string]attr.Type{}

// PolicyTransformTypeAttributes defines the Terraform attributes schema.
var PolicyTransformTypeAttributes = map[string]schema.Attribute{}

// PolicyTransformTypeTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PolicyTransformTypeTFModelToJSONClient(in *PolicyTransformTypeTFModel) (*PolicyTransformTypeJSONClientModel, error) {
	out := PolicyTransformTypeJSONClientModel{}
	return &out, nil
}

// PolicyTransformTypeJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PolicyTransformTypeJSONClientModelToTF(in *PolicyTransformTypeJSONClientModel) (PolicyTransformTypeTFModel, error) {
	out := PolicyTransformTypeTFModel{}
	return out, nil
}

// PolicyTransformerTFModel is a Terraform model struct for the PolicyTransformerAttributes schema.
type PolicyTransformerTFModel struct {
	Description           types.String `tfsdk:"description"`
	Function              types.String `tfsdk:"function"`
	ID                    types.String `tfsdk:"id"`
	InputType             types.String `tfsdk:"input_type"`
	InputTypeConstraints  types.Object `tfsdk:"input_type_constraints"`
	Name                  types.String `tfsdk:"name"`
	OutputType            types.String `tfsdk:"output_type"`
	OutputTypeConstraints types.Object `tfsdk:"output_type_constraints"`
	Parameters            types.String `tfsdk:"parameters"`
	ReuseExistingToken    types.Bool   `tfsdk:"reuse_existing_token"`
	TagIds                types.List   `tfsdk:"tag_ids"`
	TransformType         types.String `tfsdk:"transform_type"`
}

// PolicyTransformerJSONClientModel stores data for use with jsonclient for making API requests.
type PolicyTransformerJSONClientModel struct {
	Description           *string                                    `json:"description,omitempty"`
	Function              *string                                    `json:"function,omitempty"`
	ID                    *uuid.UUID                                 `json:"id,omitempty"`
	InputType             *string                                    `json:"input_type,omitempty"`
	InputTypeConstraints  *UserstoreColumnConstraintsJSONClientModel `json:"input_type_constraints,omitempty"`
	Name                  *string                                    `json:"name,omitempty"`
	OutputType            *string                                    `json:"output_type,omitempty"`
	OutputTypeConstraints *UserstoreColumnConstraintsJSONClientModel `json:"output_type_constraints,omitempty"`
	Parameters            *string                                    `json:"parameters,omitempty"`
	ReuseExistingToken    *bool                                      `json:"reuse_existing_token,omitempty"`
	TagIds                *[]uuid.UUID                               `json:"tag_ids,omitempty"`
	TransformType         *string                                    `json:"transform_type,omitempty"`
}

// PolicyTransformerAttrTypes defines the attribute types for the PolicyTransformerAttributes schema.
var PolicyTransformerAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"function":    types.StringType,
	"id":          types.StringType,
	"input_type":  types.StringType,
	"input_type_constraints": types.ObjectType{
		AttrTypes: UserstoreColumnConstraintsAttrTypes,
	},
	"name":        types.StringType,
	"output_type": types.StringType,
	"output_type_constraints": types.ObjectType{
		AttrTypes: UserstoreColumnConstraintsAttrTypes,
	},
	"parameters":           types.StringType,
	"reuse_existing_token": types.BoolType,
	"tag_ids": types.ListType{
		ElemType: types.StringType,
	},
	"transform_type": types.StringType,
}

// PolicyTransformerAttributes defines the Terraform attributes schema.
var PolicyTransformerAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"function": schema.StringAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
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
	"input_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"address", "birthdate", "boolean", "composite", "date", "e164_phonenumber", "email", "integer", "phonenumber", "ssn", "string", "timestamp", "uuid"}...),
		},
		Description:         "Valid values: `address`, `birthdate`, `boolean`, `composite`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		MarkdownDescription: "Valid values: `address`, `birthdate`, `boolean`, `composite`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		Required:            true,
	},
	"input_type_constraints": schema.SingleNestedAttribute{
		Attributes:          UserstoreColumnConstraintsAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"name": schema.StringAttribute{
		Description:         "",
		MarkdownDescription: "",
		Required:            true,
	},
	"output_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"address", "birthdate", "boolean", "composite", "date", "e164_phonenumber", "email", "integer", "phonenumber", "ssn", "string", "timestamp", "uuid"}...),
		},
		Computed:            true,
		Description:         "Valid values: `address`, `birthdate`, `boolean`, `composite`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		MarkdownDescription: "Valid values: `address`, `birthdate`, `boolean`, `composite`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		Optional:            true,
	},
	"output_type_constraints": schema.SingleNestedAttribute{
		Attributes:          UserstoreColumnConstraintsAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"parameters": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"reuse_existing_token": schema.BoolAttribute{
		Computed:            true,
		Description:         "Specifies if the tokenizing transformer should return existing token instead of creating a new one.",
		MarkdownDescription: "Specifies if the tokenizing transformer should return existing token instead of creating a new one.",
		Optional:            true,
	},
	"tag_ids": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"transform_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"passthrough", "tokenizebyreference", "tokenizebyvalue", "transform"}...),
		},
		Description:         "Valid values: `passthrough`, `tokenizebyreference`, `tokenizebyvalue`, `transform`",
		MarkdownDescription: "Valid values: `passthrough`, `tokenizebyreference`, `tokenizebyvalue`, `transform`",
		Required:            true,
	},
}

// PolicyTransformerTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func PolicyTransformerTFModelToJSONClient(in *PolicyTransformerTFModel) (*PolicyTransformerJSONClientModel, error) {
	out := PolicyTransformerJSONClientModel{}
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
	out.Function, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Function)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"function\" field: %+v", err)
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
	out.InputType, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.InputType)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"input_type\" field: %+v", err)
	}
	out.InputTypeConstraints, err = func(val *types.Object) (*UserstoreColumnConstraintsJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreColumnConstraintsTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreColumnConstraintsTFModelToJSONClient(&tfModel)
	}(&in.InputTypeConstraints)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"input_type_constraints\" field: %+v", err)
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
	out.OutputType, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.OutputType)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"output_type\" field: %+v", err)
	}
	out.OutputTypeConstraints, err = func(val *types.Object) (*UserstoreColumnConstraintsJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := UserstoreColumnConstraintsTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return UserstoreColumnConstraintsTFModelToJSONClient(&tfModel)
	}(&in.OutputTypeConstraints)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"output_type_constraints\" field: %+v", err)
	}
	out.Parameters, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Parameters)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"parameters\" field: %+v", err)
	}
	out.ReuseExistingToken, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.ReuseExistingToken)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"reuse_existing_token\" field: %+v", err)
	}
	out.TagIds, err = func(val *types.List) (*[]uuid.UUID, error) {
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
	}(&in.TagIds)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"tag_ids\" field: %+v", err)
	}
	out.TransformType, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.TransformType)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"transform_type\" field: %+v", err)
	}
	return &out, nil
}

// PolicyTransformerJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func PolicyTransformerJSONClientModelToTF(in *PolicyTransformerJSONClientModel) (PolicyTransformerTFModel, error) {
	out := PolicyTransformerTFModel{}
	var err error
	out.Description, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Description)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"description\" field: %+v", err)
	}
	out.Function, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Function)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"function\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.InputType, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.InputType)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"input_type\" field: %+v", err)
	}
	out.InputTypeConstraints, err = func(val *UserstoreColumnConstraintsJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreColumnConstraintsAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreColumnConstraintsJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreColumnConstraintsTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.InputTypeConstraints)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"input_type_constraints\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.OutputType, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.OutputType)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"output_type\" field: %+v", err)
	}
	out.OutputTypeConstraints, err = func(val *UserstoreColumnConstraintsJSONClientModel) (types.Object, error) {
		attrTypes := UserstoreColumnConstraintsAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := UserstoreColumnConstraintsJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreColumnConstraintsTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.OutputTypeConstraints)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"output_type_constraints\" field: %+v", err)
	}
	out.Parameters, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Parameters)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"parameters\" field: %+v", err)
	}
	out.ReuseExistingToken, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.ReuseExistingToken)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"reuse_existing_token\" field: %+v", err)
	}
	out.TagIds, err = func(val *[]uuid.UUID) (types.List, error) {
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
	}(in.TagIds)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"tag_ids\" field: %+v", err)
	}
	out.TransformType, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.TransformType)
	if err != nil {
		return PolicyTransformerTFModel{}, ucerr.Errorf("failed to convert \"transform_type\" field: %+v", err)
	}
	return out, nil
}

// TokenizerCreateAccessPolicyRequestTFModel is a Terraform model struct for the TokenizerCreateAccessPolicyRequestAttributes schema.
type TokenizerCreateAccessPolicyRequestTFModel struct {
	AccessPolicy types.Object `tfsdk:"access_policy"`
}

// TokenizerCreateAccessPolicyRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerCreateAccessPolicyRequestJSONClientModel struct {
	AccessPolicy *PolicyAccessPolicyJSONClientModel `json:"access_policy,omitempty"`
}

// TokenizerCreateAccessPolicyRequestAttrTypes defines the attribute types for the TokenizerCreateAccessPolicyRequestAttributes schema.
var TokenizerCreateAccessPolicyRequestAttrTypes = map[string]attr.Type{
	"access_policy": types.ObjectType{
		AttrTypes: PolicyAccessPolicyAttrTypes,
	},
}

// TokenizerCreateAccessPolicyRequestAttributes defines the Terraform attributes schema.
var TokenizerCreateAccessPolicyRequestAttributes = map[string]schema.Attribute{
	"access_policy": schema.SingleNestedAttribute{
		Attributes:          PolicyAccessPolicyAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerCreateAccessPolicyRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerCreateAccessPolicyRequestTFModelToJSONClient(in *TokenizerCreateAccessPolicyRequestTFModel) (*TokenizerCreateAccessPolicyRequestJSONClientModel, error) {
	out := TokenizerCreateAccessPolicyRequestJSONClientModel{}
	var err error
	out.AccessPolicy, err = func(val *types.Object) (*PolicyAccessPolicyJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := PolicyAccessPolicyTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return PolicyAccessPolicyTFModelToJSONClient(&tfModel)
	}(&in.AccessPolicy)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerCreateAccessPolicyRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerCreateAccessPolicyRequestJSONClientModelToTF(in *TokenizerCreateAccessPolicyRequestJSONClientModel) (TokenizerCreateAccessPolicyRequestTFModel, error) {
	out := TokenizerCreateAccessPolicyRequestTFModel{}
	var err error
	out.AccessPolicy, err = func(val *PolicyAccessPolicyJSONClientModel) (types.Object, error) {
		attrTypes := PolicyAccessPolicyAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := PolicyAccessPolicyJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyAccessPolicyTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.AccessPolicy)
	if err != nil {
		return TokenizerCreateAccessPolicyRequestTFModel{}, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	return out, nil
}

// TokenizerCreateAccessPolicyTemplateRequestTFModel is a Terraform model struct for the TokenizerCreateAccessPolicyTemplateRequestAttributes schema.
type TokenizerCreateAccessPolicyTemplateRequestTFModel struct {
	AccessPolicyTemplate types.Object `tfsdk:"access_policy_template"`
}

// TokenizerCreateAccessPolicyTemplateRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerCreateAccessPolicyTemplateRequestJSONClientModel struct {
	AccessPolicyTemplate *PolicyAccessPolicyTemplateJSONClientModel `json:"access_policy_template,omitempty"`
}

// TokenizerCreateAccessPolicyTemplateRequestAttrTypes defines the attribute types for the TokenizerCreateAccessPolicyTemplateRequestAttributes schema.
var TokenizerCreateAccessPolicyTemplateRequestAttrTypes = map[string]attr.Type{
	"access_policy_template": types.ObjectType{
		AttrTypes: PolicyAccessPolicyTemplateAttrTypes,
	},
}

// TokenizerCreateAccessPolicyTemplateRequestAttributes defines the Terraform attributes schema.
var TokenizerCreateAccessPolicyTemplateRequestAttributes = map[string]schema.Attribute{
	"access_policy_template": schema.SingleNestedAttribute{
		Attributes:          PolicyAccessPolicyTemplateAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerCreateAccessPolicyTemplateRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerCreateAccessPolicyTemplateRequestTFModelToJSONClient(in *TokenizerCreateAccessPolicyTemplateRequestTFModel) (*TokenizerCreateAccessPolicyTemplateRequestJSONClientModel, error) {
	out := TokenizerCreateAccessPolicyTemplateRequestJSONClientModel{}
	var err error
	out.AccessPolicyTemplate, err = func(val *types.Object) (*PolicyAccessPolicyTemplateJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := PolicyAccessPolicyTemplateTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return PolicyAccessPolicyTemplateTFModelToJSONClient(&tfModel)
	}(&in.AccessPolicyTemplate)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy_template\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerCreateAccessPolicyTemplateRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerCreateAccessPolicyTemplateRequestJSONClientModelToTF(in *TokenizerCreateAccessPolicyTemplateRequestJSONClientModel) (TokenizerCreateAccessPolicyTemplateRequestTFModel, error) {
	out := TokenizerCreateAccessPolicyTemplateRequestTFModel{}
	var err error
	out.AccessPolicyTemplate, err = func(val *PolicyAccessPolicyTemplateJSONClientModel) (types.Object, error) {
		attrTypes := PolicyAccessPolicyTemplateAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := PolicyAccessPolicyTemplateJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyAccessPolicyTemplateTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.AccessPolicyTemplate)
	if err != nil {
		return TokenizerCreateAccessPolicyTemplateRequestTFModel{}, ucerr.Errorf("failed to convert \"access_policy_template\" field: %+v", err)
	}
	return out, nil
}

// TokenizerCreateTokenRequestTFModel is a Terraform model struct for the TokenizerCreateTokenRequestAttributes schema.
type TokenizerCreateTokenRequestTFModel struct {
	AccessPolicyRid types.String `tfsdk:"access_policy_rid"`
	Data            types.String `tfsdk:"data"`
	TransformerRid  types.String `tfsdk:"transformer_rid"`
}

// TokenizerCreateTokenRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerCreateTokenRequestJSONClientModel struct {
	AccessPolicyRid *UserstoreResourceIDJSONClientModel `json:"access_policy_rid,omitempty"`
	Data            *string                             `json:"data,omitempty"`
	TransformerRid  *UserstoreResourceIDJSONClientModel `json:"transformer_rid,omitempty"`
}

// TokenizerCreateTokenRequestAttrTypes defines the attribute types for the TokenizerCreateTokenRequestAttributes schema.
var TokenizerCreateTokenRequestAttrTypes = map[string]attr.Type{
	"access_policy_rid": types.StringType,
	"data":              types.StringType,
	"transformer_rid":   types.StringType,
}

// TokenizerCreateTokenRequestAttributes defines the Terraform attributes schema.
var TokenizerCreateTokenRequestAttributes = map[string]schema.Attribute{
	"access_policy_rid": schema.StringAttribute{
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
	"data": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"transformer_rid": schema.StringAttribute{
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

// TokenizerCreateTokenRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerCreateTokenRequestTFModelToJSONClient(in *TokenizerCreateTokenRequestTFModel) (*TokenizerCreateTokenRequestJSONClientModel, error) {
	out := TokenizerCreateTokenRequestJSONClientModel{}
	var err error
	out.AccessPolicyRid, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.AccessPolicyRid)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy_rid\" field: %+v", err)
	}
	out.Data, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Data)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	out.TransformerRid, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.TransformerRid)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"transformer_rid\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerCreateTokenRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerCreateTokenRequestJSONClientModelToTF(in *TokenizerCreateTokenRequestJSONClientModel) (TokenizerCreateTokenRequestTFModel, error) {
	out := TokenizerCreateTokenRequestTFModel{}
	var err error
	out.AccessPolicyRid, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
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
	}(in.AccessPolicyRid)
	if err != nil {
		return TokenizerCreateTokenRequestTFModel{}, ucerr.Errorf("failed to convert \"access_policy_rid\" field: %+v", err)
	}
	out.Data, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Data)
	if err != nil {
		return TokenizerCreateTokenRequestTFModel{}, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	out.TransformerRid, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
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
	}(in.TransformerRid)
	if err != nil {
		return TokenizerCreateTokenRequestTFModel{}, ucerr.Errorf("failed to convert \"transformer_rid\" field: %+v", err)
	}
	return out, nil
}

// TokenizerCreateTokenResponseTFModel is a Terraform model struct for the TokenizerCreateTokenResponseAttributes schema.
type TokenizerCreateTokenResponseTFModel struct {
	Data types.String `tfsdk:"data"`
}

// TokenizerCreateTokenResponseJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerCreateTokenResponseJSONClientModel struct {
	Data *string `json:"data,omitempty"`
}

// TokenizerCreateTokenResponseAttrTypes defines the attribute types for the TokenizerCreateTokenResponseAttributes schema.
var TokenizerCreateTokenResponseAttrTypes = map[string]attr.Type{
	"data": types.StringType,
}

// TokenizerCreateTokenResponseAttributes defines the Terraform attributes schema.
var TokenizerCreateTokenResponseAttributes = map[string]schema.Attribute{
	"data": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerCreateTokenResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerCreateTokenResponseTFModelToJSONClient(in *TokenizerCreateTokenResponseTFModel) (*TokenizerCreateTokenResponseJSONClientModel, error) {
	out := TokenizerCreateTokenResponseJSONClientModel{}
	var err error
	out.Data, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Data)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerCreateTokenResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerCreateTokenResponseJSONClientModelToTF(in *TokenizerCreateTokenResponseJSONClientModel) (TokenizerCreateTokenResponseTFModel, error) {
	out := TokenizerCreateTokenResponseTFModel{}
	var err error
	out.Data, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Data)
	if err != nil {
		return TokenizerCreateTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	return out, nil
}

// TokenizerCreateTransformerRequestTFModel is a Terraform model struct for the TokenizerCreateTransformerRequestAttributes schema.
type TokenizerCreateTransformerRequestTFModel struct {
	Transformer types.Object `tfsdk:"transformer"`
}

// TokenizerCreateTransformerRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerCreateTransformerRequestJSONClientModel struct {
	Transformer *PolicyTransformerJSONClientModel `json:"transformer,omitempty"`
}

// TokenizerCreateTransformerRequestAttrTypes defines the attribute types for the TokenizerCreateTransformerRequestAttributes schema.
var TokenizerCreateTransformerRequestAttrTypes = map[string]attr.Type{
	"transformer": types.ObjectType{
		AttrTypes: PolicyTransformerAttrTypes,
	},
}

// TokenizerCreateTransformerRequestAttributes defines the Terraform attributes schema.
var TokenizerCreateTransformerRequestAttributes = map[string]schema.Attribute{
	"transformer": schema.SingleNestedAttribute{
		Attributes:          PolicyTransformerAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerCreateTransformerRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerCreateTransformerRequestTFModelToJSONClient(in *TokenizerCreateTransformerRequestTFModel) (*TokenizerCreateTransformerRequestJSONClientModel, error) {
	out := TokenizerCreateTransformerRequestJSONClientModel{}
	var err error
	out.Transformer, err = func(val *types.Object) (*PolicyTransformerJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := PolicyTransformerTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return PolicyTransformerTFModelToJSONClient(&tfModel)
	}(&in.Transformer)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"transformer\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerCreateTransformerRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerCreateTransformerRequestJSONClientModelToTF(in *TokenizerCreateTransformerRequestJSONClientModel) (TokenizerCreateTransformerRequestTFModel, error) {
	out := TokenizerCreateTransformerRequestTFModel{}
	var err error
	out.Transformer, err = func(val *PolicyTransformerJSONClientModel) (types.Object, error) {
		attrTypes := PolicyTransformerAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := PolicyTransformerJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyTransformerTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Transformer)
	if err != nil {
		return TokenizerCreateTransformerRequestTFModel{}, ucerr.Errorf("failed to convert \"transformer\" field: %+v", err)
	}
	return out, nil
}

// TokenizerInspectTokenRequestTFModel is a Terraform model struct for the TokenizerInspectTokenRequestAttributes schema.
type TokenizerInspectTokenRequestTFModel struct {
	Token types.String `tfsdk:"token"`
}

// TokenizerInspectTokenRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerInspectTokenRequestJSONClientModel struct {
	Token *string `json:"token,omitempty"`
}

// TokenizerInspectTokenRequestAttrTypes defines the attribute types for the TokenizerInspectTokenRequestAttributes schema.
var TokenizerInspectTokenRequestAttrTypes = map[string]attr.Type{
	"token": types.StringType,
}

// TokenizerInspectTokenRequestAttributes defines the Terraform attributes schema.
var TokenizerInspectTokenRequestAttributes = map[string]schema.Attribute{
	"token": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerInspectTokenRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerInspectTokenRequestTFModelToJSONClient(in *TokenizerInspectTokenRequestTFModel) (*TokenizerInspectTokenRequestJSONClientModel, error) {
	out := TokenizerInspectTokenRequestJSONClientModel{}
	var err error
	out.Token, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Token)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"token\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerInspectTokenRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerInspectTokenRequestJSONClientModelToTF(in *TokenizerInspectTokenRequestJSONClientModel) (TokenizerInspectTokenRequestTFModel, error) {
	out := TokenizerInspectTokenRequestTFModel{}
	var err error
	out.Token, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Token)
	if err != nil {
		return TokenizerInspectTokenRequestTFModel{}, ucerr.Errorf("failed to convert \"token\" field: %+v", err)
	}
	return out, nil
}

// TokenizerInspectTokenResponseTFModel is a Terraform model struct for the TokenizerInspectTokenResponseAttributes schema.
type TokenizerInspectTokenResponseTFModel struct {
	AccessPolicy               types.Object `tfsdk:"access_policy"`
	Created                    types.String `tfsdk:"created"`
	CurrentAccessPolicyVersion types.Int64  `tfsdk:"current_access_policy_version"`
	ID                         types.String `tfsdk:"id"`
	Token                      types.String `tfsdk:"token"`
	Transformer                types.Object `tfsdk:"transformer"`
	Updated                    types.String `tfsdk:"updated"`
}

// TokenizerInspectTokenResponseJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerInspectTokenResponseJSONClientModel struct {
	AccessPolicy               *PolicyAccessPolicyJSONClientModel `json:"access_policy,omitempty"`
	Created                    *string                            `json:"created,omitempty"`
	CurrentAccessPolicyVersion *int64                             `json:"current_access_policy_version,omitempty"`
	ID                         *uuid.UUID                         `json:"id,omitempty"`
	Token                      *string                            `json:"token,omitempty"`
	Transformer                *PolicyTransformerJSONClientModel  `json:"transformer,omitempty"`
	Updated                    *string                            `json:"updated,omitempty"`
}

// TokenizerInspectTokenResponseAttrTypes defines the attribute types for the TokenizerInspectTokenResponseAttributes schema.
var TokenizerInspectTokenResponseAttrTypes = map[string]attr.Type{
	"access_policy": types.ObjectType{
		AttrTypes: PolicyAccessPolicyAttrTypes,
	},
	"created":                       types.StringType,
	"current_access_policy_version": types.Int64Type,
	"id":                            types.StringType,
	"token":                         types.StringType,
	"transformer": types.ObjectType{
		AttrTypes: PolicyTransformerAttrTypes,
	},
	"updated": types.StringType,
}

// TokenizerInspectTokenResponseAttributes defines the Terraform attributes schema.
var TokenizerInspectTokenResponseAttributes = map[string]schema.Attribute{
	"access_policy": schema.SingleNestedAttribute{
		Attributes:          PolicyAccessPolicyAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"created": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"current_access_policy_version": schema.Int64Attribute{
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
	"token": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"transformer": schema.SingleNestedAttribute{
		Attributes:          PolicyTransformerAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"updated": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerInspectTokenResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerInspectTokenResponseTFModelToJSONClient(in *TokenizerInspectTokenResponseTFModel) (*TokenizerInspectTokenResponseJSONClientModel, error) {
	out := TokenizerInspectTokenResponseJSONClientModel{}
	var err error
	out.AccessPolicy, err = func(val *types.Object) (*PolicyAccessPolicyJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := PolicyAccessPolicyTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return PolicyAccessPolicyTFModelToJSONClient(&tfModel)
	}(&in.AccessPolicy)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	out.Created, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Created)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"created\" field: %+v", err)
	}
	out.CurrentAccessPolicyVersion, err = func(val *types.Int64) (*int64, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueInt64()
		return &converted, nil
	}(&in.CurrentAccessPolicyVersion)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"current_access_policy_version\" field: %+v", err)
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
	out.Token, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Token)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"token\" field: %+v", err)
	}
	out.Transformer, err = func(val *types.Object) (*PolicyTransformerJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := PolicyTransformerTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return PolicyTransformerTFModelToJSONClient(&tfModel)
	}(&in.Transformer)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"transformer\" field: %+v", err)
	}
	out.Updated, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Updated)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"updated\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerInspectTokenResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerInspectTokenResponseJSONClientModelToTF(in *TokenizerInspectTokenResponseJSONClientModel) (TokenizerInspectTokenResponseTFModel, error) {
	out := TokenizerInspectTokenResponseTFModel{}
	var err error
	out.AccessPolicy, err = func(val *PolicyAccessPolicyJSONClientModel) (types.Object, error) {
		attrTypes := PolicyAccessPolicyAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := PolicyAccessPolicyJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyAccessPolicyTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.AccessPolicy)
	if err != nil {
		return TokenizerInspectTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	out.Created, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Created)
	if err != nil {
		return TokenizerInspectTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"created\" field: %+v", err)
	}
	out.CurrentAccessPolicyVersion, err = func(val *int64) (types.Int64, error) {
		return types.Int64PointerValue(val), nil
	}(in.CurrentAccessPolicyVersion)
	if err != nil {
		return TokenizerInspectTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"current_access_policy_version\" field: %+v", err)
	}
	out.ID, err = func(val *uuid.UUID) (types.String, error) {
		if val == nil {
			return types.StringNull(), nil
		}
		return types.StringValue(val.String()), nil
	}(in.ID)
	if err != nil {
		return TokenizerInspectTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"id\" field: %+v", err)
	}
	out.Token, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Token)
	if err != nil {
		return TokenizerInspectTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"token\" field: %+v", err)
	}
	out.Transformer, err = func(val *PolicyTransformerJSONClientModel) (types.Object, error) {
		attrTypes := PolicyTransformerAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := PolicyTransformerJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyTransformerTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.Transformer)
	if err != nil {
		return TokenizerInspectTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"transformer\" field: %+v", err)
	}
	out.Updated, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Updated)
	if err != nil {
		return TokenizerInspectTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"updated\" field: %+v", err)
	}
	return out, nil
}

// TokenizerLookupOrCreateTokensRequestTFModel is a Terraform model struct for the TokenizerLookupOrCreateTokensRequestAttributes schema.
type TokenizerLookupOrCreateTokensRequestTFModel struct {
	AccessPolicyRids types.List `tfsdk:"access_policy_rids"`
	Data             types.List `tfsdk:"data"`
	TransformerRids  types.List `tfsdk:"transformer_rids"`
}

// TokenizerLookupOrCreateTokensRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerLookupOrCreateTokensRequestJSONClientModel struct {
	AccessPolicyRids *[]UserstoreResourceIDJSONClientModel `json:"access_policy_rids,omitempty"`
	Data             *[]string                             `json:"data,omitempty"`
	TransformerRids  *[]UserstoreResourceIDJSONClientModel `json:"transformer_rids,omitempty"`
}

// TokenizerLookupOrCreateTokensRequestAttrTypes defines the attribute types for the TokenizerLookupOrCreateTokensRequestAttributes schema.
var TokenizerLookupOrCreateTokensRequestAttrTypes = map[string]attr.Type{
	"access_policy_rids": types.ListType{
		ElemType: types.StringType,
	},
	"data": types.ListType{
		ElemType: types.StringType,
	},
	"transformer_rids": types.ListType{
		ElemType: types.StringType,
	},
}

// TokenizerLookupOrCreateTokensRequestAttributes defines the Terraform attributes schema.
var TokenizerLookupOrCreateTokensRequestAttributes = map[string]schema.Attribute{
	"access_policy_rids": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"data": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"transformer_rids": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerLookupOrCreateTokensRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerLookupOrCreateTokensRequestTFModelToJSONClient(in *TokenizerLookupOrCreateTokensRequestTFModel) (*TokenizerLookupOrCreateTokensRequestJSONClientModel, error) {
	out := TokenizerLookupOrCreateTokensRequestJSONClientModel{}
	var err error
	out.AccessPolicyRids, err = func(val *types.List) (*[]UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.AccessPolicyRids)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy_rids\" field: %+v", err)
	}
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
	out.TransformerRids, err = func(val *types.List) (*[]UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.TransformerRids)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"transformer_rids\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerLookupOrCreateTokensRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerLookupOrCreateTokensRequestJSONClientModelToTF(in *TokenizerLookupOrCreateTokensRequestJSONClientModel) (TokenizerLookupOrCreateTokensRequestTFModel, error) {
	out := TokenizerLookupOrCreateTokensRequestTFModel{}
	var err error
	out.AccessPolicyRids, err = func(val *[]UserstoreResourceIDJSONClientModel) (types.List, error) {
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
	}(in.AccessPolicyRids)
	if err != nil {
		return TokenizerLookupOrCreateTokensRequestTFModel{}, ucerr.Errorf("failed to convert \"access_policy_rids\" field: %+v", err)
	}
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
		return TokenizerLookupOrCreateTokensRequestTFModel{}, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	out.TransformerRids, err = func(val *[]UserstoreResourceIDJSONClientModel) (types.List, error) {
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
	}(in.TransformerRids)
	if err != nil {
		return TokenizerLookupOrCreateTokensRequestTFModel{}, ucerr.Errorf("failed to convert \"transformer_rids\" field: %+v", err)
	}
	return out, nil
}

// TokenizerLookupOrCreateTokensResponseTFModel is a Terraform model struct for the TokenizerLookupOrCreateTokensResponseAttributes schema.
type TokenizerLookupOrCreateTokensResponseTFModel struct {
	Tokens types.List `tfsdk:"tokens"`
}

// TokenizerLookupOrCreateTokensResponseJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerLookupOrCreateTokensResponseJSONClientModel struct {
	Tokens *[]string `json:"tokens,omitempty"`
}

// TokenizerLookupOrCreateTokensResponseAttrTypes defines the attribute types for the TokenizerLookupOrCreateTokensResponseAttributes schema.
var TokenizerLookupOrCreateTokensResponseAttrTypes = map[string]attr.Type{
	"tokens": types.ListType{
		ElemType: types.StringType,
	},
}

// TokenizerLookupOrCreateTokensResponseAttributes defines the Terraform attributes schema.
var TokenizerLookupOrCreateTokensResponseAttributes = map[string]schema.Attribute{
	"tokens": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerLookupOrCreateTokensResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerLookupOrCreateTokensResponseTFModelToJSONClient(in *TokenizerLookupOrCreateTokensResponseTFModel) (*TokenizerLookupOrCreateTokensResponseJSONClientModel, error) {
	out := TokenizerLookupOrCreateTokensResponseJSONClientModel{}
	var err error
	out.Tokens, err = func(val *types.List) (*[]string, error) {
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
	}(&in.Tokens)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"tokens\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerLookupOrCreateTokensResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerLookupOrCreateTokensResponseJSONClientModelToTF(in *TokenizerLookupOrCreateTokensResponseJSONClientModel) (TokenizerLookupOrCreateTokensResponseTFModel, error) {
	out := TokenizerLookupOrCreateTokensResponseTFModel{}
	var err error
	out.Tokens, err = func(val *[]string) (types.List, error) {
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
	}(in.Tokens)
	if err != nil {
		return TokenizerLookupOrCreateTokensResponseTFModel{}, ucerr.Errorf("failed to convert \"tokens\" field: %+v", err)
	}
	return out, nil
}

// TokenizerLookupTokensRequestTFModel is a Terraform model struct for the TokenizerLookupTokensRequestAttributes schema.
type TokenizerLookupTokensRequestTFModel struct {
	AccessPolicyRid types.String `tfsdk:"access_policy_rid"`
	Data            types.String `tfsdk:"data"`
	TransformerRid  types.String `tfsdk:"transformer_rid"`
}

// TokenizerLookupTokensRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerLookupTokensRequestJSONClientModel struct {
	AccessPolicyRid *UserstoreResourceIDJSONClientModel `json:"access_policy_rid,omitempty"`
	Data            *string                             `json:"data,omitempty"`
	TransformerRid  *UserstoreResourceIDJSONClientModel `json:"transformer_rid,omitempty"`
}

// TokenizerLookupTokensRequestAttrTypes defines the attribute types for the TokenizerLookupTokensRequestAttributes schema.
var TokenizerLookupTokensRequestAttrTypes = map[string]attr.Type{
	"access_policy_rid": types.StringType,
	"data":              types.StringType,
	"transformer_rid":   types.StringType,
}

// TokenizerLookupTokensRequestAttributes defines the Terraform attributes schema.
var TokenizerLookupTokensRequestAttributes = map[string]schema.Attribute{
	"access_policy_rid": schema.StringAttribute{
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
	"data": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"transformer_rid": schema.StringAttribute{
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

// TokenizerLookupTokensRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerLookupTokensRequestTFModelToJSONClient(in *TokenizerLookupTokensRequestTFModel) (*TokenizerLookupTokensRequestJSONClientModel, error) {
	out := TokenizerLookupTokensRequestJSONClientModel{}
	var err error
	out.AccessPolicyRid, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.AccessPolicyRid)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy_rid\" field: %+v", err)
	}
	out.Data, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Data)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	out.TransformerRid, err = func(val *types.String) (*UserstoreResourceIDJSONClientModel, error) {
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
	}(&in.TransformerRid)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"transformer_rid\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerLookupTokensRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerLookupTokensRequestJSONClientModelToTF(in *TokenizerLookupTokensRequestJSONClientModel) (TokenizerLookupTokensRequestTFModel, error) {
	out := TokenizerLookupTokensRequestTFModel{}
	var err error
	out.AccessPolicyRid, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
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
	}(in.AccessPolicyRid)
	if err != nil {
		return TokenizerLookupTokensRequestTFModel{}, ucerr.Errorf("failed to convert \"access_policy_rid\" field: %+v", err)
	}
	out.Data, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Data)
	if err != nil {
		return TokenizerLookupTokensRequestTFModel{}, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	out.TransformerRid, err = func(val *UserstoreResourceIDJSONClientModel) (types.String, error) {
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
	}(in.TransformerRid)
	if err != nil {
		return TokenizerLookupTokensRequestTFModel{}, ucerr.Errorf("failed to convert \"transformer_rid\" field: %+v", err)
	}
	return out, nil
}

// TokenizerLookupTokensResponseTFModel is a Terraform model struct for the TokenizerLookupTokensResponseAttributes schema.
type TokenizerLookupTokensResponseTFModel struct {
	Tokens types.List `tfsdk:"tokens"`
}

// TokenizerLookupTokensResponseJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerLookupTokensResponseJSONClientModel struct {
	Tokens *[]string `json:"tokens,omitempty"`
}

// TokenizerLookupTokensResponseAttrTypes defines the attribute types for the TokenizerLookupTokensResponseAttributes schema.
var TokenizerLookupTokensResponseAttrTypes = map[string]attr.Type{
	"tokens": types.ListType{
		ElemType: types.StringType,
	},
}

// TokenizerLookupTokensResponseAttributes defines the Terraform attributes schema.
var TokenizerLookupTokensResponseAttributes = map[string]schema.Attribute{
	"tokens": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerLookupTokensResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerLookupTokensResponseTFModelToJSONClient(in *TokenizerLookupTokensResponseTFModel) (*TokenizerLookupTokensResponseJSONClientModel, error) {
	out := TokenizerLookupTokensResponseJSONClientModel{}
	var err error
	out.Tokens, err = func(val *types.List) (*[]string, error) {
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
	}(&in.Tokens)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"tokens\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerLookupTokensResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerLookupTokensResponseJSONClientModelToTF(in *TokenizerLookupTokensResponseJSONClientModel) (TokenizerLookupTokensResponseTFModel, error) {
	out := TokenizerLookupTokensResponseTFModel{}
	var err error
	out.Tokens, err = func(val *[]string) (types.List, error) {
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
	}(in.Tokens)
	if err != nil {
		return TokenizerLookupTokensResponseTFModel{}, ucerr.Errorf("failed to convert \"tokens\" field: %+v", err)
	}
	return out, nil
}

// TokenizerResolveTokenResponseTFModel is a Terraform model struct for the TokenizerResolveTokenResponseAttributes schema.
type TokenizerResolveTokenResponseTFModel struct {
	Data  types.String `tfsdk:"data"`
	Token types.String `tfsdk:"token"`
}

// TokenizerResolveTokenResponseJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerResolveTokenResponseJSONClientModel struct {
	Data  *string `json:"data,omitempty"`
	Token *string `json:"token,omitempty"`
}

// TokenizerResolveTokenResponseAttrTypes defines the attribute types for the TokenizerResolveTokenResponseAttributes schema.
var TokenizerResolveTokenResponseAttrTypes = map[string]attr.Type{
	"data":  types.StringType,
	"token": types.StringType,
}

// TokenizerResolveTokenResponseAttributes defines the Terraform attributes schema.
var TokenizerResolveTokenResponseAttributes = map[string]schema.Attribute{
	"data": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
	"token": schema.StringAttribute{
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerResolveTokenResponseTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerResolveTokenResponseTFModelToJSONClient(in *TokenizerResolveTokenResponseTFModel) (*TokenizerResolveTokenResponseJSONClientModel, error) {
	out := TokenizerResolveTokenResponseJSONClientModel{}
	var err error
	out.Data, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Data)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	out.Token, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.Token)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"token\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerResolveTokenResponseJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerResolveTokenResponseJSONClientModelToTF(in *TokenizerResolveTokenResponseJSONClientModel) (TokenizerResolveTokenResponseTFModel, error) {
	out := TokenizerResolveTokenResponseTFModel{}
	var err error
	out.Data, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Data)
	if err != nil {
		return TokenizerResolveTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"data\" field: %+v", err)
	}
	out.Token, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Token)
	if err != nil {
		return TokenizerResolveTokenResponseTFModel{}, ucerr.Errorf("failed to convert \"token\" field: %+v", err)
	}
	return out, nil
}

// TokenizerUpdateAccessPolicyRequestTFModel is a Terraform model struct for the TokenizerUpdateAccessPolicyRequestAttributes schema.
type TokenizerUpdateAccessPolicyRequestTFModel struct {
	AccessPolicy types.Object `tfsdk:"access_policy"`
}

// TokenizerUpdateAccessPolicyRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerUpdateAccessPolicyRequestJSONClientModel struct {
	AccessPolicy *PolicyAccessPolicyJSONClientModel `json:"access_policy,omitempty"`
}

// TokenizerUpdateAccessPolicyRequestAttrTypes defines the attribute types for the TokenizerUpdateAccessPolicyRequestAttributes schema.
var TokenizerUpdateAccessPolicyRequestAttrTypes = map[string]attr.Type{
	"access_policy": types.ObjectType{
		AttrTypes: PolicyAccessPolicyAttrTypes,
	},
}

// TokenizerUpdateAccessPolicyRequestAttributes defines the Terraform attributes schema.
var TokenizerUpdateAccessPolicyRequestAttributes = map[string]schema.Attribute{
	"access_policy": schema.SingleNestedAttribute{
		Attributes:          PolicyAccessPolicyAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerUpdateAccessPolicyRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerUpdateAccessPolicyRequestTFModelToJSONClient(in *TokenizerUpdateAccessPolicyRequestTFModel) (*TokenizerUpdateAccessPolicyRequestJSONClientModel, error) {
	out := TokenizerUpdateAccessPolicyRequestJSONClientModel{}
	var err error
	out.AccessPolicy, err = func(val *types.Object) (*PolicyAccessPolicyJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := PolicyAccessPolicyTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return PolicyAccessPolicyTFModelToJSONClient(&tfModel)
	}(&in.AccessPolicy)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerUpdateAccessPolicyRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerUpdateAccessPolicyRequestJSONClientModelToTF(in *TokenizerUpdateAccessPolicyRequestJSONClientModel) (TokenizerUpdateAccessPolicyRequestTFModel, error) {
	out := TokenizerUpdateAccessPolicyRequestTFModel{}
	var err error
	out.AccessPolicy, err = func(val *PolicyAccessPolicyJSONClientModel) (types.Object, error) {
		attrTypes := PolicyAccessPolicyAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := PolicyAccessPolicyJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyAccessPolicyTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.AccessPolicy)
	if err != nil {
		return TokenizerUpdateAccessPolicyRequestTFModel{}, ucerr.Errorf("failed to convert \"access_policy\" field: %+v", err)
	}
	return out, nil
}

// TokenizerUpdateAccessPolicyTemplateRequestTFModel is a Terraform model struct for the TokenizerUpdateAccessPolicyTemplateRequestAttributes schema.
type TokenizerUpdateAccessPolicyTemplateRequestTFModel struct {
	AccessPolicyTemplate types.Object `tfsdk:"access_policy_template"`
}

// TokenizerUpdateAccessPolicyTemplateRequestJSONClientModel stores data for use with jsonclient for making API requests.
type TokenizerUpdateAccessPolicyTemplateRequestJSONClientModel struct {
	AccessPolicyTemplate *PolicyAccessPolicyTemplateJSONClientModel `json:"access_policy_template,omitempty"`
}

// TokenizerUpdateAccessPolicyTemplateRequestAttrTypes defines the attribute types for the TokenizerUpdateAccessPolicyTemplateRequestAttributes schema.
var TokenizerUpdateAccessPolicyTemplateRequestAttrTypes = map[string]attr.Type{
	"access_policy_template": types.ObjectType{
		AttrTypes: PolicyAccessPolicyTemplateAttrTypes,
	},
}

// TokenizerUpdateAccessPolicyTemplateRequestAttributes defines the Terraform attributes schema.
var TokenizerUpdateAccessPolicyTemplateRequestAttributes = map[string]schema.Attribute{
	"access_policy_template": schema.SingleNestedAttribute{
		Attributes:          PolicyAccessPolicyTemplateAttributes,
		Computed:            true,
		Description:         "",
		MarkdownDescription: "",
		Optional:            true,
	},
}

// TokenizerUpdateAccessPolicyTemplateRequestTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func TokenizerUpdateAccessPolicyTemplateRequestTFModelToJSONClient(in *TokenizerUpdateAccessPolicyTemplateRequestTFModel) (*TokenizerUpdateAccessPolicyTemplateRequestJSONClientModel, error) {
	out := TokenizerUpdateAccessPolicyTemplateRequestJSONClientModel{}
	var err error
	out.AccessPolicyTemplate, err = func(val *types.Object) (*PolicyAccessPolicyTemplateJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}

		attrs := val.Attributes()

		tfModel := PolicyAccessPolicyTemplateTFModel{}
		reflected := reflect.ValueOf(&tfModel)
		tfsdkNamesToFieldNames := map[string]string{}
		for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
			tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
		}
		for k, v := range attrs {
			reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
		}
		return PolicyAccessPolicyTemplateTFModelToJSONClient(&tfModel)
	}(&in.AccessPolicyTemplate)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"access_policy_template\" field: %+v", err)
	}
	return &out, nil
}

// TokenizerUpdateAccessPolicyTemplateRequestJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func TokenizerUpdateAccessPolicyTemplateRequestJSONClientModelToTF(in *TokenizerUpdateAccessPolicyTemplateRequestJSONClientModel) (TokenizerUpdateAccessPolicyTemplateRequestTFModel, error) {
	out := TokenizerUpdateAccessPolicyTemplateRequestTFModel{}
	var err error
	out.AccessPolicyTemplate, err = func(val *PolicyAccessPolicyTemplateJSONClientModel) (types.Object, error) {
		attrTypes := PolicyAccessPolicyTemplateAttrTypes

		if val == nil {
			return types.ObjectNull(attrTypes), nil
		}

		tfModel, err := PolicyAccessPolicyTemplateJSONClientModelToTF(val)
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
			return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert PolicyAccessPolicyTemplateTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
		}
		return objVal, nil
	}(in.AccessPolicyTemplate)
	if err != nil {
		return TokenizerUpdateAccessPolicyTemplateRequestTFModel{}, ucerr.Errorf("failed to convert \"access_policy_template\" field: %+v", err)
	}
	return out, nil
}

// UserstoreColumnConstraintsTFModel is a Terraform model struct for the UserstoreColumnConstraintsAttributes schema.
type UserstoreColumnConstraintsTFModel struct {
	Fields            types.List `tfsdk:"fields"`
	ImmutableRequired types.Bool `tfsdk:"immutable_required"`
	UniqueIDRequired  types.Bool `tfsdk:"unique_id_required"`
	UniqueRequired    types.Bool `tfsdk:"unique_required"`
}

// UserstoreColumnConstraintsJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreColumnConstraintsJSONClientModel struct {
	Fields            *[]UserstoreColumnFieldJSONClientModel `json:"fields,omitempty"`
	ImmutableRequired *bool                                  `json:"immutable_required,omitempty"`
	UniqueIDRequired  *bool                                  `json:"unique_id_required,omitempty"`
	UniqueRequired    *bool                                  `json:"unique_required,omitempty"`
}

// UserstoreColumnConstraintsAttrTypes defines the attribute types for the UserstoreColumnConstraintsAttributes schema.
var UserstoreColumnConstraintsAttrTypes = map[string]attr.Type{
	"fields": types.ListType{
		ElemType: types.ObjectType{
			AttrTypes: UserstoreColumnFieldAttrTypes,
		},
	},
	"immutable_required": types.BoolType,
	"unique_id_required": types.BoolType,
	"unique_required":    types.BoolType,
}

// UserstoreColumnConstraintsAttributes defines the Terraform attributes schema.
var UserstoreColumnConstraintsAttributes = map[string]schema.Attribute{
	"fields": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: UserstoreColumnFieldAttributes,
		},
		Computed:            true,
		Description:         "The set of fields associated with a column of type composite. Fields cannot be specified if the column type is not composite.",
		MarkdownDescription: "The set of fields associated with a column of type composite. Fields cannot be specified if the column type is not composite.",
		Optional:            true,
	},
	"immutable_required": schema.BoolAttribute{
		Computed:            true,
		Description:         "Can be enabled when unique_id_required is enabled. If true, values for the associated column cannot be modified, but can be added or removed.",
		MarkdownDescription: "Can be enabled when unique_id_required is enabled. If true, values for the associated column cannot be modified, but can be added or removed.",
		Optional:            true,
	},
	"unique_id_required": schema.BoolAttribute{
		Computed:            true,
		Description:         "Can be enabled for column type composite or address. If true, each value for the associated column must have a unique string ID, which can either be provided or generated by backend.",
		MarkdownDescription: "Can be enabled for column type composite or address. If true, each value for the associated column must have a unique string ID, which can either be provided or generated by backend.",
		Optional:            true,
	},
	"unique_required": schema.BoolAttribute{
		Computed:            true,
		Description:         "If true, each value for the associated column must be unique for the user. This is primarily useful for array columns.",
		MarkdownDescription: "If true, each value for the associated column must be unique for the user. This is primarily useful for array columns.",
		Optional:            true,
	},
}

// UserstoreColumnConstraintsTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreColumnConstraintsTFModelToJSONClient(in *UserstoreColumnConstraintsTFModel) (*UserstoreColumnConstraintsJSONClientModel, error) {
	out := UserstoreColumnConstraintsJSONClientModel{}
	var err error
	out.Fields, err = func(val *types.List) (*[]UserstoreColumnFieldJSONClientModel, error) {
		if val == nil || val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		var out = []UserstoreColumnFieldJSONClientModel{}
		for _, elem := range val.Elements() {
			elemTyped, ok := elem.(types.Object)
			if !ok {
				return nil, ucerr.Errorf("unexpected type %s in list", elem.Type(context.Background()).String())
			}
			converted, err := func(val *types.Object) (*UserstoreColumnFieldJSONClientModel, error) {
				if val == nil || val.IsNull() || val.IsUnknown() {
					return nil, nil
				}

				attrs := val.Attributes()

				tfModel := UserstoreColumnFieldTFModel{}
				reflected := reflect.ValueOf(&tfModel)
				tfsdkNamesToFieldNames := map[string]string{}
				for i := 0; i < reflect.Indirect(reflected).NumField(); i++ {
					tfsdkNamesToFieldNames[reflect.Indirect(reflected).Type().Field(i).Tag.Get("tfsdk")] = reflect.Indirect(reflected).Type().Field(i).Name
				}
				for k, v := range attrs {
					reflect.Indirect(reflected).FieldByName(tfsdkNamesToFieldNames[k]).Set(reflect.ValueOf(v))
				}
				return UserstoreColumnFieldTFModelToJSONClient(&tfModel)
			}(&elemTyped)
			if err != nil {
				return nil, ucerr.Wrap(err)
			}
			out = append(out, *converted)
		}
		return &out, nil
	}(&in.Fields)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"fields\" field: %+v", err)
	}
	out.ImmutableRequired, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.ImmutableRequired)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"immutable_required\" field: %+v", err)
	}
	out.UniqueIDRequired, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.UniqueIDRequired)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"unique_id_required\" field: %+v", err)
	}
	out.UniqueRequired, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.UniqueRequired)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"unique_required\" field: %+v", err)
	}
	return &out, nil
}

// UserstoreColumnConstraintsJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreColumnConstraintsJSONClientModelToTF(in *UserstoreColumnConstraintsJSONClientModel) (UserstoreColumnConstraintsTFModel, error) {
	out := UserstoreColumnConstraintsTFModel{}
	var err error
	out.Fields, err = func(val *[]UserstoreColumnFieldJSONClientModel) (types.List, error) {
		childAttrType := types.ObjectType{
			AttrTypes: UserstoreColumnFieldAttrTypes,
		}
		if val == nil {
			return types.ListNull(childAttrType), nil
		}
		var out []attr.Value
		for _, elem := range *val {
			converted, err := func(val *UserstoreColumnFieldJSONClientModel) (types.Object, error) {
				attrTypes := UserstoreColumnFieldAttrTypes

				if val == nil {
					return types.ObjectNull(attrTypes), nil
				}

				tfModel, err := UserstoreColumnFieldJSONClientModelToTF(val)
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
					return types.ObjectNull(attrTypes), ucerr.Errorf("failed to convert UserstoreColumnFieldTFModel to terraform basetypes.Object: %s", diag.Errors()[0].Detail())
				}
				return objVal, nil
			}(&elem)
			if err != nil {
				return types.ListNull(childAttrType), ucerr.Wrap(err)
			}
			out = append(out, converted)
		}
		return types.ListValueMust(childAttrType, out), nil
	}(in.Fields)
	if err != nil {
		return UserstoreColumnConstraintsTFModel{}, ucerr.Errorf("failed to convert \"fields\" field: %+v", err)
	}
	out.ImmutableRequired, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.ImmutableRequired)
	if err != nil {
		return UserstoreColumnConstraintsTFModel{}, ucerr.Errorf("failed to convert \"immutable_required\" field: %+v", err)
	}
	out.UniqueIDRequired, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.UniqueIDRequired)
	if err != nil {
		return UserstoreColumnConstraintsTFModel{}, ucerr.Errorf("failed to convert \"unique_id_required\" field: %+v", err)
	}
	out.UniqueRequired, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.UniqueRequired)
	if err != nil {
		return UserstoreColumnConstraintsTFModel{}, ucerr.Errorf("failed to convert \"unique_required\" field: %+v", err)
	}
	return out, nil
}

// UserstoreColumnFieldTFModel is a Terraform model struct for the UserstoreColumnFieldAttributes schema.
type UserstoreColumnFieldTFModel struct {
	CamelCaseName       types.String `tfsdk:"camel_case_name"`
	IgnoreForUniqueness types.Bool   `tfsdk:"ignore_for_uniqueness"`
	Name                types.String `tfsdk:"name"`
	Required            types.Bool   `tfsdk:"required"`
	StructName          types.String `tfsdk:"struct_name"`
	Type                types.String `tfsdk:"type"`
}

// UserstoreColumnFieldJSONClientModel stores data for use with jsonclient for making API requests.
type UserstoreColumnFieldJSONClientModel struct {
	CamelCaseName       *string `json:"camel_case_name,omitempty"`
	IgnoreForUniqueness *bool   `json:"ignore_for_uniqueness,omitempty"`
	Name                *string `json:"name,omitempty"`
	Required            *bool   `json:"required,omitempty"`
	StructName          *string `json:"struct_name,omitempty"`
	Type                *string `json:"type,omitempty"`
}

// UserstoreColumnFieldAttrTypes defines the attribute types for the UserstoreColumnFieldAttributes schema.
var UserstoreColumnFieldAttrTypes = map[string]attr.Type{
	"camel_case_name":       types.StringType,
	"ignore_for_uniqueness": types.BoolType,
	"name":                  types.StringType,
	"required":              types.BoolType,
	"struct_name":           types.StringType,
	"type":                  types.StringType,
}

// UserstoreColumnFieldAttributes defines the Terraform attributes schema.
var UserstoreColumnFieldAttributes = map[string]schema.Attribute{
	"camel_case_name": schema.StringAttribute{
		Computed:            true,
		Description:         "Read-only camel-case version of field name, with underscores stripped out. (ex. IDField1)",
		MarkdownDescription: "Read-only camel-case version of field name, with underscores stripped out. (ex. IDField1)",
		Optional:            true,
	},
	"ignore_for_uniqueness": schema.BoolAttribute{
		Computed:            true,
		Description:         "If true, field value will be ignored when comparing two composite value for a uniqueness check.",
		MarkdownDescription: "If true, field value will be ignored when comparing two composite value for a uniqueness check.",
		Optional:            true,
	},
	"name": schema.StringAttribute{
		Description:         "Each part of name must be capitalized or all-caps, separated by underscores. Names may contain alphanumeric characters, and the first part must start with a letter, while other parts may start with a number. (ex. ID_Field_1)",
		MarkdownDescription: "Each part of name must be capitalized or all-caps, separated by underscores. Names may contain alphanumeric characters, and the first part must start with a letter, while other parts may start with a number. (ex. ID_Field_1)",
		Required:            true,
	},
	"required": schema.BoolAttribute{
		Computed:            true,
		Description:         "Whether a value must be specified for the field.",
		MarkdownDescription: "Whether a value must be specified for the field.",
		Optional:            true,
	},
	"struct_name": schema.StringAttribute{
		Computed:            true,
		Description:         "Read-only snake-case version of field name, with all letters lowercase. (ex. id_field_1)",
		MarkdownDescription: "Read-only snake-case version of field name, with all letters lowercase. (ex. id_field_1)",
		Optional:            true,
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf([]string{"address", "birthdate", "boolean", "composite", "date", "e164_phonenumber", "email", "integer", "phonenumber", "ssn", "string", "timestamp", "uuid"}...),
		},
		Description:         "Valid values: `address`, `birthdate`, `boolean`, `composite`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		MarkdownDescription: "Valid values: `address`, `birthdate`, `boolean`, `composite`, `date`, `e164_phonenumber`, `email`, `integer`, `phonenumber`, `ssn`, `string`, `timestamp`, `uuid`",
		Required:            true,
	},
}

// UserstoreColumnFieldTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UserstoreColumnFieldTFModelToJSONClient(in *UserstoreColumnFieldTFModel) (*UserstoreColumnFieldJSONClientModel, error) {
	out := UserstoreColumnFieldJSONClientModel{}
	var err error
	out.CamelCaseName, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.CamelCaseName)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"camel_case_name\" field: %+v", err)
	}
	out.IgnoreForUniqueness, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.IgnoreForUniqueness)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"ignore_for_uniqueness\" field: %+v", err)
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
	out.Required, err = func(val *types.Bool) (*bool, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueBool()
		return &converted, nil
	}(&in.Required)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"required\" field: %+v", err)
	}
	out.StructName, err = func(val *types.String) (*string, error) {
		if val.IsNull() || val.IsUnknown() {
			return nil, nil
		}
		converted := val.ValueString()
		return &converted, nil
	}(&in.StructName)
	if err != nil {
		return nil, ucerr.Errorf("failed to convert \"struct_name\" field: %+v", err)
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

// UserstoreColumnFieldJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UserstoreColumnFieldJSONClientModelToTF(in *UserstoreColumnFieldJSONClientModel) (UserstoreColumnFieldTFModel, error) {
	out := UserstoreColumnFieldTFModel{}
	var err error
	out.CamelCaseName, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.CamelCaseName)
	if err != nil {
		return UserstoreColumnFieldTFModel{}, ucerr.Errorf("failed to convert \"camel_case_name\" field: %+v", err)
	}
	out.IgnoreForUniqueness, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.IgnoreForUniqueness)
	if err != nil {
		return UserstoreColumnFieldTFModel{}, ucerr.Errorf("failed to convert \"ignore_for_uniqueness\" field: %+v", err)
	}
	out.Name, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Name)
	if err != nil {
		return UserstoreColumnFieldTFModel{}, ucerr.Errorf("failed to convert \"name\" field: %+v", err)
	}
	out.Required, err = func(val *bool) (types.Bool, error) {
		return types.BoolPointerValue(val), nil
	}(in.Required)
	if err != nil {
		return UserstoreColumnFieldTFModel{}, ucerr.Errorf("failed to convert \"required\" field: %+v", err)
	}
	out.StructName, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.StructName)
	if err != nil {
		return UserstoreColumnFieldTFModel{}, ucerr.Errorf("failed to convert \"struct_name\" field: %+v", err)
	}
	out.Type, err = func(val *string) (types.String, error) {
		return types.StringPointerValue(val), nil
	}(in.Type)
	if err != nil {
		return UserstoreColumnFieldTFModel{}, ucerr.Errorf("failed to convert \"type\" field: %+v", err)
	}
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

// UuidarrayUuidarrayTFModel is a Terraform model struct for the UuidarrayUuidarrayAttributes schema.
type UuidarrayUuidarrayTFModel struct {
}

// UuidarrayUuidarrayJSONClientModel stores data for use with jsonclient for making API requests.
type UuidarrayUuidarrayJSONClientModel struct {
}

// UuidarrayUuidarrayAttrTypes defines the attribute types for the UuidarrayUuidarrayAttributes schema.
var UuidarrayUuidarrayAttrTypes = map[string]attr.Type{}

// UuidarrayUuidarrayAttributes defines the Terraform attributes schema.
var UuidarrayUuidarrayAttributes = map[string]schema.Attribute{}

// UuidarrayUuidarrayTFModelToJSONClient converts a Terraform model struct to a jsonclient model struct.
func UuidarrayUuidarrayTFModelToJSONClient(in *UuidarrayUuidarrayTFModel) (*UuidarrayUuidarrayJSONClientModel, error) {
	out := UuidarrayUuidarrayJSONClientModel{}
	return &out, nil
}

// UuidarrayUuidarrayJSONClientModelToTF converts a jsonclient model struct to a Terraform model struct.
func UuidarrayUuidarrayJSONClientModelToTF(in *UuidarrayUuidarrayJSONClientModel) (UuidarrayUuidarrayTFModel, error) {
	out := UuidarrayUuidarrayTFModel{}
	return out, nil
}

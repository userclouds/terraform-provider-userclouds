// NOTE: automatically generated file -- DO NOT EDIT

package tokenizer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"userclouds.com/infra/jsonclient"
)

// Note: revive is complaining about stuttering in the generated schema names (e.g. an OpenAPI
// schema might be called "UserstoreColumnTFModel", and then we generate it in the "userstore"
// package, so it becomes "userstore.UserstoreColumnTFModel"), but these names are coming from the
// OpenAPI spec and no one is going to be reading this generated code anyways, so we should just
// silence the rule.
//revive:disable:exported

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TransformerResource{}
var _ resource.ResourceWithImportState = &TransformerResource{}
var _ resource.ResourceWithModifyPlan = &TransformerResource{}

// NewTransformerResource returns a new instance of the resource.
func NewTransformerResource() resource.Resource {
	return &TransformerResource{}
}

// TransformerResource defines the resource implementation.
type TransformerResource struct {
	client *jsonclient.Client
}

// Metadata describes the resource metadata.
func (r *TransformerResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transformer"
}

// Schema describes the resource schema.
func (r *TransformerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "Manages a Tokenizer transformer. For more details refer to the transformer documentation: https://docs.userclouds.com/docs/token-transformers",
		MarkdownDescription: "Manages a Tokenizer transformer. For more details refer to the [transformer documentation](https://docs.userclouds.com/docs/token-transformers).",
		Attributes:          PolicyTransformerAttributes,
	}
}

// Configure configures the resource.
func (r *TransformerResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*jsonclient.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *jsonclient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Create creates a new resource.
func (r *TransformerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PolicyTransformerTFModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	jsonclientModel, err := PolicyTransformerTFModelToJSONClient(data)
	if err != nil {
		resp.Diagnostics.AddError("Error converting userclouds_transformer to JSON", err.Error())
		return
	}

	url := "/tokenizer/policies/transformation"
	// The collection path won't include all path parameters, but it should be fine to ReplaceAll
	// for all the path parameters anyways, since ReplaceAll will just be a no-op if a path
	// parameter doesn't appear in the string.
	url = strings.ReplaceAll(url, "{id}", data.ID.ValueString())
	body := TokenizerCreateTransformerRequestJSONClientModel{
		Transformer: jsonclientModel,
	}

	marshaled, err := json.Marshal(body)
	if err != nil {
		resp.Diagnostics.AddError("Error serializing userclouds_transformer JSON request body", err.Error())
		return
	}
	tflog.Trace(ctx, fmt.Sprintf("POST %s: %s", url, string(marshaled)))

	var apiResp PolicyTransformerJSONClientModel
	if err := r.client.Post(ctx, url, body, &apiResp); err != nil {
		resp.Diagnostics.AddError("Error creating userclouds_transformer", err.Error())
		return
	}
	created := apiResp
	createdTF, err := PolicyTransformerJSONClientModelToTF(&created)
	if err != nil {
		resp.Diagnostics.AddError("Error converting userclouds_transformer response JSON to Terraform state", err.Error())
		return
	}

	tflog.Trace(ctx, "successfully created userclouds_transformer with ID "+created.ID.String())
	resp.Diagnostics.Append(resp.State.Set(ctx, &createdTF)...)
}

// Read reads the existing resource state.
func (r *TransformerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var oldState *PolicyTransformerTFModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &oldState)...)

	if resp.Diagnostics.HasError() {
		return
	}

	url := "/tokenizer/policies/transformation/{id}"
	url = strings.ReplaceAll(url, "{id}", oldState.ID.ValueString())
	tflog.Trace(ctx, fmt.Sprintf("GET %s", url))
	var apiResp PolicyTransformerJSONClientModel
	if err := r.client.Get(ctx, url, &apiResp); err != nil {
		var jce jsonclient.Error
		if errors.As(err, &jce) && (jce.StatusCode == http.StatusNotFound) {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError("Error reading userclouds_transformer", err.Error())
		return
	}
	current := apiResp

	newState, err := PolicyTransformerJSONClientModelToTF(&current)
	if err != nil {
		resp.Diagnostics.AddError("Error converting userclouds_transformer response JSON to Terraform state", err.Error())
		return
	}

	tflog.Trace(ctx, "successfully read userclouds_transformer with ID "+current.ID.String())
	resp.Diagnostics.Append(resp.State.Set(ctx, &newState)...)
}

// Update updates an existing resource.
func (r *TransformerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddError(
		"Updates are not supported for userclouds_transformer",
		"Terraform should have suggested destroying and re-creating the resource. Please report this as a provider bug.",
	)
}

// Delete deletes an existing resource.
func (r *TransformerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PolicyTransformerTFModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	url := "/tokenizer/policies/transformation/{id}"
	url = strings.ReplaceAll(url, "{id}", data.ID.ValueString())
	tflog.Trace(ctx, fmt.Sprintf("GET %s", url))
	if err := r.client.Delete(ctx, url, nil); err != nil {
		resp.Diagnostics.AddError("Error deleting userclouds_transformer", err.Error())
		return
	}
}

// ImportState imports an existing resource into Terraform.
func (r *TransformerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// ModifyPlan forces replacement on modification, since updates are not supported for this resource
func (r *TransformerResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Do not replace on resource creation.
	if req.State.Raw.IsNull() {
		return
	}

	// Do not replace on resource destroy.
	if req.Plan.Raw.IsNull() {
		return
	}

	// Do not replace if the plans and states are equal.
	if req.Plan.Raw.Equal(req.State.Raw) {
		return
	}

	// TODO: does this work, or do we need to enumerate all the fields with path.Root()?
	resp.RequiresReplace.Append(path.Empty())
}

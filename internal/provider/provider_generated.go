// NOTE: automatically generated file -- DO NOT EDIT

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/userclouds/userclouds/terraform-provider-userclouds/internal/provider/tokenizer"
	"github.com/userclouds/userclouds/terraform-provider-userclouds/internal/provider/userstore"
)

var generatedResources = []func() resource.Resource{
	userstore.NewUserstoreColumnResource,
	userstore.NewUserstoreColumnPreDeleteRetentionDurationResource,
	userstore.NewUserstoreColumnPostDeleteRetentionDurationResource,
	userstore.NewUserstoreAccessorResource,
	userstore.NewUserstoreMutatorResource,
	userstore.NewUserstorePurposeResource,
	tokenizer.NewAccessPolicyResource,
	tokenizer.NewAccessPolicyTemplateResource,
	tokenizer.NewTransformerResource,
}

var generatedDataSources = []func() datasource.DataSource{}

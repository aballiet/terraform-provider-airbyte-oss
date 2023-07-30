// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDefinitionManifestResourceModel) ToCreateSDKType() *shared.DeclarativeSourceDefinitionCreateManifestRequestBody {
	description := r.DeclarativeManifest.Description.ValueString()
	manifest := shared.DeclarativeManifest{}
	spec := shared.SourceDefinitionSpecification{}
	version := r.DeclarativeManifest.Version.ValueInt64()
	declarativeManifest := shared.DeclarativeSourceManifest{
		Description: description,
		Manifest:    manifest,
		Spec:        spec,
		Version:     version,
	}
	setAsActiveManifest := r.SetAsActiveManifest.ValueBool()
	out := shared.DeclarativeSourceDefinitionCreateManifestRequestBody{
		DeclarativeManifest: declarativeManifest,
		SetAsActiveManifest: setAsActiveManifest,
	}
	return &out
}

func (r *SourceDefinitionManifestResourceModel) ToUpdateSDKType() *shared.UpdateActiveManifestRequestBody {
	out := shared.UpdateActiveManifestRequestBody{}
	return &out
}

func (r *SourceDefinitionManifestResourceModel) RefreshFromCreateResponse(resp *shared.NotFoundKnownExceptionInfo) {
	if resp.ExceptionClassName != nil {
		r.ExceptionClassName = types.StringValue(*resp.ExceptionClassName)
	} else {
		r.ExceptionClassName = types.StringNull()
	}
	r.ExceptionStack = nil
	for _, v := range resp.ExceptionStack {
		r.ExceptionStack = append(r.ExceptionStack, types.StringValue(v))
	}
	if resp.ID != nil {
		r.ID = types.StringValue(*resp.ID)
	} else {
		r.ID = types.StringNull()
	}
	r.Message = types.StringValue(resp.Message)
	if resp.RootCauseExceptionClassName != nil {
		r.RootCauseExceptionClassName = types.StringValue(*resp.RootCauseExceptionClassName)
	} else {
		r.RootCauseExceptionClassName = types.StringNull()
	}
	r.RootCauseExceptionStack = nil
	for _, v := range resp.RootCauseExceptionStack {
		r.RootCauseExceptionStack = append(r.RootCauseExceptionStack, types.StringValue(v))
	}
}

func (r *SourceDefinitionManifestResourceModel) RefreshFromUpdateResponse(resp *shared.NotFoundKnownExceptionInfo) {
	r.RefreshFromCreateResponse(resp)
}

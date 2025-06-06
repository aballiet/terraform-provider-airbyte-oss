// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *OperationDataSourceModel) RefreshFromSharedOperationRead(resp *shared.OperationRead) {
	r.Name = types.StringValue(resp.Name)
	r.OperationID = types.StringValue(resp.OperationID)
	if resp.OperatorConfiguration.Dbt == nil {
		r.OperatorConfiguration.Dbt = nil
	} else {
		r.OperatorConfiguration.Dbt = &OperatorDbt{}
		r.OperatorConfiguration.Dbt.DbtArguments = types.StringPointerValue(resp.OperatorConfiguration.Dbt.DbtArguments)
		r.OperatorConfiguration.Dbt.DockerImage = types.StringPointerValue(resp.OperatorConfiguration.Dbt.DockerImage)
		r.OperatorConfiguration.Dbt.GitRepoBranch = types.StringPointerValue(resp.OperatorConfiguration.Dbt.GitRepoBranch)
		r.OperatorConfiguration.Dbt.GitRepoURL = types.StringValue(resp.OperatorConfiguration.Dbt.GitRepoURL)
	}
	if resp.OperatorConfiguration.Normalization == nil {
		r.OperatorConfiguration.Normalization = nil
	} else {
		r.OperatorConfiguration.Normalization = &OperatorNormalization{}
		if resp.OperatorConfiguration.Normalization.Option != nil {
			r.OperatorConfiguration.Normalization.Option = types.StringValue(string(*resp.OperatorConfiguration.Normalization.Option))
		} else {
			r.OperatorConfiguration.Normalization.Option = types.StringNull()
		}
	}
	r.OperatorConfiguration.OperatorType = types.StringValue(string(resp.OperatorConfiguration.OperatorType))
	if resp.OperatorConfiguration.Webhook == nil {
		r.OperatorConfiguration.Webhook = nil
	} else {
		r.OperatorConfiguration.Webhook = &OperatorWebhook{}
		if resp.OperatorConfiguration.Webhook.DbtCloud == nil {
			r.OperatorConfiguration.Webhook.DbtCloud = nil
		} else {
			r.OperatorConfiguration.Webhook.DbtCloud = &DbtCloud{}
			r.OperatorConfiguration.Webhook.DbtCloud.AccountID = types.Int64Value(resp.OperatorConfiguration.Webhook.DbtCloud.AccountID)
			r.OperatorConfiguration.Webhook.DbtCloud.JobID = types.Int64Value(resp.OperatorConfiguration.Webhook.DbtCloud.JobID)
		}
		r.OperatorConfiguration.Webhook.ExecutionBody = types.StringPointerValue(resp.OperatorConfiguration.Webhook.ExecutionBody)
		r.OperatorConfiguration.Webhook.ExecutionURL = types.StringPointerValue(resp.OperatorConfiguration.Webhook.ExecutionURL)
		r.OperatorConfiguration.Webhook.WebhookConfigID = types.StringPointerValue(resp.OperatorConfiguration.Webhook.WebhookConfigID)
		if resp.OperatorConfiguration.Webhook.WebhookType != nil {
			r.OperatorConfiguration.Webhook.WebhookType = types.StringValue(string(*resp.OperatorConfiguration.Webhook.WebhookType))
		} else {
			r.OperatorConfiguration.Webhook.WebhookType = types.StringNull()
		}
	}
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

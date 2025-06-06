// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SynchronousJobRead struct {
	ConfigID                      types.String   `tfsdk:"config_id"`
	ConfigType                    types.String   `tfsdk:"config_type"`
	ConnectorConfigurationUpdated types.Bool     `tfsdk:"connector_configuration_updated"`
	CreatedAt                     types.Int64    `tfsdk:"created_at"`
	EndedAt                       types.Int64    `tfsdk:"ended_at"`
	FailureReason                 *FailureReason `tfsdk:"failure_reason"`
	ID                            types.String   `tfsdk:"id"`
	Logs                          *LogRead       `tfsdk:"logs"`
	Succeeded                     types.Bool     `tfsdk:"succeeded"`
}

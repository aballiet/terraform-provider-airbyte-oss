// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type OperatorConfiguration struct {
	Dbt           *OperatorDbt           `tfsdk:"dbt"`
	Normalization *OperatorNormalization `tfsdk:"normalization"`
	OperatorType  types.String           `tfsdk:"operator_type"`
	Webhook       *OperatorWebhook       `tfsdk:"webhook"`
}
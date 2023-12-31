// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type OperatorDbt struct {
	DbtArguments  types.String `tfsdk:"dbt_arguments"`
	DockerImage   types.String `tfsdk:"docker_image"`
	GitRepoBranch types.String `tfsdk:"git_repo_branch"`
	GitRepoURL    types.String `tfsdk:"git_repo_url"`
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDefinitionResourceModel) ToCreateSDKType() *shared.CustomSourceDefinitionCreate {
	dockerImageTag := r.SourceDefinition.DockerImageTag.ValueString()
	dockerRepository := r.SourceDefinition.DockerRepository.ValueString()
	documentationURL := r.SourceDefinition.DocumentationURL.ValueString()
	icon := new(string)
	if !r.SourceDefinition.Icon.IsUnknown() && !r.SourceDefinition.Icon.IsNull() {
		*icon = r.SourceDefinition.Icon.ValueString()
	} else {
		icon = nil
	}
	name := r.SourceDefinition.Name.ValueString()
	var resourceRequirements *shared.ActorDefinitionResourceRequirements
	if r.SourceDefinition.ResourceRequirements != nil {
		var default1 *shared.ResourceRequirements
		if r.SourceDefinition.ResourceRequirements.Default != nil {
			cpuLimit := new(string)
			if !r.SourceDefinition.ResourceRequirements.Default.CPULimit.IsUnknown() && !r.SourceDefinition.ResourceRequirements.Default.CPULimit.IsNull() {
				*cpuLimit = r.SourceDefinition.ResourceRequirements.Default.CPULimit.ValueString()
			} else {
				cpuLimit = nil
			}
			cpuRequest := new(string)
			if !r.SourceDefinition.ResourceRequirements.Default.CPURequest.IsUnknown() && !r.SourceDefinition.ResourceRequirements.Default.CPURequest.IsNull() {
				*cpuRequest = r.SourceDefinition.ResourceRequirements.Default.CPURequest.ValueString()
			} else {
				cpuRequest = nil
			}
			memoryLimit := new(string)
			if !r.SourceDefinition.ResourceRequirements.Default.MemoryLimit.IsUnknown() && !r.SourceDefinition.ResourceRequirements.Default.MemoryLimit.IsNull() {
				*memoryLimit = r.SourceDefinition.ResourceRequirements.Default.MemoryLimit.ValueString()
			} else {
				memoryLimit = nil
			}
			memoryRequest := new(string)
			if !r.SourceDefinition.ResourceRequirements.Default.MemoryRequest.IsUnknown() && !r.SourceDefinition.ResourceRequirements.Default.MemoryRequest.IsNull() {
				*memoryRequest = r.SourceDefinition.ResourceRequirements.Default.MemoryRequest.ValueString()
			} else {
				memoryRequest = nil
			}
			default1 = &shared.ResourceRequirements{
				CPULimit:      cpuLimit,
				CPURequest:    cpuRequest,
				MemoryLimit:   memoryLimit,
				MemoryRequest: memoryRequest,
			}
		}
		jobSpecific := make([]shared.JobTypeResourceLimit, 0)
		for _, jobSpecificItem := range r.SourceDefinition.ResourceRequirements.JobSpecific {
			jobType := shared.JobType(jobSpecificItem.JobType.ValueString())
			cpuLimit1 := new(string)
			if !jobSpecificItem.ResourceRequirements.CPULimit.IsUnknown() && !jobSpecificItem.ResourceRequirements.CPULimit.IsNull() {
				*cpuLimit1 = jobSpecificItem.ResourceRequirements.CPULimit.ValueString()
			} else {
				cpuLimit1 = nil
			}
			cpuRequest1 := new(string)
			if !jobSpecificItem.ResourceRequirements.CPURequest.IsUnknown() && !jobSpecificItem.ResourceRequirements.CPURequest.IsNull() {
				*cpuRequest1 = jobSpecificItem.ResourceRequirements.CPURequest.ValueString()
			} else {
				cpuRequest1 = nil
			}
			memoryLimit1 := new(string)
			if !jobSpecificItem.ResourceRequirements.MemoryLimit.IsUnknown() && !jobSpecificItem.ResourceRequirements.MemoryLimit.IsNull() {
				*memoryLimit1 = jobSpecificItem.ResourceRequirements.MemoryLimit.ValueString()
			} else {
				memoryLimit1 = nil
			}
			memoryRequest1 := new(string)
			if !jobSpecificItem.ResourceRequirements.MemoryRequest.IsUnknown() && !jobSpecificItem.ResourceRequirements.MemoryRequest.IsNull() {
				*memoryRequest1 = jobSpecificItem.ResourceRequirements.MemoryRequest.ValueString()
			} else {
				memoryRequest1 = nil
			}
			resourceRequirements1 := shared.ResourceRequirements{
				CPULimit:      cpuLimit1,
				CPURequest:    cpuRequest1,
				MemoryLimit:   memoryLimit1,
				MemoryRequest: memoryRequest1,
			}
			jobSpecific = append(jobSpecific, shared.JobTypeResourceLimit{
				JobType:              jobType,
				ResourceRequirements: resourceRequirements1,
			})
		}
		resourceRequirements = &shared.ActorDefinitionResourceRequirements{
			Default:     default1,
			JobSpecific: jobSpecific,
		}
	}
	sourceDefinition := shared.SourceDefinitionCreate{
		DockerImageTag:       dockerImageTag,
		DockerRepository:     dockerRepository,
		DocumentationURL:     documentationURL,
		Icon:                 icon,
		Name:                 name,
		ResourceRequirements: resourceRequirements,
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.CustomSourceDefinitionCreate{
		SourceDefinition: sourceDefinition,
		WorkspaceID:      workspaceID,
	}
	return &out
}

func (r *SourceDefinitionResourceModel) ToUpdateSDKType() *shared.SourceDefinitionUpdate {
	dockerImageTag := r.DockerImageTag.ValueString()
	sourceDefinitionID := r.SourceDefinitionID.ValueString()
	out := shared.SourceDefinitionUpdate{
		DockerImageTag:     dockerImageTag,
		SourceDefinitionID: sourceDefinitionID,
	}
	return &out
}

func (r *SourceDefinitionResourceModel) ToDeleteSDKType() *shared.SourceDefinitionIDRequestBody {
	sourceDefinitionID := r.SourceDefinitionID.ValueString()
	out := shared.SourceDefinitionIDRequestBody{
		SourceDefinitionID: sourceDefinitionID,
	}
	return &out
}

func (r *SourceDefinitionResourceModel) RefreshFromCreateResponse(resp *shared.InvalidInputExceptionInfo) {
	if resp.ExceptionClassName != nil {
		r.ExceptionClassName = types.StringValue(*resp.ExceptionClassName)
	} else {
		r.ExceptionClassName = types.StringNull()
	}
	r.ExceptionStack = nil
	for _, v := range resp.ExceptionStack {
		r.ExceptionStack = append(r.ExceptionStack, types.StringValue(v))
	}
	r.Message = types.StringValue(resp.Message)
	r.ValidationErrors = nil
	for _, validationErrorsItem := range resp.ValidationErrors {
		var validationErrors1 InvalidInputProperty
		if validationErrorsItem.InvalidValue != nil {
			validationErrors1.InvalidValue = types.StringValue(*validationErrorsItem.InvalidValue)
		} else {
			validationErrors1.InvalidValue = types.StringNull()
		}
		if validationErrorsItem.Message != nil {
			validationErrors1.Message = types.StringValue(*validationErrorsItem.Message)
		} else {
			validationErrors1.Message = types.StringNull()
		}
		validationErrors1.PropertyPath = types.StringValue(validationErrorsItem.PropertyPath)
		r.ValidationErrors = append(r.ValidationErrors, validationErrors1)
	}
}

func (r *SourceDefinitionResourceModel) RefreshFromUpdateResponse(resp *shared.InvalidInputExceptionInfo) {
	r.RefreshFromCreateResponse(resp)
}

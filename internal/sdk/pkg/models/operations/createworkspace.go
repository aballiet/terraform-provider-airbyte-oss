// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	WorkspaceRead *shared.WorkspaceRead
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
}

func (o *CreateWorkspaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWorkspaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWorkspaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWorkspaceResponse) GetWorkspaceRead() *shared.WorkspaceRead {
	if o == nil {
		return nil
	}
	return o.WorkspaceRead
}

func (o *CreateWorkspaceResponse) GetInvalidInputExceptionInfo() *shared.InvalidInputExceptionInfo {
	if o == nil {
		return nil
	}
	return o.InvalidInputExceptionInfo
}

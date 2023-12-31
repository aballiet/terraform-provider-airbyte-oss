// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreatePermissionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	PermissionRead *shared.PermissionRead
	// Exception occurred; see message for details.
	KnownExceptionInfo *shared.KnownExceptionInfo
}

func (o *CreatePermissionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePermissionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePermissionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePermissionResponse) GetPermissionRead() *shared.PermissionRead {
	if o == nil {
		return nil
	}
	return o.PermissionRead
}

func (o *CreatePermissionResponse) GetKnownExceptionInfo() *shared.KnownExceptionInfo {
	if o == nil {
		return nil
	}
	return o.KnownExceptionInfo
}

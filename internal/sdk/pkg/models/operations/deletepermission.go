// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeletePermissionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Operation forbidden
	KnownExceptionInfo *shared.KnownExceptionInfo
	// Object with given id was not found.
	NotFoundKnownExceptionInfo *shared.NotFoundKnownExceptionInfo
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
}

func (o *DeletePermissionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePermissionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePermissionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeletePermissionResponse) GetKnownExceptionInfo() *shared.KnownExceptionInfo {
	if o == nil {
		return nil
	}
	return o.KnownExceptionInfo
}

func (o *DeletePermissionResponse) GetNotFoundKnownExceptionInfo() *shared.NotFoundKnownExceptionInfo {
	if o == nil {
		return nil
	}
	return o.NotFoundKnownExceptionInfo
}

func (o *DeletePermissionResponse) GetInvalidInputExceptionInfo() *shared.InvalidInputExceptionInfo {
	if o == nil {
		return nil
	}
	return o.InvalidInputExceptionInfo
}

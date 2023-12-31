// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateUserResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	UserRead *shared.UserRead
	// Exception occurred; see message for details.
	KnownExceptionInfo *shared.KnownExceptionInfo
}

func (o *CreateUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUserResponse) GetUserRead() *shared.UserRead {
	if o == nil {
		return nil
	}
	return o.UserRead
}

func (o *CreateUserResponse) GetKnownExceptionInfo() *shared.KnownExceptionInfo {
	if o == nil {
		return nil
	}
	return o.KnownExceptionInfo
}

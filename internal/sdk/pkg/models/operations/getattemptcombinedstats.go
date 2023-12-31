// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetAttemptCombinedStatsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Operation
	AttemptStats *shared.AttemptStats
	// Object with given id was not found.
	NotFoundKnownExceptionInfo *shared.NotFoundKnownExceptionInfo
}

func (o *GetAttemptCombinedStatsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAttemptCombinedStatsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAttemptCombinedStatsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAttemptCombinedStatsResponse) GetAttemptStats() *shared.AttemptStats {
	if o == nil {
		return nil
	}
	return o.AttemptStats
}

func (o *GetAttemptCombinedStatsResponse) GetNotFoundKnownExceptionInfo() *shared.NotFoundKnownExceptionInfo {
	if o == nil {
		return nil
	}
	return o.NotFoundKnownExceptionInfo
}

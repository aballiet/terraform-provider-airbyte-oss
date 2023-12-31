// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteStreamResetRecordsForJobRequestBody struct {
	ConnectionID string `json:"connectionId"`
	JobID        int64  `json:"jobId"`
}

func (o *DeleteStreamResetRecordsForJobRequestBody) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DeleteStreamResetRecordsForJobRequestBody) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

type DeleteStreamResetRecordsForJobResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteStreamResetRecordsForJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteStreamResetRecordsForJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteStreamResetRecordsForJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

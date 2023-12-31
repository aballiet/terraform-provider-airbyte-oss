// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/aballiet/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListLatestDestinationDefinitionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	DestinationDefinitionReadList *shared.DestinationDefinitionReadList
}

func (o *ListLatestDestinationDefinitionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLatestDestinationDefinitionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLatestDestinationDefinitionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListLatestDestinationDefinitionsResponse) GetDestinationDefinitionReadList() *shared.DestinationDefinitionReadList {
	if o == nil {
		return nil
	}
	return o.DestinationDefinitionReadList
}

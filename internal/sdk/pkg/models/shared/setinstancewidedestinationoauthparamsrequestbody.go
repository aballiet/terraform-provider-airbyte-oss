// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SetInstancewideDestinationOauthParamsRequestBody struct {
	DestinationDefinitionID string                 `json:"destinationDefinitionId"`
	Params                  map[string]interface{} `json:"params"`
}

func (o *SetInstancewideDestinationOauthParamsRequestBody) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}

func (o *SetInstancewideDestinationOauthParamsRequestBody) GetParams() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Params
}

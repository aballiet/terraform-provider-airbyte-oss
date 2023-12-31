// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectionStreamHistoryRequestBody struct {
	ConnectionID string `json:"connectionId"`
	Timezone     string `json:"timezone"`
}

func (o *ConnectionStreamHistoryRequestBody) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ConnectionStreamHistoryRequestBody) GetTimezone() string {
	if o == nil {
		return ""
	}
	return o.Timezone
}

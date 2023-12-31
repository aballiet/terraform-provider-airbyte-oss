// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationCoreConfig struct {
	DestinationID           *string `json:"destinationId,omitempty"`
	DestinationDefinitionID string  `json:"destinationDefinitionId"`
	// The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	WorkspaceID             string      `json:"workspaceId"`
}

func (o *DestinationCoreConfig) GetDestinationID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *DestinationCoreConfig) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}

func (o *DestinationCoreConfig) GetConnectionConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.ConnectionConfiguration
}

func (o *DestinationCoreConfig) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

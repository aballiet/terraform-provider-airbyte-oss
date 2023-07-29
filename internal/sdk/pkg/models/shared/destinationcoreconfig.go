// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationCoreConfig struct {
	// The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	DestinationDefinitionID string      `json:"destinationDefinitionId"`
	DestinationID           *string     `json:"destinationId,omitempty"`
	WorkspaceID             string      `json:"workspaceId"`
}
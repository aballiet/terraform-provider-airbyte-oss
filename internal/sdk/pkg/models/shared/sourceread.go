// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceRead struct {
	SourceDefinitionID string  `json:"sourceDefinitionId"`
	SourceID           string  `json:"sourceId"`
	WorkspaceID        string  `json:"workspaceId"`
	Name               string  `json:"name"`
	SourceName         string  `json:"sourceName"`
	Icon               *string `json:"icon,omitempty"`
}

func (o *SourceRead) GetSourceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.SourceDefinitionID
}

func (o *SourceRead) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *SourceRead) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *SourceRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceRead) GetSourceName() string {
	if o == nil {
		return ""
	}
	return o.SourceName
}

func (o *SourceRead) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

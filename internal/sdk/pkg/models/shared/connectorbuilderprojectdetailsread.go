// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectorBuilderProjectDetailsRead struct {
	Name                             string  `json:"name"`
	BuilderProjectID                 string  `json:"builderProjectId"`
	SourceDefinitionID               *string `json:"sourceDefinitionId,omitempty"`
	ActiveDeclarativeManifestVersion *int64  `json:"activeDeclarativeManifestVersion,omitempty"`
	HasDraft                         bool    `json:"hasDraft"`
}

func (o *ConnectorBuilderProjectDetailsRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectorBuilderProjectDetailsRead) GetBuilderProjectID() string {
	if o == nil {
		return ""
	}
	return o.BuilderProjectID
}

func (o *ConnectorBuilderProjectDetailsRead) GetSourceDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.SourceDefinitionID
}

func (o *ConnectorBuilderProjectDetailsRead) GetActiveDeclarativeManifestVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.ActiveDeclarativeManifestVersion
}

func (o *ConnectorBuilderProjectDetailsRead) GetHasDraft() bool {
	if o == nil {
		return false
	}
	return o.HasDraft
}

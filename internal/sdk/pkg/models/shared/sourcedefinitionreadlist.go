// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceDefinitionReadList struct {
	SourceDefinitions []SourceDefinitionRead `json:"sourceDefinitions"`
}

func (o *SourceDefinitionReadList) GetSourceDefinitions() []SourceDefinitionRead {
	if o == nil {
		return []SourceDefinitionRead{}
	}
	return o.SourceDefinitions
}
